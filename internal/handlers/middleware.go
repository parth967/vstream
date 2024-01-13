package handlers

import (
	"fmt"
	"os"

	"github.com/gofiber/fiber/v2"
	jtoken "github.com/golang-jwt/jwt/v4"
	"github.com/joho/godotenv"
)

func AuthMiddleware(ctx *fiber.Ctx) error {
	tokenSignedString := ctx.Cookies("Auth")
	godotenv.Load()
	secretKey := os.Getenv("TOKEN_SECRET")

	token, err := jtoken.Parse(tokenSignedString, func(token *jtoken.Token) (interface{}, error) {
		if _, ok := token.Method.(*jtoken.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(secretKey), nil
	})

	if err != nil || !token.Valid {
		return ctx.Status(fiber.StatusUnauthorized).SendString("Invalid Token")
	}

	username := getTokenInfo(token, "username")
	ctx.Locals("username", username)
	return ctx.Next()
}

func getTokenInfo(token *jtoken.Token, KeyName string) string {
	claims, ok := token.Claims.(jtoken.MapClaims)

	if !ok {
		return ""
	}

	username, ok := claims["username"].(string)

	if !ok {
		return ""
	}

	return username
}
