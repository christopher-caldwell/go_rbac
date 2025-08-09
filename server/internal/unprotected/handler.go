package unprotected

import (
	"context"
	"gorbac/internal/server"
)

type UnprotectedHandler struct{}

var _ server.UnprotectedHandler = (*UnprotectedHandler)(nil)

func (h *UnprotectedHandler) GetUnprotectedResource(ctx context.Context) (server.GetUnprotectedResourceRes, error) {
	return &server.UnprotectedResource{
		Message: "Congrats, you captured the unprotected flag.",
	}, nil
}

var Handlers = UnprotectedHandler{}
