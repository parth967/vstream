package models

import (
	"errors"
)

func FindByCredentials(email, password string) (User, error) {
	if email == "test@mail.com" && password == "test12345" {
		return User{
			ID:             1,
			Email:          "test@mail.com",
			Password:       "test12345",
			FavoritePhrase: "Hello, World!",
		}, nil
	}
	return User{
		ID:             0,
		Email:          "",
		Password:       "",
		FavoritePhrase: "",
	}, errors.New("user not found")
}
