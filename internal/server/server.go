package server

import (
	"fmt"
	logs "github.com/rusystem/notes-log/pkg/proto"
	"google.golang.org/grpc"
	"net"
)

type Server struct {
	grpcSrv   *grpc.Server
	logServer logs.LogsServer
}

func NewServer(logServer logs.LogsServer) *Server {
	return &Server{
		grpcSrv:   grpc.NewServer(),
		logServer: logServer,
	}
}

func (s *Server) Run(host string, port int) error {
	addr := fmt.Sprintf("%s:%d", host, port)

	lis, err := net.Listen("tcp", addr)
	if err != nil {
		return err
	}

	logs.RegisterLogsServer(s.grpcSrv, s.logServer)

	if err := s.grpcSrv.Serve(lis); err != nil {
		return err
	}

	return nil
}

func (s *Server) Stop() func() {
	return s.grpcSrv.GracefulStop
}
