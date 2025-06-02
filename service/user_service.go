package service

import (
	"apisix-backend/repository"
	"apisix-backend/structs"

	"golang.org/x/crypto/bcrypt"
)

func LoginService(username, password string) (string, *structs.MyUsersEntity, error) {
	user, err := repository.FindUserByUsername(username)
	if err != nil {
		return "", nil, err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return "", nil, err
	}

	token, err := CreateToken(user.ID, user.Username)
	if err != nil {
		return "", nil, err
	}

	return token, user, nil
}
