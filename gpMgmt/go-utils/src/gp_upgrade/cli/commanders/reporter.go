package commanders

import (
	"context"
	"fmt"
	pb "gp_upgrade/idl"

	"github.com/greenplum-db/gp-common-go-libs/gplog"
	"github.com/pkg/errors"
)

type Reporter struct {
	client pb.CliToHubClient
}

// UpgradeStepsMessage encode the proper checklist item string to go with a step
//
// Future steps include:
//logger.Info("PENDING - Validate compatible versions for upgrade")
//logger.Info("PENDING - Master server upgrade")
//logger.Info("PENDING - Master OID file shared with segments")
//logger.Info("PENDING - Primary segment upgrade")
//logger.Info("PENDING - Validate cluster start")
//logger.Info("PENDING - Adjust upgrade cluster ports")
var UpgradeStepsMessage = map[pb.UpgradeSteps]string{
	pb.UpgradeSteps_UNKNOWN_STEP:         "- Unknown step",
	pb.UpgradeSteps_CHECK_CONFIG:         "- Configuration Check",
	pb.UpgradeSteps_SEGINSTALL:           "- Install binaries on segments",
	pb.UpgradeSteps_PREPARE_INIT_CLUSTER: "- Initialize upgrade target cluster",
	pb.UpgradeSteps_PREPARE_START_AGENTS: "- Agents Started on Cluster",
	pb.UpgradeSteps_MASTERUPGRADE:        "- Run pg_upgrade on master",
	pb.UpgradeSteps_STOPPED_CLUSTER:      "- Shutdown clusters",
}

func NewReporter(client pb.CliToHubClient) *Reporter {
	return &Reporter{
		client: client,
	}
}

func (r *Reporter) OverallUpgradeStatus() error {
	status, err := r.client.StatusUpgrade(context.Background(), &pb.StatusUpgradeRequest{})
	if err != nil {
		// find some way to expound on the error message? Integration test failing because we no longer log here
		return errors.New("Unable to connect to hub: " + err.Error())
	}

	if status == nil || len(status.ListOfUpgradeStepStatuses) == 0 {
		return errors.New("Received no list of upgrade statuses from hub")
	}

	for _, step := range status.ListOfUpgradeStepStatuses {
		reportString := fmt.Sprintf("%v %s", step.Status,
			UpgradeStepsMessage[step.Step])
		gplog.Info(reportString)
	}

	return nil
}

func (r *Reporter) OverallConversionStatus() error {
	status, err := r.client.StatusConversion(context.Background(), &pb.StatusConversionRequest{})
	if err != nil {
		return errors.New("Unable to connect to hub: " + err.Error())
	}

	if status == nil || status.ConversionStatus == "" {
		return errors.New("Received no conversion status from hub")
	}

	gplog.Info(status.ConversionStatus)

	return nil
}
