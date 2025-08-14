package me

import (
	"context"
	"encoding/json"
	"fmt"
	"gorbac/server"
	"gorbac/session"

	"github.com/casbin/casbin/v2"
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

func (h *MeHandler) GetPermissions(ctx context.Context) (server.GetPermissionsRes, error) {
	user, err := session.GetSessionFromContext(ctx)
	if err != nil {
		return nil, err
	}
	permissions, err := casbin.CasbinJsGetPermissionForUser(session.CasbinEnforcer, user.UserId)
	if err != nil {
		return nil, err
	}

	var perms server.Permissions
	if err := json.Unmarshal([]byte(permissions), &perms); err != nil {
		return nil, fmt.Errorf("failed to unmarshal permissions: %w", err)
	}

	return &perms, nil
}

// func RegisterCasbinPolicyHandler(mux *http.ServeMux) {
// 	fmt.Println("Registering casbin policy handler")
// 	mux.Handle("/permissions", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		permissions, err := casbin.CasbinJsGetPermissionForUser(session.CasbinEnforcer, "123")
// 		if err != nil {
// 			w.WriteHeader(http.StatusInternalServerError)
// 			w.Write([]byte(err.Error()))
// 			return
// 		}
// 		w.WriteHeader(http.StatusOK)
// 		w.Write([]byte(permissions))
// 	}))
// }

var Handlers = MeHandler{}
