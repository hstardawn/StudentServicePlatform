package router

import (
	"StudentServicePlatform/internal/handler/student"
	"StudentServicePlatform/internal/handler/user"

	"github.com/gin-gonic/gin"
)

func Init(r *gin.Engine){
	const pre = "/api"
	api:=r.Group(pre)
	{
		api.POST("/user/register",user.Register)
		api.POST("/user/login",user.Login)
		api.PUT("/user/update",user.UpdateUser)

		api.POST("/student/post",student.CreatePost)
		api.PUT("/student/post",student.UpdatePost)
		api.DELETE("/student/post",student.DeletePost)
		api.GET("/student/post",student.GetPostList)

		api.GET("/student/response",student.GetResponse)
		api.POST("/student/response",student.CreateResponseRating)
	}
}