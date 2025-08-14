package main

import (
	"gorbac/server"
	auth "gorbac/session"
	"log"
	"net/http"
	"os"

	"gorbac/internal/me"
	"gorbac/internal/protected"
	"gorbac/internal/rbac"
	"gorbac/internal/unprotected"
	"gorbac/openapi"

	"github.com/rs/cors"
)

type handler struct {
	unprotected.UnprotectedHandler
	protected.ProtectedHandler
	rbac.RbacHandler
	me.MeHandler
}

func main() {

	service := &handler{
		UnprotectedHandler: unprotected.Handlers,
		ProtectedHandler:   protected.Handlers,
		RbacHandler:        rbac.Handlers,
		MeHandler:          me.Handlers,
	}

	srv, err := server.NewServer(service, auth.SecurityHandler, server.NotFoundHandler, server.GenericErrorHandler)
	if err != nil {
		log.Fatal(err)
	}
	mux := http.NewServeMux()

	mux.Handle("/", srv)
	openapi.RegisterOpenApiSpec(mux)
	// me.RegisterCasbinPolicyHandler(mux)

	c := cors.New(cors.Options{
		AllowCredentials: true,
		AllowedOrigins:   []string{"http://localhost:*", "https://www.realizedproperties.cloud"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	handler := c.Handler(mux)
	if err := http.ListenAndServe(":"+port, handler); err != nil {
		log.Fatal(err)
	}
}
