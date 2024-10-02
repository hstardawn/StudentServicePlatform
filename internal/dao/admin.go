package dao

import (
	"StudentServicePlatform/internal/model"
	"context"
)

func (d *Dao) QueryPost(ctx context.Context, status int) ([]model.Post, error) {
	var postList []model.Post
	err := d.orm.WithContext(ctx).Find(&postList, "status = ?",status).Error
	return postList, err
}

func (d *Dao) UpdatePostStatus(ctx context.Context,adminID int, postID int, approval int) error{
	err := d.orm.WithContext(ctx).Model(&model.Post{}).Where("id=?", postID).Updates(map[string]interface{}{
		"status": approval,
		"admin_id": adminID,
	}).Error
	return err
}

func(d *Dao) ReceivePost(ctx context.Context,UserID int, postID int, content string) error {
	response := model.Response{Response: content, PostID: postID, AdminID: UserID}
	err := d.orm.WithContext(ctx).Create(&response).Error
	return err
}

func (d *Dao)DeleteResponse(ctx context.Context, postID int) error{
	err := d.orm.WithContext(ctx).Delete(&model.Response{}, postID).Error
	return err
}

func (d *Dao)ChangeResonse(ctx context.Context, postID int, response string) error{
	err := d.orm.WithContext(ctx).Model(&model.Response{}).Where("post_id= ?", postID).Update("post_id", postID).Error
	return err
}