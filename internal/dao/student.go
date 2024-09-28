package dao

import (
	"StudentServicePlatform/internal/model"
	"context"
)

func (d *Dao) GetPostByPostId(ctx context.Context, postId int)(*model.Post, error){
	var post model.Post
	err := d.orm.WithContext(ctx).First(&post, "id = ?",postId).Error
	return &post, err
}

func (d *Dao) SavePost(ctx context.Context, post model.Post)error{
	err := d.orm.WithContext(ctx).Save(post).Error
	return err
}