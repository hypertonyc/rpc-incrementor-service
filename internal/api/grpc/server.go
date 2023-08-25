package grpc

import (
	"fmt"
	"net"

	"google.golang.org/grpc"

	"github.com/hypertonyc/rpc-incrementor-service/internal/api/grpc/pb"
	"github.com/hypertonyc/rpc-incrementor-service/internal/config"
	"github.com/hypertonyc/rpc-incrementor-service/internal/logger"
	"github.com/hypertonyc/rpc-incrementor-service/internal/service"
)

type Server struct {
	srv *grpc.Server
	cfg config.AppConfig
	svc service.IncrementService
}

func NewServer(cfg config.AppConfig, svc service.IncrementService) *Server {
	return &Server{
		cfg: cfg,
		svc: svc,
	}
}

func (s *Server) Start() error {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", s.cfg.GrpcPort))
	if err != nil {
		logger.Logger.Sugar().Fatalf("failed to listen: %w", err)
		return err
	}

	handler := NewIncrementHandler(s.svc)
	s.srv = grpc.NewServer()
	pb.RegisterIncrementServiceServer(s.srv, handler)
	go s.srv.Serve(lis)

	logger.Logger.Sugar().Infof("gRPC server is listening on port %d", s.cfg.GrpcPort)

	return nil
}

func (s *Server) Stop() {
	s.srv.GracefulStop()
	logger.Logger.Info("gRPC server gracefully stopped")
}
