package protected

import (
	"context"
	"gorbac/server"
)

type ProtectedHandler struct{}

var _ server.ProtectedHandler = (*ProtectedHandler)(nil)

func (h *ProtectedHandler) GetProtectedResource(ctx context.Context) (server.GetProtectedResourceRes, error) {
	return &server.ProtectedResource{
		Message:         "Congrats, you captured the protected flag.",
		SomethingSecret: "This is a secret message that you are authorized to see because you're authenticated.",
	}, nil
}

var Handlers = ProtectedHandler{}
