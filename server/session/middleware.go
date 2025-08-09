package session

import (
	"gorbac/server"
)

type contextKey string

const sessionKey contextKey = "session"

type securityHandler struct{}

var _ server.SecurityHandler = (*securityHandler)(nil)

var SecurityHandler = &securityHandler{}
