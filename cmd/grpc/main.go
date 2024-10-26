package main

import (
	"fmt"
	"log"
	"net"
	"time"

	grpcZap "github.com/grpc-ecosystem/go-grpc-middleware/logging/zap"
	grpcRecovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	grpcCtxTags "github.com/grpc-ecosystem/go-grpc-middleware/tags"
	grpcPrometheus "github.com/grpc-ecosystem/go-grpc-prometheus"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"
	"ports-server/configs"
	adapter "ports-server/internal/adapter/handler/grpc"
	repository "ports-server/internal/adapter/repository/in"
	service "ports-server/internal/core/service/grpc"
	"ports-server/internal/core/util"
	grpc2 "ports-server/pkg/api/grpc"
)

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	c, err := configs.NewConfig()
	if err != nil {
		return err
	}

	l, err := util.New(c)
	if err != nil {
		return err
	}
	defer func() {
		if err := l.Sync(); err != nil {
			log.Printf("failed sync logger: %v", err)
		}
	}()

	grpcServer := newGrpcServer(c, l.Logger)
	storagePorts := repository.NewStorageIn(c)
	portsLogic := service.NewPortsLogic(l, storagePorts)

	server := adapter.NewServerGRPC(l.Logger, portsLogic)
	grpc2.RegisterEmulatorPortsServer(grpcServer, server)

	lis, err := net.Listen("tcp", fmt.Sprintf("%s:%s", c.Grpc.Host, c.Grpc.Port))
	if err != nil {
		return fmt.Errorf("failed to strat listener: %w", err)
	}

	deferFunc := func(net.Listener) {
		if err = lis.Close(); err != nil {
			l.Error("error while closing listener", zap.Error(err))
		}
	}

	if err != nil {
		return err
	}
	defer deferFunc(lis)

	if err = grpcServer.Serve(lis); err != nil {
		return fmt.Errorf("failed to serve: %w", err)
	}

	return nil
}

func newGrpcServer(c *configs.Config, l *zap.Logger) *grpc.Server {
	s := grpc.NewServer(
		grpc.KeepaliveParams(keepalive.ServerParameters{Timeout: time.Second * time.Duration(c.Grpc.Timeout)}),
		grpc.MaxRecvMsgSize(c.Grpc.MaxGrpcReceive),
		grpc.MaxSendMsgSize(c.Grpc.MaxGrpcSend),
		grpc.ChainUnaryInterceptor(
			grpcRecovery.UnaryServerInterceptor(),
			grpcCtxTags.UnaryServerInterceptor(),
			grpcPrometheus.UnaryServerInterceptor,
			grpcZap.UnaryServerInterceptor(l, grpcZap.WithMessageProducer(util.MessageProducer)),
		),
		grpc.ChainStreamInterceptor(
			grpcRecovery.StreamServerInterceptor(),
			grpcCtxTags.StreamServerInterceptor(),
			grpcPrometheus.StreamServerInterceptor,
			grpcZap.StreamServerInterceptor(l, grpcZap.WithMessageProducer(util.MessageProducer)),
		),
	)

	return s
}
