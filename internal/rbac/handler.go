package rbac

import (
	"context"
	"gorbac/internal/server"
)

type RbacHandler struct{}

var _ server.RbacHandler = (*RbacHandler)(nil)

func (h *RbacHandler) CreateRbacResource(ctx context.Context, req *server.RbacResource) (server.CreateRbacResourceRes, error) {
	return &server.RbacResource{
		Message:             req.Message,
		SomethingSecret:     req.SomethingSecret,
		SomethingVerySecret: req.SomethingVerySecret,
	}, nil
}

func (h *RbacHandler) GetRbacResource(ctx context.Context) (server.GetRbacResourceRes, error) {
	return &server.RbacResource{
		Message:             "Congrats, you captured the rbac flag.",
		SomethingSecret:     "This is a secret message that you are authorized to see because you're authenticated",
		SomethingVerySecret: "This is a very secret message that you are authorized to see because you have the right rbac permissions.",
	}, nil
}

var Handlers = RbacHandler{}
