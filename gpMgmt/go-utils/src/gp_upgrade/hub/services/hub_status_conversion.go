package services

import (
	"golang.org/x/net/context"

	"strings"

	pb "gp_upgrade/idl"
)

func (h *HubClient) StatusConversion(ctx context.Context, in *pb.StatusConversionRequest) (*pb.StatusConversionReply, error) {
	conns, err := h.AgentConns()
	if err != nil {
		return nil, err
	}

	var statuses []string
	for _, conn := range conns {
		status, err := pb.NewAgentClient(conn).CheckConversionStatus(context.Background(), &pb.CheckConversionStatusRequest{})
		if err != nil {
			return nil, err
		}
		statuses = append(statuses, status.GetStatus()...)
	}

	return &pb.StatusConversionReply{
		ConversionStatus: strings.Join(statuses, "\n"),
	}, nil
}
