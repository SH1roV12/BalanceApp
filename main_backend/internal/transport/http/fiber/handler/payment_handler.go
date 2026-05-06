package handler

import (
	"github.com/SH1roV12/balance/internal/transport/dto/request"
	"github.com/gofiber/fiber/v2"
)

func(h *Handlers)CreatePayment(ctx *fiber.Ctx)error{
	
	user_id := ctx.Locals("user_id").(string)
	var req request.Payment
	if err := ctx.BodyParser(&req); err != nil{
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error":"bad request"})
	}
	
	res,err := h.paymentService.CreatePayment(ctx.Context(), request.ToDomain(req,user_id))
	if err != nil{
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error":err.Error()})
	}
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{"result":res.PaymentUrl})
}