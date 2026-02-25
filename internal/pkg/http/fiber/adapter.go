package fiber

import (
	customHttp "balance/internal/pkg/http"

	"github.com/gofiber/fiber/v2"
)

type App struct{
	app *fiber.App
}


func New(config fiber.Config)customHttp.Server{
	return &App{
		app: fiber.New(config),
	}
	
}