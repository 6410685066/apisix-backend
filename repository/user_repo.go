package repository

import (
	"apisix-backend/share"
	"apisix-backend/structs"
)

func FindUserByUsername(username string) (*structs.MyUsersEntity, error) {
	var user structs.MyUsersEntity
	err := share.DB.Where("username = ?", username).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}
