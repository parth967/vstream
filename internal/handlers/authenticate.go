package handlers

import (
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	jtoken "github.com/golang-jwt/jwt/v4"
	"github.com/joho/godotenv"
	"github.com/vstream/internal/db"
	"github.com/vstream/internal/models"
)

func HandleLogin(ctx *fiber.Ctx) error {
	username := ctx.FormValue("username")
	password := ctx.FormValue("password")

	token, err := findUserByCredentials(username, password, ctx)
	if err != nil {
		ctx.Redirect("/")
	} else {
		ctx.Cookie(&fiber.Cookie{
			Name:  "Auth",
			Value: token,
		})

		ctx.Redirect("/home")
	}
	return nil
}

func findUserByCredentials(username, password string, ctx *fiber.Ctx) (string, error) {
	var user models.User

	isValid, err := db.ValidateUser("username", username, password, &user, ctx)
	if err != nil {
		return "", err
	}

	if isValid {
		token, err := generateToken(user.Name, user.Username, user.UserID)
		if err != nil {
			return "", err
		}
		return token, nil
	}

	return "", errors.New("wrong username or password ")
}

func generateToken(name, username string, userid uint8) (string, error) {
	token := jtoken.NewWithClaims(jtoken.SigningMethodHS256, jtoken.MapClaims{
		"username": username,
		"userid":   userid,
		"name":     name,
		"exp":      time.Now().Add(time.Hour * 24).Unix(),
	})

	godotenv.Load()
	secretKey := os.Getenv("TOKEN_SECRET")

	tokenString, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func HandleSignup(ctx *fiber.Ctx) error {
	newUsername := ctx.FormValue("new-username")
	newPassword := ctx.FormValue("new-password")
	newName := ctx.FormValue("new-name")
	defaultPermission := "admin"

	err := registerUser(newName, newUsername, newPassword, defaultPermission, ctx)
	if err != nil {
		return err
	}
	return nil
}

func registerUser(name, username, password, permission string, ctx *fiber.Ctx) error {
	err := db.AddUser(name, username, password, permission, ctx)

	if err != nil {
		errString := err.Error()
		redirectURL := fmt.Sprintf("/?error=%s", errString)
		return ctx.Redirect(redirectURL)
	}

	redirectURL := fmt.Sprintf("/?success=%s", "Signup successful")
	return ctx.Redirect(redirectURL)

}
