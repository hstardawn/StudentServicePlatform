package dao

import (
	"StudentServicePlatform/internal/model"
	"context"
	"time"
)

func (d *Dao) QueryPost(ctx context.Context, status int) ([]model.Post, error) {
	var postList []model.Post
	err := d.orm.WithContext(ctx).Find(&postList, "status = ?",status).Error
	return postList, err
}

func (d *Dao) UpdatePostStatus(ctx context.Context,adminID int, postID int, status int) error{
	err := d.orm.WithContext(ctx).Model(&model.Post{}).Where("id=?", postID).Updates(map[string]interface{}{
		"status": status,
		"admin_id": adminID,
	}).Error
	return err
}

// func (d *Dao)GetPostResponseTime(ctx context.Context, postID int) (time.Time, error){
// 	var response model.Response
// 	err := d.orm.WithContext(ctx).Model(&model.Response{}).Where("post_id=?", postID).Find(&response).Error
// 	create_at := response.CreateAt
// 	return create_at, err
// }
// func (d *Dao)UpdatePostResponseTime(ctx context.Context, ID int,response_time time.Time) error{
// 	err := d.orm.WithContext(ctx).Model(&model.Post{}).Where("id=?", ID).Updates(&model.Post{ResponseAt: response_time}).Error
// 	return err
// }

func(d *Dao) ReceivePost(ctx context.Context,response *model.Response) error {
	// response := model.Response{Response: content, PostID: post_id, AdminID: user_id}
	response.CreateAt = time.Now()
	err := d.orm.WithContext(ctx).Create(response).Error
	return err
}

func (d *Dao)DeleteResponse(ctx context.Context, postID int) error{
	err := d.orm.WithContext(ctx).Delete(&model.Response{}, postID).Error
	return err
}

func (d *Dao)ChangeResponse(ctx context.Context, postID int, response string) error{
	err := d.orm.WithContext(ctx).Model(&model.Response{}).Where("post_id= ?", postID).Update("response", response).Error
	return err
}

func (d *Dao)QueryAdmin(ctx context.Context) ([]model.User, error) {
	var adminList []model.User
	err := d.orm.WithContext(ctx).Find(&adminList).Error
	return adminList, err
}

func (d *Dao) UpdateUserType(ctx context.Context,userID int,userType int) error{
	err := d.orm.WithContext(ctx).Model(&model.User{}).Where("ID=?", userID).Updates(map[string]interface{}{
		"user_type": userType,
	}).Error
	return err
}

func (d *Dao) GetPostByAdminID(ctx context.Context, admin_id int) ([]model.Post, error){
	var posts []model.Post
	err:= d.orm.WithContext(ctx).Model(&model.Post{}).Where("admin_id=?", admin_id).Find(&posts).Error
	// if err!= nil {
	// 	return nil, err
	// }

	// PostID := posts.ID
	return posts, err
}