package service

import "StudentServicePlatform/internal/model"

func QueryUnhandlePost() ([]model.Post, error) {
	return d.QueryPost(ctx, 0)
}

func UpdatePostStatus(adminID int, postID int, approval int) error {
	return d.UpdatePostStatus(ctx, adminID, postID, approval)
}
func ReceivePost(UserID int,postID int, response string) error {
	return d.ReceivePost(ctx, UserID, postID, response)
}

func QuashPost(postID int) error {
	return d.DeleteResponse(ctx, postID)
}

func ChangeResonse(postID int, response string) error {
	return d.ChangeResonse(ctx,  postID, response)
}