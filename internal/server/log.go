package server

import (
	"context"
	"github.com/rusystem/notes-log/internal/service"
	log "github.com/rusystem/notes-log/pkg/domain"
)

type LogServer struct {
	service service.Service
}

func NewLogServer(service service.Service) *LogServer {
	return &LogServer{
		service: service,
	}
}

func (h *LogServer) Log(ctx context.Context, req *log.LogRequest) (*log.Empty, error) {
	err := h.service.Insert(ctx, req)

	return &log.Empty{}, err
}
