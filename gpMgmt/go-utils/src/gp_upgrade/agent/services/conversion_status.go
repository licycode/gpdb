package services

import (
	"context"

	pb "gp_upgrade/idl"
)

func (s *AgentServer) CheckConversionStatus(ctx context.Context, in *pb.CheckConversionStatusRequest) (*pb.CheckConversionStatusReply, error) {
	return nil, nil
}
