package middleware

import (
	"log"

	"github.com/SH1roV12/balance/internal/pkg/jwt"
	"github.com/gofiber/fiber/v2"
)


func JWTMidleware(c *fiber.Ctx)error{
	user_id,err := jwt.ParseAccessToken(c)
	if err != nil{
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error":"empty or invalid token"})
	}
	log.Println("jwt middleware", "checking middleware", "user_id", user_id)
	c.Locals("user_id",user_id)
	return c.Next()

}