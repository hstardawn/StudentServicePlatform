package service

import (
	"StudentServicePlatform/internal/model"
	"time"
	// "time"
)

func QueryUnhandlePost() ([]model.Post, error) {
	postList, err := d.QueryPost(ctx, 0)
	return postList, err
}

func UpdatePostStatus(adminID int, postID int, status int) error {
	err := d.UpdatePostStatus(ctx, adminID, postID, status)
	return err
}

func ReceivePost(adminID int, postID int, response string) error {
	err := d.ReceivePost(ctx, &model.Response{
		PostID:   postID,
		AdminID:  adminID,
		Response: response,
		CreateAt: time.Now(),
	})
	return err
}
func QuashPost(postID int) error {
	err := d.DeleteResponse(ctx, postID)
	return err
}

func ChangeResponse(postID int, response string) error {
	err := d.ChangeResponse(ctx, postID, response)
	return err
}

func QueryTrash() ([]model.Post, error) {
	postList, err := d.QueryPost(ctx, 2)
	return postList, err
}

func HandleTrash(adminID int, postID int, approval int) error {
	if approval == 1 {
		err := DeletePost(adminID, postID)
		return err
	} else {
		err := UpdatePostStatus(adminID, postID, 0)
		return err
	}
}

func QueryAdmin() ([]model.User, error) {
	adminList, err := d.QueryAdmin(ctx)
	return adminList, err
}

func UpdateUserType(userID int, userType int) error {
	err := d.UpdateUserType(ctx, userID, userType)
	return err
}

func GetPostByAdminID(admin_id int) ([]model.Post, error) {
	return d.GetPostByAdminID(ctx, admin_id)
}

func GetResponseByPID(postID int) (*model.Response, error) {
	return d.GetResponseByPID(ctx, postID)
}
