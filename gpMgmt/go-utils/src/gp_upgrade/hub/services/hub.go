package services

import (
	"context"
	"os"
	"path/filepath"
	"time"

	"gp_upgrade/hub/cluster"
	"gp_upgrade/hub/upgradestatus"

	"github.com/greenplum-db/gp-common-go-libs/gplog"
	"github.com/pkg/errors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/connectivity"
)

const (
	port = "6416"
)

var DialTimeout = 3 * time.Second

type dialer func(ctx context.Context, target string, opts ...grpc.DialOption) (*grpc.ClientConn, error)

type reader interface {
	GetHostnames() ([]string, error)
	OfOldClusterConfig()
}

type HubClient struct {
	Bootstrapper

	agentConns   []*grpc.ClientConn
	clusterPair  cluster.PairOperator
	configreader reader
	grpcDialer   dialer
}

func NewHub(pair cluster.PairOperator, configReader reader, grpcDialer dialer) (*HubClient, func()) {
	// refactor opportunity -- don't use this pattern,
	// use different types or separate functions for old/new or set the config path at reader initialization time
	configReader.OfOldClusterConfig()
	gpUpgradeDir := filepath.Join(os.Getenv("HOME"), ".gp_upgrade")

	h := &HubClient{
		clusterPair:  pair,
		configreader: configReader,
		grpcDialer:   grpcDialer,
		Bootstrapper: Bootstrapper{
			hostnameGetter: configReader,
			remoteExecutor: NewClusterSsher(upgradestatus.NewChecklistManager(gpUpgradeDir), NewPingerManager()),
		},
	}

	return h, h.closeConns
}

func (h *HubClient) AgentConns() ([]*grpc.ClientConn, error) {
	if h.agentConns != nil {
		err := h.ensureConnsAreReady()
		if err != nil {
			return nil, err
		}

		return h.agentConns, nil
	}

	hostnames, err := h.configreader.GetHostnames()
	if err != nil {
		return nil, err
	}

	for _, host := range hostnames {
		ctx, _ := context.WithTimeout(context.TODO(), DialTimeout)
		conn, err := h.grpcDialer(ctx, host+":"+port, grpc.WithInsecure(), grpc.WithBlock())
		if err != nil {
			return nil, err
		}
		h.agentConns = append(h.agentConns, conn)
	}

	return h.agentConns, nil
}

func (h *HubClient) ensureConnsAreReady() error {
	for i := 0; i < 3; i++ {
		ready := 0
		for _, conn := range h.agentConns {
			if conn.GetState() == connectivity.Ready {
				ready++
			}
		}

		if ready == len(h.agentConns) {
			return nil
		}

		time.Sleep(500 * time.Millisecond)
	}

	return errors.New("at least one connection is not ready")
}

func (h *HubClient) closeConns() {
	for _, conn := range h.agentConns {
		err := conn.Close()
		if err != nil {
			gplog.Info("Error closing hub to agent connection: ", err.Error())
		}
	}
}
