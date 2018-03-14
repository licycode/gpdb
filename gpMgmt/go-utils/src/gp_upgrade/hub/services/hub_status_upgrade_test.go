package services_test

import (
	"errors"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"gp_upgrade/hub/configutils"
	pb "gp_upgrade/idl"
	"gp_upgrade/utils"

	"github.com/greenplum-db/gp-common-go-libs/testhelper"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"google.golang.org/grpc"

	"gp_upgrade/hub/services"
	"gp_upgrade/testutils"
)

var _ = Describe("status upgrade", func() {
	var (
		listener                 *services.HubClient
		fakeStatusUpgradeRequest *pb.StatusUpgradeRequest
		shutdownHub              func()
	)

	BeforeEach(func() {
		testhelper.SetupTestLogger() // extend to capture the values in a var if future tests need it
		//any mocking of utils.System function pointers should be reset by calling InitializeSystemFunctions
		utils.System = utils.InitializeSystemFunctions()

		reader := configutils.NewReader()
		listener, shutdownHub = services.NewHub(nil, &reader, grpc.DialContext)
	})

	AfterEach(func() {
		shutdownHub()
	})

	It("responds with the statuses of the steps based on files on disk", func() {
		dir, err := ioutil.TempDir("", "")
		Expect(err).ToNot(HaveOccurred())
		defer os.RemoveAll(dir)
		services.HomeDir = dir

		setStateFile(dir, "seginstall", "completed")
		setStateFile(dir, "share-oids", "failed")

		f, err := os.Create(filepath.Join(dir, ".gp_upgrade", "cluster_config.json"))
		Expect(err).ToNot(HaveOccurred())
		f.Close()

		resp, err := listener.StatusUpgrade(nil, &pb.StatusUpgradeRequest{})
		Expect(err).To(BeNil())

		Expect(resp.ListOfUpgradeStepStatuses).To(ConsistOf(
			[]*pb.UpgradeStepStatus{
				{
					Step:   pb.UpgradeSteps_CHECK_CONFIG,
					Status: pb.StepStatus_COMPLETE,
				}, {
					Step:   pb.UpgradeSteps_PREPARE_INIT_CLUSTER,
					Status: pb.StepStatus_PENDING,
				}, {
					Step:   pb.UpgradeSteps_SEGINSTALL,
					Status: pb.StepStatus_COMPLETE,
				}, {
					Step:   pb.UpgradeSteps_STOPPED_CLUSTER,
					Status: pb.StepStatus_PENDING,
				}, {
					Step:   pb.UpgradeSteps_MASTERUPGRADE,
					Status: pb.StepStatus_PENDING,
				}, {
					Step:   pb.UpgradeSteps_PREPARE_START_AGENTS,
					Status: pb.StepStatus_PENDING,
				}, {
					Step:   pb.UpgradeSteps_SHARE_OIDS,
					Status: pb.StepStatus_FAILED,
				},
			},
		))
	})

	// TODO: Get rid of these tests once full rewritten test coverage exists
	Describe("creates a reply", func() {
		It("sends status messages under good condition", func() {
			formulatedResponse, err := listener.StatusUpgrade(nil, fakeStatusUpgradeRequest)
			Expect(err).To(BeNil())
			countOfStatuses := len(formulatedResponse.GetListOfUpgradeStepStatuses())
			Expect(countOfStatuses).ToNot(BeZero())
		})

		It("reports that prepare start-agents is pending", func() {
			utils.System.FilePathGlob = func(string) ([]string, error) {
				return []string{"somefile"}, nil
			}

			var fakeStatusUpgradeRequest *pb.StatusUpgradeRequest

			formulatedResponse, err := listener.StatusUpgrade(nil, fakeStatusUpgradeRequest)
			Expect(err).To(BeNil())

			stepStatuses := formulatedResponse.GetListOfUpgradeStepStatuses()

			var stepStatusSaved *pb.UpgradeStepStatus
			for _, stepStatus := range stepStatuses {

				if stepStatus.GetStep() == pb.UpgradeSteps_PREPARE_START_AGENTS {
					stepStatusSaved = stepStatus
				}
			}
			Expect(stepStatusSaved.GetStep()).ToNot(BeZero())
			Expect(stepStatusSaved.GetStatus()).To(Equal(pb.StepStatus_PENDING))
		})

		It("reports that prepare start-agents is running and then complete", func() {
			var numInvocations int
			utils.System.FilePathGlob = func(input string) ([]string, error) {
				numInvocations += 1
				if numInvocations == 1 {
					return []string{filepath.Join(filepath.Dir(input), "in.progress")}, nil
				}
				return []string{filepath.Join(filepath.Dir(input), "completed")}, nil
			}
			utils.System.Stat = func(name string) (os.FileInfo, error) {
				return nil, nil
			}
			pollStatusUpgrade := func() pb.StepStatus {
				response, _ := listener.StatusUpgrade(nil, &pb.StatusUpgradeRequest{})

				stepStatuses := response.GetListOfUpgradeStepStatuses()

				var stepStatusSaved *pb.UpgradeStepStatus
				for _, stepStatus := range stepStatuses {

					if stepStatus.GetStep() == pb.UpgradeSteps_PREPARE_START_AGENTS {
						stepStatusSaved = stepStatus
					}
				}
				return stepStatusSaved.GetStatus()

			}

			//Expect(stepStatusSaved.GetStep()).ToNot(BeZero())
			Eventually(pollStatusUpgrade).Should(Equal(pb.StepStatus_COMPLETE))
		})

		It("reports that master upgrade is pending when pg_upgrade dir does not exist", func() {
			utils.System.IsNotExist = func(error) bool {
				return true
			}

			formulatedResponse, err := listener.StatusUpgrade(nil, fakeStatusUpgradeRequest)
			Expect(err).To(BeNil())

			stepStatuses := formulatedResponse.GetListOfUpgradeStepStatuses()

			for _, stepStatus := range stepStatuses {
				if stepStatus.GetStep() == pb.UpgradeSteps_MASTERUPGRADE {
					Expect(stepStatus.GetStatus()).To(Equal(pb.StepStatus_PENDING))
				}
			}
		})

		It("reports that master upgrade is running when pg_upgrade/*.inprogress files exists", func() {
			utils.System.IsNotExist = func(error) bool {
				return false
			}
			utils.System.FilePathGlob = func(string) ([]string, error) {
				return []string{"somefile.inprogress"}, nil
			}
			utils.System.ExecCmdOutput = func(cmd string, args ...string) ([]byte, error) {
				return []byte("123"), nil
			}

			formulatedResponse, err := listener.StatusUpgrade(nil, fakeStatusUpgradeRequest)
			Expect(err).To(BeNil())

			stepStatuses := formulatedResponse.GetListOfUpgradeStepStatuses()

			for _, stepStatus := range stepStatuses {
				if stepStatus.GetStep() == pb.UpgradeSteps_MASTERUPGRADE {
					Expect(stepStatus.GetStatus()).To(Equal(pb.StepStatus_RUNNING))
				}
			}
		})

		It("reports that master upgrade is done when no *.inprogress files exist in ~/.gp_upgrade/pg_upgrade", func() {
			utils.System.IsNotExist = func(error) bool {
				return false
			}
			utils.System.FilePathGlob = func(glob string) ([]string, error) {
				if strings.Contains(glob, "inprogress") {
					return nil, errors.New("fake error")
				} else if strings.Contains(glob, "done") {
					return []string{"found something"}, nil
				}

				return nil, errors.New("test not configured for this glob")
			}
			utils.System.ExecCmdOutput = func(cmd string, args ...string) ([]byte, error) {
				return []byte(""), errors.New("bogus error")
			}
			utils.System.Stat = func(filename string) (os.FileInfo, error) {
				if strings.Contains(filename, "found something") {
					return &testutils.FakeFileInfo{}, nil
				}
				return nil, nil
			}

			utils.System.Open = func(name string) (*os.File, error) {
				// Temporarily create a file that we can read as a real file descriptor
				fd, err := ioutil.TempFile("/tmp", "hub_status_upgrade_test")
				Expect(err).To(BeNil())

				filename := fd.Name()
				fd.WriteString("12312312;Upgrade complete;\n")
				fd.Close()
				return os.Open(filename)

			}

			formulatedResponse, err := listener.StatusUpgrade(nil, fakeStatusUpgradeRequest)
			Expect(err).To(BeNil())

			stepStatuses := formulatedResponse.GetListOfUpgradeStepStatuses()

			for _, stepStatus := range stepStatuses {
				if stepStatus.GetStep() == pb.UpgradeSteps_MASTERUPGRADE {
					Expect(stepStatus.GetStatus()).To(Equal(pb.StepStatus_COMPLETE))
				}
			}
		})

		It("reports pg_upgrade has failed", func() {
			utils.System.IsNotExist = func(error) bool {
				return false
			}
			utils.System.FilePathGlob = func(glob string) ([]string, error) {
				if strings.Contains(glob, "inprogress") {
					return nil, errors.New("fake error")
				} else if strings.Contains(glob, "done") {
					return []string{"found something"}, nil
				}

				return nil, errors.New("test not configured for this glob")
			}
			utils.System.ExecCmdOutput = func(cmd string, args ...string) ([]byte, error) {
				return []byte(""), errors.New("bogus error")
			}
			utils.System.Open = func(name string) (*os.File, error) {
				// Temporarily create a file that we can read as a real file descriptor
				fd, err := ioutil.TempFile("/tmp", "hub_status_upgrade_test")
				Expect(err).To(BeNil())

				filename := fd.Name()
				fd.WriteString("12312312;Upgrade failed;\n")
				fd.Close()
				return os.Open(filename)

			}
			formulatedResponse, err := listener.StatusUpgrade(nil, fakeStatusUpgradeRequest)
			Expect(err).To(BeNil())

			stepStatuses := formulatedResponse.GetListOfUpgradeStepStatuses()

			for _, stepStatus := range stepStatuses {
				if stepStatus.GetStep() == pb.UpgradeSteps_MASTERUPGRADE {
					Expect(stepStatus.GetStatus()).To(Equal(pb.StepStatus_FAILED))
				}
			}
		})
	})

	Describe("Status of PrepareNewClusterConfig", func() {
		It("marks this step pending if there's no new cluster config file", func() {
			utils.System.Stat = func(filename string) (os.FileInfo, error) {
				return nil, errors.New("cannot find file") /* This is normally a PathError */
			}
			stepStatus, err := services.GetPrepareNewClusterConfigStatus()
			Expect(err).To(BeNil()) // convert file-not-found errors into stepStatus
			Expect(stepStatus.Step).To(Equal(pb.UpgradeSteps_PREPARE_INIT_CLUSTER))
			Expect(stepStatus.Status).To(Equal(pb.StepStatus_PENDING))
		})

		It("marks this step complete if there is a new cluster config file", func() {
			utils.System.Stat = func(filename string) (os.FileInfo, error) {
				return nil, nil
			}

			stepStatus, err := services.GetPrepareNewClusterConfigStatus()
			Expect(err).To(BeNil())
			Expect(stepStatus.Step).To(Equal(pb.UpgradeSteps_PREPARE_INIT_CLUSTER))
			Expect(stepStatus.Status).To(Equal(pb.StepStatus_COMPLETE))

		})

	})

	Describe("Status of ShutdownClusters", func() {
		It("We're sending the status of shutdown clusters", func() {
			formulatedResponse, err := listener.StatusUpgrade(nil, fakeStatusUpgradeRequest)
			Expect(err).To(BeNil())
			countOfStatuses := len(formulatedResponse.GetListOfUpgradeStepStatuses())
			Expect(countOfStatuses).ToNot(BeZero())
			found := false
			for _, v := range formulatedResponse.GetListOfUpgradeStepStatuses() {
				if pb.UpgradeSteps_STOPPED_CLUSTER == v.Step {
					found = true
				}
			}
			Expect(found).To(Equal(true))
		})
	})
})

func setStateFile(dir string, step string, state string) {
	err := os.MkdirAll(filepath.Join(dir, ".gp_upgrade", step), os.ModePerm)
	Expect(err).ToNot(HaveOccurred())

	f, err := os.Create(filepath.Join(dir, ".gp_upgrade", step, state))
	Expect(err).ToNot(HaveOccurred())
	f.Close()
}
