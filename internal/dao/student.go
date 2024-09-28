package dao

import (
	"StudentServicePlatform/internal/model"
	"context"
	"time"
)

func (d *Dao) CreatePost(ctx context.Context, post *model.Post) error {
	return d.orm.WithContext(ctx).Create(post).Error
}

func (d *Dao) UpdatePost(ctx context.Context, user_id int, post_id int, is_anonymous int, is_urgent int, post_type int, title string, content string) error {
	err := d.orm.WithContext(ctx).Model(&model.Post{}).Where("post_id=?", post_id).Updates(map[string]interface{}{
		"user_id":      user_id,
		"post_id":      post_id,
		"is_anonymous": is_anonymous,
		"is_urgent":    is_urgent,
		"post_type":    post_type,
		"title":        title,
		"content":      content,
	}).Error
	return err
}

func (d *Dao) DeletePost(ctx context.Context, user_id int, post_id int) error {
	return d.orm.WithContext(ctx).Where("post_id=?", post_id).Delete(&model.Post{}).Error
}

func (d *Dao) GetPost(ctx context.Context) ([]model.Post, error) {
	var posts []model.Post
	err := d.orm.WithContext(ctx).Find(&posts).Error
	if err != nil {
		return nil, err
	}
	for i := range posts {
		posts[i].CreateAt = time.Now()
	}
	return posts, nil
}