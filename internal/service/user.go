package service

import (
	"StudentServicePlatform/internal/model"
)

func GetUserByUserid(userid int) (*model.User, error) {
	return d.GetUserByUserid(ctx, userid)
}