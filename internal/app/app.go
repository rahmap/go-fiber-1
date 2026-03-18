package app

import (
	"fmt"
	"log"

	"fiber-rest-api/internal/config"
	"fiber-rest-api/internal/http/middleware"
	"fiber-rest-api/internal/modules/user"
	"fiber-rest-api/internal/platform/database"

	"github.com/gofiber/fiber/v2"
)

type App struct {
	fiber  *fiber.App
	config config.Config
}

func New() (*App, error) {
	cfg := config.Load()

	db, err := database.NewMySQL(cfg.DatabaseDSN)
	if err != nil {
		return nil, fmt.Errorf("connect database: %w", err)
	}

	userRepository := user.NewRepository(db)
	userService := user.NewService(userRepository)
	userHandler := user.NewHandler(userService)

	fiberApp := fiber.New(fiber.Config{
		AppName: cfg.AppName,
	})

	apiV1 := fiberApp.Group("/api/v1", middleware.ValidateToken(cfg.APIToken))
	user.RegisterRoutes(apiV1, userHandler)

	return &App{
		fiber:  fiberApp,
		config: cfg,
	}, nil
}

func (a *App) Run() error {
	log.Printf("database connected")
	log.Printf("starting %s on %s", a.config.AppName, a.config.Address())

	return a.fiber.Listen(a.config.Address())
}
