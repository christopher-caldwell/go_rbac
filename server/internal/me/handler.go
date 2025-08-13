package me

import (
	"context"
	"gorbac/server"
	"gorbac/session"
)

type MeHandler struct{}

var _ server.MeHandler = (*MeHandler)(nil)

func (h *MeHandler) GetMe(ctx context.Context) (server.GetMeRes, error) {
	user, err := session.GetSessionFromContext(ctx)
	if err != nil {
		return nil, err
	}
	return &server.User{
		UserID:    user.UserId,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
	}, nil
}

var Handlers = MeHandler{}
