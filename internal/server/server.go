package server

import (
	"fmt"
	log "github.com/rusystem/notes-log/pkg/domain"
	"google.golang.org/grpc"
	"net"
)

type Server struct {
	grpcSrv   *grpc.Server
	logServer log.LogsServer
}

func NewServer(logServer log.LogsServer) *Server {
	return &Server{
		grpcSrv:   grpc.NewServer(),
		logServer: logServer,
	}
}

func (s *Server) Run(port int) error {
	addr := fmt.Sprintf(":%d", port)

	lis, err := net.Listen("tcp", addr)
	if err != nil {
		return err
	}

	log.RegisterLogsServer(s.grpcSrv, s.logServer)

	if err := s.grpcSrv.Serve(lis); err != nil {
		return err
	}

	return nil
}

func (s *Server) Stop() func() {
	return s.grpcSrv.Stop
}
