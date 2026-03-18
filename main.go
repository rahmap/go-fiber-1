package main

import (
	"log"

	"fiber-rest-api/internal/app"
)

func main() {
	application, err := app.New()
	if err != nil {
		log.Fatalf("failed to initialize application: %v", err)
	}

	if err := application.Run(); err != nil {
		log.Fatalf("failed to run application: %v", err)
	}
}
