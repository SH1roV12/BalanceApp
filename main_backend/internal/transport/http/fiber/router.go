package http

import (
	"github.com/SH1roV12/balance/internal/transport/http/fiber/handler"
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App,handlers *handler.Handlers){
	api := app.Group("/api/v1")
	api.Post("/new",handlers.Register)
}


