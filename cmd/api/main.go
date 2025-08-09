package main

import (
	"gorbac/internal/server"
	auth "gorbac/internal/session"
	"log"
	"net/http"
	"os"

	"gorbac/internal/protected"
	"gorbac/internal/rbac"
	"gorbac/internal/unprotected"
	"gorbac/openapi"
)

type handler struct {
	unprotected.UnprotectedHandler
	protected.ProtectedHandler
	rbac.RbacHandler
}

func main() {

	service := &handler{
		UnprotectedHandler: unprotected.Handlers,
		ProtectedHandler:   protected.Handlers,
		RbacHandler:        rbac.Handlers,
	}

	srv, err := server.NewServer(service, auth.SecurityHandler, server.NotFoundHandler, server.GenericErrorHandler)
	if err != nil {
		log.Fatal(err)
	}
	mux := http.NewServeMux()
	mux.Handle("/", srv)
	openapi.RegisterOpenApiSpec(mux)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	if err := http.ListenAndServe(":"+port, mux); err != nil {
		log.Fatal(err)
	}

}
