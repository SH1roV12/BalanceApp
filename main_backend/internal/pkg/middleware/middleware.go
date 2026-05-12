package middleware

import (
	"log"

	"github.com/SH1roV12/balance/internal/pkg/config"
	"github.com/SH1roV12/balance/internal/pkg/jwt"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)




func JWTMidleware(cfg *config.Api)fiber.Handler{
	return func(c *fiber.Ctx)error{
	user_id,err := jwt.ParseAccessToken(c,cfg)
	if err != nil{
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error":"empty or invalid token"})
	}
	log.Println("jwt middleware", "checking middleware", "user_id", user_id)
	c.Locals("user_id",user_id)
	return c.Next()

}
}


func Logger(sugar *zap.SugaredLogger)fiber.Handler{
	return func(c *fiber.Ctx)error{
		sugar.Infow("incoming request",
            "method", c.Method(),
            "path",   c.Path(),
            "ip",     c.IP(),
        )

        err := c.Next()

        sugar.Infow("request completed",
            "status", c.Response().StatusCode(),
            "path",   c.Path(),
        )

        return err
	}
}