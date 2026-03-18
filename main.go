package main

import (
	"log"

	_ "fiber-rest-api/docs"
	"fiber-rest-api/internal/app"
)

// @title Fiber REST API
// @version 1.0
// @description Simple Fiber REST API with Swagger documentation for the user module.
// @BasePath /
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name X-TOKEN

func main() {
	application, err := app.New()
	if err != nil {
		log.Fatalf("failed to initialize application: %v", err)
	}

	if err := application.Run(); err != nil {
		log.Fatalf("failed to run application: %v", err)
	}
}
