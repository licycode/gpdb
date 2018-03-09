package services_test

import (
	"errors"
	"strings"

	"gp_upgrade/testutils"
	"gp_upgrade/utils"

	"github.com/greenplum-db/gp-common-go-libs/testhelper"

	"google.golang.org/grpc"

	pb "gp_upgrade/idl"

	"gp_upgrade/hub/services"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("hub", func() {
	var (
		hubClient   *services.HubClient
		shutdownHub func()
		agentA      *testutils.MockAgentServer
	)

	BeforeEach(func() {
		testhelper.SetupTestLogger()
		utils.System = utils.InitializeSystemFunctions()

		agentA = testutils.NewMockAgentServer()

		reader := spyReader{}
		reader.hostnames = []string{"localhost", "localhost"}
		hubClient, shutdownHub = services.NewHub(nil, &reader, grpc.DialContext)
	})

	AfterEach(func() {
		shutdownHub()

		agentA.Stop()
	})

	It("receives a conversion status from each agent and returns all as single message", func() {
		statusMessages := []string{"status", "status"}
		agentA.StatusConversionResponse = &pb.CheckConversionStatusReply{
			Status: statusMessages,
		}

		status, err := hubClient.StatusConversion(nil, &pb.StatusConversionRequest{})
		Expect(err).To(BeNil())

		statusList := strings.Split(status.GetConversionStatus(), "\n")
		Expect(statusList).To(Equal([]string{"status", "status", "status", "status"}))
	})

	It("returns an error when AgentConns returns an error", func() {
		agentA.Stop()

		_, err := hubClient.StatusConversion(nil, &pb.StatusConversionRequest{})
		Expect(err).To(HaveOccurred())
	})

	It("returns an error when Agent server returns an error", func() {
		agentA.StatusConversionErr = errors.New("any error")

		_, err := hubClient.StatusConversion(nil, &pb.StatusConversionRequest{})
		Expect(err).To(HaveOccurred())
	})
})
