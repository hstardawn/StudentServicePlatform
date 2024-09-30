package service

import (
	"StudentServicePlatform/internal/model"
	"time"
)


//提出反馈
func CreatePost(user_id int, is_anonymous int, is_urgent int, post_type int, title string, content string) error {
	return d.CreatePost(ctx, &model.Post{
		UserID:      user_id,
		IsAnonymous: is_anonymous,
		IsUrgent:    is_urgent,
		PostType:    post_type,
		Title:       title,
		Content:     content,
		CreateAt: time.Now(),
	})
}

func GetUserByUserID(user_id int) (*model.User, error) {
	return d.GetUserByUserID(ctx, user_id)
}


//修改反馈
func GetPostByID(id int)(*model.Post,error){
	return d.GetPostByID(ctx,id)
}

func UpdatePost(user_id int, id int, is_anonymous int, is_urgent int, post_type int, title string, content string) error {
	return d.UpdatePost(ctx, user_id, id, is_anonymous, is_urgent, post_type, title, content)
}


//删除反馈
func DeletePost(user_id int, post_id int) error {
	return d.DeletePost(ctx, user_id, post_id)
}


//获取反馈列表
func GetPostList()([]model.Post,error){
	return d.GetPostList(ctx)
}

//查看回复
func GetResponse(user_id int)([]model.Response,error){
	return d.GetResponse(ctx,user_id)
}

func GetResponseByPostID(id int)(*model.Response,error){
	return d.GetResponseByPostID(ctx,id)
}