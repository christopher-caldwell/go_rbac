package session

import (
	"context"
	"errors"
	"gorbac/server"
	logger "gorbac/util/log"

	"github.com/ogen-go/ogen/ogenerrors"
)

type User struct {
	UserId    string   `json:"user_id"`
	FirstName string   `json:"first_name"`
	LastName  string   `json:"last_name"`
	Email     string   `json:"email"`
	Roles     []string `json:"roles"`
}

var mockUser = User{
	UserId:    "123",
	FirstName: "User",
	LastName:  "Person",
	Email:     "user@example.com",
}
var mockAdmin = User{
	UserId:    "456",
	FirstName: "Admin",
	LastName:  "Person",
	Email:     "admin@example.com",
}

func (h *securityHandler) HandleBearerAuth(ctx context.Context, operationName server.OperationName, t server.BearerAuth) (context.Context, error) {
	// JWT check
	sessionToken := t.Token

	var user *User
	switch sessionToken {
	case "123":
		user = &mockUser
	case "456":
		logger.Logger.Info("Session token", "token", sessionToken)
		user = &mockAdmin
	default:
		return nil, ogenerrors.ErrSecurityRequirementIsNotSatisfied
	}

	ctx = addSessionToContext(ctx, user)
	return ctx, nil
}

func addSessionToContext(ctx context.Context, session *User) context.Context {
	return context.WithValue(ctx, sessionKey, session)
}

func GetSessionFromContext(ctx context.Context) (*User, error) {
	session, ok := ctx.Value(sessionKey).(*User)
	if !ok || session == nil {
		logger.Logger.Error("Session not found in context")
		return nil, errors.New("session not found in context")
	}
	return session, nil
}
