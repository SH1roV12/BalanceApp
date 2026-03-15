package handler

import (
	"github.com/SH1roV12/balance/internal/transport/http/dto/request"
	"github.com/SH1roV12/balance/internal/transport/http/dto/response"
	"github.com/gofiber/fiber/v2"
)





func(h *Handlers) Register(ctx *fiber.Ctx)error{
	var req request.RegisterUser
	
	if err := ctx.BodyParser(&req); err !=nil{
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error":"bad request"})
	}
	err := h.userService.NewUser(ctx.Context(),&req)
	if err != nil{
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error":err.Error()})
	}

	return ctx.Status(fiber.StatusCreated).JSON(fiber.Map{"message":"successfully created"})
}

func(h *Handlers)GetAllUsers(ctx *fiber.Ctx)error{
	users,err := h.userService.GetAll(ctx.Context())
	if err != nil{
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error":err.Error()})
	}
	return ctx.Status(fiber.StatusOK).JSON(response.FromEntitysToDTOs(users))
}