package dao

import (
	"StudentServicePlatform/internal/model"
	"context"
)

func (d *Dao) StorePostPicture(ctx context.Context,ID int,filename string) error{
	return d.orm.WithContext(ctx).Model(&model.Post{}).Where("post_id=?", ID).Updates(map[string]interface{}{
		"pictures": filename,
	}).Error
}

func (d *Dao)StoreUserPicture(ctx context.Context,ID int,filename string) error{
	return d.orm.WithContext(ctx).Model(&model.User{}).Where("user_id=?", ID).Updates(map[string]interface{}{
		"pictures": filename,
	}).Error
}