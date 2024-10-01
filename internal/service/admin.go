package service

import "StudentServicePlatform/internal/model"

func QueryUnhandlePost() ([]*model.Post, error) {
	return d.QueryUnhandlePost(ctx)
}

func ReceivePost(post *model.Post) error {
	return d.SavePost(ctx, *post)
}

func QuashPost(post *model.Post) error {
	return d.SavePost(ctx, *post)
}

func ChangeResonse(post *model.Post) error {
	return d.SavePost(ctx, *post)
}