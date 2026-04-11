package handler

import (
	"errors"
	"time"

	customErrors "github.com/SH1roV12/balance/internal/pkg/errors"
	"github.com/SH1roV12/balance/internal/pkg/jwt"
	"github.com/SH1roV12/balance/internal/transport/http/dto/request"
	"github.com/SH1roV12/balance/internal/transport/http/dto/response"
	"github.com/gofiber/fiber/v2"
)





func(h *Handlers) Register(ctx *fiber.Ctx)error{
	var req request.RegisterUser
	if err := ctx.BodyParser(&req); err !=nil{
		h.sugar.Errorw("register", "parse req","error",err.Error())
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error":"bad request"})
	}
	err := h.userService.NewUser(ctx.Context(),&req)
	if err != nil{
		var customErr customErrors.MyError
		errors.As(err,&customErr)
		if errors.Is(err, customErrors.Repo.User.AlreadyExist){
			h.sugar.Errorw(customErr.Location,"error",customErr.RawError )
			return ctx.Status(fiber.StatusConflict).JSON(fiber.Map{"error":err.Error()})
		}
		h.sugar.Errorw(customErr.Location,"error",customErr.RawError )
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error":err.Error()})
	}
	
	return ctx.Status(fiber.StatusCreated).JSON(fiber.Map{"message":"successfully registered"})
}


func (h *Handlers) Login(ctx *fiber.Ctx)error{
	var req request.Login
	if err := ctx.BodyParser(&req); err != nil{
		h.sugar.Errorw("login", "parse req","error",err.Error())
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error":"bad request"})
	}
	user,err := h.userService.GetByEmail(ctx.Context(),req.Email, req.Password)
	
	
	if err!=nil{
		var customErr customErrors.MyError
		errors.As(err,&customErr)
		h.sugar.Errorw(customErr.Location,"error",customErr.RawError )
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error":err.Error()}) // Match to password wrong 
	}
	
	access,err := jwt.GenAccessToken(user.ID,h.sugar)
	if err != nil{
		h.sugar.Errorw("jwt_refresh","get access token", "error", err)
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error":err.Error()})
	}
	refresh,err := jwt.GenRefreshToken(user.ID,h.sugar)
	if err != nil{
		h.sugar.Errorw("jwt_refresh", "get refresh token","error", err)
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error":err.Error()})
	}
	ctx.Cookie(&fiber.Cookie{
		Name: "access_token",
		Value: access,
		Expires: time.Now().Add(time.Minute*20),
		HTTPOnly: true,
		Secure: false,
		SameSite: "Lax",
	})

	ctx.Cookie(&fiber.Cookie{
		Name: "refresh_token",
		Value: refresh,
		Expires: time.Now().Add(time.Hour*24*7),
		HTTPOnly: true,
		Secure: false,
		SameSite: "Lax",
	})

	return ctx.Status(fiber.StatusOK).JSON(response.FromEntityToDTO(user))
}


func(h *Handlers)GetAllUsers(ctx *fiber.Ctx)error{
	users,err := h.userService.GetAll(ctx.Context())
	
	if err != nil{
		var customErr customErrors.MyError
		errors.As(err,&customErr)
		h.sugar.Errorw(customErr.Location,"error",customErr.RawError )
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error":err.Error()})
	}
	
	return ctx.Status(fiber.StatusOK).JSON(response.FromEntitysToDTOs(users))
}


func(h *Handlers)Refresh(ctx *fiber.Ctx)error{
	user_id,err := jwt.ParseRefreshToken(ctx)
	if err != nil{
		h.sugar.Errorw("jwt_refresh", "error", err)
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error":"empty or invalid token"})
	}
	access,err := jwt.GenAccessToken(user_id,h.sugar)
	if err != nil{
		h.sugar.Errorw("jwt_refresh","get access token", "error", err)
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error":"failed to generate new access token"})
	}
	refresh,err := jwt.GenRefreshToken(user_id,h.sugar)
	if err != nil{
		h.sugar.Errorw("jwt_refresh", "get refresh token","error", err)
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error":"failed to generate new refresh token"})
	}
	ctx.Cookie(&fiber.Cookie{
		Name: "access_token",
		Value: access,
		Expires: time.Now().Add(time.Minute*20),
		HTTPOnly: true,
		Secure: false,
		SameSite: "Lax",
	})

	ctx.Cookie(&fiber.Cookie{
		Name: "refresh_token",
		Value: refresh,
		Expires: time.Now().Add(time.Hour*24*7),
		HTTPOnly: true,
		Secure: false,
		SameSite: "Lax",
	})
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{})
}

func(h *Handlers)GetUserByID(ctx *fiber.Ctx)error{
	user_id := ctx.Locals("user_id").(string)
	user,err := h.userService.GetById(ctx.Context(),user_id)

	if err != nil{
		var customErr customErrors.MyError
		errors.As(err,&customErr)
		h.sugar.Errorw(customErr.Location,"error",customErr.RawError )
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error":err.Error()})
	}
	return ctx.Status(fiber.StatusOK).JSON(response.FromEntityToDTO(user))
}