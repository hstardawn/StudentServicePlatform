package dao

import (
	"StudentServicePlatform/internal/model"
	"context"
)

func (d *Dao) GetUserByUserid(ctx context.Context, userId int) (*model.User, error) {
	var user *model.User
	err := d.orm.WithContext(ctx).First(&user, "id = ?",userId).Error
	return user, err
}