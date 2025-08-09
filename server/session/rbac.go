package session

import (
	"context"
	"gorbac/internal/server"
	logger "gorbac/util/log"

	"github.com/casbin/casbin/v2"
)

var enforcer *casbin.Enforcer

func (h *securityHandler) HandleRbac(ctx context.Context, operationName server.OperationName, t server.Rbac) (context.Context, error) {
	if enforcer == nil {
		var err error
		enforcer, err = casbin.NewEnforcer("rbac_model.conf", "rbac_policy.csv")
		if err != nil {
			logger.Logger.Error("Failed to create enforcer", "error", err)
		}
	}

	sub := "alice" // the user that wants to access a resource.
	obj := "data1" // the resource that is going to be accessed.
	act := "read"  // the operation that the user performs on the resource.

	ok, err := enforcer.Enforce(sub, obj, act)

	if err != nil {
		logger.Logger.Error("Failed to enforce", "error", err)
	}

	if ok {
		logger.Logger.Info("permit alice to read data1")
	} else {
		logger.Logger.Error("deny the request, show an error")
	}

	return ctx, nil
}
