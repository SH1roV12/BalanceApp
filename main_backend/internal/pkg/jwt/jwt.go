package jwt

import (
	"errors"
	"log"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"go.uber.org/zap"
)

var (
	accessSecret  = []byte(os.Getenv("JWT_ACCESS_SECRET"))
	refreshSecret = []byte(os.Getenv("JWT_REFRESH_SECRET"))
)
func GenAccessToken(id string,sugar *zap.SugaredLogger)(string,error){
	sugar.Infow("gen access token", "generating...", id)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,jwt.MapClaims{
		"user_id":  id,
		"exp":  time.Now().Add(time.Minute * 20).Unix(), 
	})
	
	return token.SignedString(accessSecret)
}

func GenRefreshToken(id string,sugar *zap.SugaredLogger)(string,error){
	sugar.Infow("gen refresh token", "generating...", id)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,jwt.MapClaims{
		"user_id":  id,
		"exp":  time.Now().Add(time.Hour * 24*7).Unix(), 
	})
	
	return token.SignedString(refreshSecret)
}


func ParseAccessToken(ctx *fiber.Ctx)(string,error){
	log.Println("parse access token", "checking" )
	tokenString := ctx.Cookies("access_token")
	if tokenString == ""{
		return "",errors.New("empty token")
	}
	token,err := jwt.Parse(tokenString,func(t *jwt.Token) (any, error) {
		return []byte(accessSecret),nil
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

func ParseRefreshToken(ctx *fiber.Ctx)(string,error){
	log.Println("parse refresh token", "checking" )
	tokenString := ctx.Cookies("refresh_token")
	if tokenString == ""{
		return "",errors.New("empty token")
	}
	token,err := jwt.Parse(tokenString,func(t *jwt.Token) (any, error) {
		return []byte(refreshSecret),nil
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
