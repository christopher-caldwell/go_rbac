package session

import (
	"context"
	"gorbac/server"
	logger "gorbac/util/log"
	"os"
	"path/filepath"

	"github.com/casbin/casbin/v2"
	"github.com/ogen-go/ogen/ogenerrors"
)

func (h *securityHandler) HandleRbac(ctx context.Context, operationName server.OperationName, t server.Rbac) (context.Context, error) {
	logger.Logger.Info("Handling RBAC", "operationName", operationName, "token", t.Token)

	user, err := GetSessionFromContext(ctx)
	if err != nil {
		logger.Logger.Error("Failed to get session from context", "error", err)
		return ctx, err
	}

	modelPath, policyPath, err := getPathsForRbacEnforcer()
	if err != nil {
		logger.Logger.Error("Failed to get paths for RBAC enforcer", "error", err)
		return ctx, err
	}

	enforcer, err := casbin.NewEnforcer(*modelPath, *policyPath)
	if err != nil {
		logger.Logger.Error("Failed to create enforcer", "error", err)
		return ctx, err
	}

	permitted, err := enforcer.Enforce(user.UserId, operationName, "data1")
	if err != nil {
		logger.Logger.Error("Failed to enforce", "error", err)
		return ctx, err
	}

	if permitted {
		logger.Logger.Info("permit admin to read data1")
	} else {
		logger.Logger.Error("deny the request, show an error")
		return ctx, ogenerrors.ErrSecurityRequirementIsNotSatisfied
	}

	return ctx, nil
}

func getPathsForRbacEnforcer() (*string, *string, error) {

	// Get the current working directory and build paths to the session directory
	wd, err := os.Getwd()
	if err != nil {
		logger.Logger.Error("Failed to get working directory", "error", err)
		return nil, nil, err
	}

	// Build paths relative to the working directory
	// Assuming we're running from the server directory or project root
	sessionDir := filepath.Join(wd, "session")
	if _, err := os.Stat(sessionDir); os.IsNotExist(err) {
		// Try from project root
		sessionDir = filepath.Join(wd, "server", "session")
		if _, err := os.Stat(sessionDir); os.IsNotExist(err) {
			logger.Logger.Error("Could not find session directory", "tried", filepath.Join(wd, "session"), "and", filepath.Join(wd, "server", "session"))
			return nil, nil, err
		}
	}

	modelPath := filepath.Join(sessionDir, "rbac_model.conf")
	policyPath := filepath.Join(sessionDir, "rbac_policy.csv")

	// Check if files exist
	if _, err := os.Stat(modelPath); os.IsNotExist(err) {
		logger.Logger.Error("RBAC model file does not exist", "path", modelPath)
		return nil, nil, err
	}
	if _, err := os.Stat(policyPath); os.IsNotExist(err) {
		logger.Logger.Error("RBAC policy file does not exist", "path", policyPath)
		return nil, nil, err
	}

	return &modelPath, &policyPath, nil
}
