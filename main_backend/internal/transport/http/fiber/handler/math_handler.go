package handler

import (
	"github.com/SH1roV12/balance/internal/transport/http/dto/request"
	"github.com/gofiber/fiber/v2"
)

func(h *Handlers)GetSum(ctx *fiber.Ctx)error{
	var req request.Sum
	if err := ctx.BodyParser(&req); err != nil{
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error":"bad request"})
	}
	
	sum,err := h.mathService.Sum(ctx.Context(), request.ToDomain(req))
	if err != nil{
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error":err.Error()})
	}
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{"result":sum.Result})
}