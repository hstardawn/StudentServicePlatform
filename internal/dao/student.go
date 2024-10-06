package dao

import (
	"StudentServicePlatform/internal/model"
	"context"
	"time"
)
//提出反馈
func (d *Dao) GetUserByUserID(ctx context.Context, id int) (*model.User, error) {
	var user model.User
	err := d.orm.WithContext(ctx).Model(&model.User{}).Where("id=?", id).First(&user).Error
	return &user, err
}

func (d *Dao) CreatePost(ctx context.Context, post *model.Post) error {
	post.CreateAt = time.Now()
	return d.orm.WithContext(ctx).Create(post).Error
}


// 修改反馈
func (d *Dao) GetPostByID(ctx context.Context, id int) (*model.Post, error) {
	var post model.Post
	err := d.orm.WithContext(ctx).Model(&model.Post{}).Where("id=?", id).First(&post).Error
	return &post, err
}

func (d *Dao) UpdatePost(ctx context.Context, user_id int, id int, is_anonymous int, is_urgent int, post_type int, title string, content string) error {
	err := d.orm.WithContext(ctx).Model(&model.Post{}).Where("id=?", id).Updates(map[string]interface{}{
		// "user_id":      user_id,
		// "id":           id,
		"is_anonymous": is_anonymous,
		"is_urgent":    is_urgent,
		"post_type":    post_type,
		"title":        title,
		"content":      content,
		"updated_at":   time.Now(),
	}).Error
	return err
}


// 删除反馈
func (d *Dao) DeletePost(ctx context.Context, user_id int, id int) error {
	return d.orm.WithContext(ctx).Where("id=?", id).Delete(&model.Post{}).Error
}


// 查看反馈
func (d *Dao) GetPostList(ctx context.Context) ([]model.Post, error) {
	var posts []model.Post
	err := d.orm.WithContext(ctx).Find(&posts).Error
	if err != nil {
		return nil, err
	}
	// for i := range posts {
	// 	posts[i].CreateAt = time.Now()
	// }
	return posts, nil
}

func (d *Dao) GetResponseList(ctx context.Context) ([]model.Response, error) {
	var responses []model.Response
	err := d.orm.WithContext(ctx).Find(&responses).Error
	return responses, err
}

func (d *Dao) GetResponse(ctx context.Context, post_id int) ([]model.Response, error) {
	var responses []model.Response
	err:= d.orm.WithContext(ctx).Model(&model.Response{}).Where("post_id=?", post_id).Find(&responses).Error
	return responses, err
}

func (d *Dao) GetPostByUserID(ctx context.Context, user_id int) ([]model.Post, error){
	var posts []model.Post
	err:= d.orm.WithContext(ctx).Model(&model.Post{}).Where("user_id=?", user_id).Find(&posts).Error
	return posts, err
}

func (d *Dao) GetResponseByPostID(ctx context.Context, post_id int) (*model.Response, error) {
	var response model.Response
	err := d.orm.WithContext(ctx).Model(&model.Response{}).Where("post_id=?", post_id).Find(&response).Error
	return &response, err
}
