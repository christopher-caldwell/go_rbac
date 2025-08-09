package session

import (
	"context"
	"errors"
	"gorbac/internal/server"

	"github.com/clerk/clerk-sdk-go/v2"
	"github.com/clerk/clerk-sdk-go/v2/jwt"
	"github.com/docker/docker/daemon/logger"
	"github.com/ogen-go/ogen/ogenerrors"
)

func (h *securityHandler) HandleBearerAuth(ctx context.Context, operationName server.OperationName, t server.BearerAuth) (context.Context, error) {
	sessionToken := t.Token

	// Verify the session
	claims, err := jwt.Verify(ctx, &jwt.VerifyParams{
		Token: sessionToken,
	})
	if err != nil {
		return nil, ogenerrors.ErrSecurityRequirementIsNotSatisfied
	}

	usr, err := clerkuser.Get(ctx, claims.Subject)
	if err != nil {
		return nil, ogenerrors.ErrSecurityRequirementIsNotSatisfied
	}

	// Can do RBAC here, using operation name and by getting the session from Redis. Need to rename APIKey to SessionID tho
	ctx = addSessionToContext(ctx, usr)
	return ctx, nil
}

func addSessionToContext(ctx context.Context, session *clerk.User) context.Context {
	return context.WithValue(ctx, sessionKey, session)
}

func GetSessionFromContext(ctx context.Context) (*clerk.User, error) {
	session, ok := ctx.Value(sessionKey).(*clerk.User)
	if !ok || session == nil {
		logger.Logger.Error("Session not found in context")
		return nil, errors.New("session not found in context")
	}
	return session, nil
}
