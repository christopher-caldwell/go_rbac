package session

import (
	"context"
	"errors"
	"gorbac/internal/server"
	logger "gorbac/util/log"
)

type User struct {
	UserId    string `json:"user_id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
}

var mockUser = User{
	UserId:    "123",
	FirstName: "John",
	LastName:  "Doe",
	Email:     "john.doe@example.com",
}

func (h *securityHandler) HandleBearerAuth(ctx context.Context, operationName server.OperationName, t server.BearerAuth) (context.Context, error) {
	// JWT check
	sessionToken := t.Token
	logger.Logger.Info("Session token", "token", sessionToken)

	ctx = addSessionToContext(ctx, &mockUser)
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
