package service

import (
	"StudentServicePlatform/internal/model"
)

func GetPostByPostId(postId int) (*model.Post, error) {
	return d.GetPostByPostId(ctx, postId)
}

func SavePost(post model.Post) error {
	return d.SavePost(ctx, post)
}
