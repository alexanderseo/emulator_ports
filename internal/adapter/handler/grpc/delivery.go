package adapter

import (
	"context"

	"go.uber.org/zap"
	"google.golang.org/protobuf/types/known/emptypb"
	service "ports-server/internal/core/service/grpc"
	"ports-server/pkg/api/grpc"
)

type ServerGRPC struct {
	l          *zap.Logger
	portsLogic service.PortsLogicTester
	grpc.UnimplementedEmulatorPortsServer
}

func NewServerGRPC(
	l *zap.Logger,
	portsLogic service.PortsLogicTester,
) *ServerGRPC {
	return &ServerGRPC{
		l:          l,
		portsLogic: portsLogic,
	}
}

func (s *ServerGRPC) Read(ctx context.Context, in *emptypb.Empty) (*grpc.Answer, error) {
	answer, err := s.portsLogic.Read(ctx)
	return nil, nil
}

func (s *ServerGRPC) Write(ctx context.Context, in *emptypb.Empty) (*grpc.Answer, error) {

	return nil, nil
}
