package services

import (
	"errors"
	"gp_upgrade/hub/configutils"
	"gp_upgrade/hub/upgradestatus"
	pb "gp_upgrade/idl"
	"gp_upgrade/utils"
	"os"
	"path/filepath"

	"github.com/greenplum-db/gp-common-go-libs/gplog"
	"golang.org/x/net/context"
)

func (s *HubClient) StatusUpgrade(ctx context.Context, in *pb.StatusUpgradeRequest) (*pb.StatusUpgradeReply, error) {
	gplog.Info("starting StatusUpgrade")

	demoStepStatus := &pb.UpgradeStepStatus{
		Step:   pb.UpgradeSteps_CHECK_CONFIG,
		Status: pb.StepStatus_PENDING,
	}

	prepareInitStatus, _ := GetPrepareNewClusterConfigStatus()

	homeDirectory := os.Getenv("HOME")
	if homeDirectory == "" {
		return nil, errors.New("could not find the HOME environment")
	}

	seginstallStatePath := filepath.Join(homeDirectory, ".gp_upgrade/seginstall")
	gplog.Debug("looking for seginstall State at %s", seginstallStatePath)
	seginstallState := upgradestatus.NewStateCheck(seginstallStatePath, pb.UpgradeSteps_SEGINSTALL)
	seginstallStatus, _ := seginstallState.GetStatus()

	gpstopStatePath := filepath.Join(homeDirectory, ".gp_upgrade/gpstop")
	clusterPair := upgradestatus.NewShutDownClusters(gpstopStatePath)
	shutdownClustersStatus, _ := clusterPair.GetStatus()

	pgUpgradePath := filepath.Join(homeDirectory, ".gp_upgrade/pg_upgrade")
	convertMaster := upgradestatus.NewConvertMaster(pgUpgradePath)
	masterUpgradeStatus, _ := convertMaster.GetStatus()

	startAgentsStatePath := filepath.Join(homeDirectory, ".gp_upgrade/start-agents")
	prepareStartAgentsState := upgradestatus.NewStateCheck(startAgentsStatePath, pb.UpgradeSteps_PREPARE_START_AGENTS)
	startAgentsStatus, _ := prepareStartAgentsState.GetStatus()

	shareOidsPath := filepath.Join(homeDirectory, ".gp_upgrade/share-oids")
	shareOidsState := upgradestatus.NewStateCheck(shareOidsPath, pb.UpgradeSteps_SHARE_OIDS)
	shareOidsStatus, _ := shareOidsState.GetStatus()

	reply := &pb.StatusUpgradeReply{}
	reply.ListOfUpgradeStepStatuses = append(reply.ListOfUpgradeStepStatuses, demoStepStatus, seginstallStatus,
		prepareInitStatus, shutdownClustersStatus, masterUpgradeStatus, startAgentsStatus, shareOidsStatus)
	return reply, nil
}

func GetPrepareNewClusterConfigStatus() (*pb.UpgradeStepStatus, error) {
	/* Treat all stat failures as cannot find file. Conceal worse failures atm.*/
	_, err := utils.System.Stat(configutils.GetNewClusterConfigFilePath())

	if err != nil {
		gplog.Debug("%v", err)
		return &pb.UpgradeStepStatus{Step: pb.UpgradeSteps_PREPARE_INIT_CLUSTER,
			Status: pb.StepStatus_PENDING}, nil
	}

	return &pb.UpgradeStepStatus{Step: pb.UpgradeSteps_PREPARE_INIT_CLUSTER,
		Status: pb.StepStatus_COMPLETE}, nil
}
