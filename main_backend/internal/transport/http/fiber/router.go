package http

import (
	"github.com/SH1roV12/balance/internal/pkg/config"
	"github.com/SH1roV12/balance/internal/pkg/middleware"
	"github.com/SH1roV12/balance/internal/transport/http/fiber/handler"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"go.uber.org/zap"
)

func SetupRoutes(app *fiber.App,handlers *handler.Handlers,sugar *zap.SugaredLogger,config config.Api){
	app.Use(cors.New(cors.Config{
    	AllowOrigins: "http://localhost:5500, http://127.0.0.1:5500",
    	AllowHeaders: "Origin, Content-Type, Accept, Authorization",
		AllowCredentials: true,
    
    	AllowMethods:     "GET,POST,HEAD,PUT,DELETE,PATCH,OPTIONS",
	}))
	api := app.Group("/api/v1",middleware.Logger(sugar) )
	
	
	auth := api.Group("/auth",middleware.JWTMidleware(&config))
	auth.Post("/payment", handlers.CreatePayment)
	api.Post("/new",handlers.Register)
	api.Get("/users",handlers.GetAllUsers)
	api.Get("/refresh",handlers.Refresh)
	auth.Get("/get",handlers.GetUserByID)
	api.Post("/login",handlers.Login)
}


