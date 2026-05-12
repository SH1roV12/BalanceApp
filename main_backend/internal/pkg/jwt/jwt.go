package jwt

import (
	"errors"
	"log"
	"time"

	"github.com/SH1roV12/balance/internal/pkg/config"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"go.uber.org/zap"
)


func GenAccessToken(id string,sugar *zap.SugaredLogger,config *config.Api)(string,error){
	sugar.Infow("gen access token", "generating...", id)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,jwt.MapClaims{
		"user_id":  id,
		"exp":  time.Now().Add(time.Minute * 20).Unix(), 
	})
	
	return token.SignedString([]byte(config.Access_Secret))
}

func GenRefreshToken(id string,sugar *zap.SugaredLogger,config *config.Api)(string,error){
	sugar.Infow("gen refresh token", "generating...", id)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,jwt.MapClaims{
		"user_id":  id,
		"exp":  time.Now().Add(time.Hour * 24*7).Unix(), 
	})
	
	return token.SignedString([]byte(config.Refresh_Secret))
}


func ParseAccessToken(ctx *fiber.Ctx,config *config.Api)(string,error){
	log.Println("parse access token", "checking" )
	tokenString := ctx.Cookies("access_token")
	if tokenString == ""{
		return "",errors.New("empty token")
	}
	token,err := jwt.Parse(tokenString,func(t *jwt.Token) (any, error) {
		return []byte(config.Access_Secret),nil
	})
	if err != nil || !token.Valid{
		return "", err
	}
	
	if claims,ok := token.Claims.(jwt.MapClaims); ok{
		return claims["user_id"].(string),nil
	}else{
		return "",errors.New("failed to get user id")
	}
}

func ParseRefreshToken(ctx *fiber.Ctx,config *config.Api)(string,error){
	log.Println("parse refresh token", "checking" )
	tokenString := ctx.Cookies("refresh_token")
	if tokenString == ""{
		return "",errors.New("empty token")
	}
	token,err := jwt.Parse(tokenString,func(t *jwt.Token) (any, error) {
		return []byte(config.Refresh_Secret),nil
	})
	if err != nil || !token.Valid{
		return "", err
	}
	if claims,ok := token.Claims.(jwt.MapClaims); ok{
		return claims["user_id"].(string),nil
	}else{
		return "",errors.New("failed to get user id")
	}
}
