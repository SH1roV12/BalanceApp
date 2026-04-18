package http

import (
	"github.com/SH1roV12/balance/internal/pkg/middleware"
	"github.com/SH1roV12/balance/internal/transport/http/fiber/handler"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

func SetupRoutes(app *fiber.App,handlers *handler.Handlers,sugar *zap.SugaredLogger){
	api := app.Group("/api/v1",middleware.Logger(sugar) )
	
	api.Post("/sum", handlers.GetSum)
	auth := api.Group("/auth",middleware.JWTMidleware)
	
	api.Post("/new",handlers.Register)
	api.Get("/users",handlers.GetAllUsers)
	api.Get("/refresh",handlers.Refresh)
	auth.Get("/get",handlers.GetUserByID)
	api.Get("/login",handlers.Login)
}


