package dao

import (
	// "CMS/internal/model"
	"StudentServicePlatform/internal/model"
	"context"
)

func (d *Dao) CreateUser(ctx context.Context, user *model.User) error {
	return d.orm.WithContext(ctx).Create(user).Error
}

func (d *Dao) GetUserPassword(ctx context.Context, username int) (*model.User, error) {
	var user model.User
	err := d.orm.WithContext(ctx).Where("username=?", username).First(&user).Error
	return &user, err
}

func (d *Dao) GetUserByUsername(ctx context.Context, username int) (*model.User, error) {
	var user model.User
	err := d.orm.WithContext(ctx).Where("username=?", username).First(&user).Error
	return &user, err
}

func (d *Dao) UpdateUser(ctx context.Context, username int, name string, sex string, phone_num int, email string, password string) error {
	return d.orm.WithContext(ctx).Model(&model.User{}).Where("username=?", username).Updates(map[string]interface{}{
		"username":  username,
		"name":      name,
		"sex":       sex,
		"phone_num": phone_num,
		"email":     email,
		"password":  password,
	}).Error
}