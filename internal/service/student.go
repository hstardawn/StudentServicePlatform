package service

import "StudentServicePlatform/internal/model"

func CreatePost(user_id int, is_anonymous int, is_urgent int, post_type int, title string, content string) error {
	return d.CreatePost(ctx, &model.Post{
		UserID:      user_id,
		IsAnonymous: is_anonymous,
		IsUrgent:    is_urgent,
		PostType:    post_type,
		Title:       title,
		Content:     content,
	})
}

func UpdatePost(user_id int, post_id int, is_anonymous int, is_urgent int, post_type int, title string, content string) error {
	return d.UpdatePost(ctx, user_id, post_id, is_anonymous, is_urgent, post_type, title, content)
	// return d.CreatePost(ctx,&model.Post{
	// 	UserID:user_id,
	// 	PostID:post_id,
	// 	IsAnonymous:is_anonymous,
	// 	IsUrgent:is_urgent,
	// 	PostType:post_type,
	// 	Title:title,
	// 	Content:content,
	// })
}

func DeletePost(user_id int, post_id int) error {
	return d.DeletePost(ctx, user_id, post_id)
}

func GetPost() ([]model.Post, error) {
	return d.GetPost(ctx)
}