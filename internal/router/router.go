package router

import (
	"StudentServicePlatform/internal/handler/student"
	"StudentServicePlatform/internal/handler/user"
	"StudentServicePlatform/internal/handler/admin"

	"github.com/gin-gonic/gin"
)

func Init(r *gin.Engine){
	const pre = "/api"
	api:=r.Group(pre)
	{	
		userGroup := api.Group("/user")
		{
			userGroup.POST("/user/register",user.Register)
			userGroup.POST("/user/login",user.Login)
			userGroup.PUT("/user/update",user.UpdateUser)
		}

		studentGroup := api.Group("/student")
		{
			studentGroup.POST("/student/post",student.CreatePost)
			studentGroup.PUT("/student/post",student.UpdatePost)
			studentGroup.DELETE("/student/post",student.DeletePost)
			studentGroup.GET("/student/post",student.GetPostList)
			studentGroup.GET("/student/response",student.GetResponse)
			studentGroup.POST("/student/response",student.CreateResponseRating)
		}

			adminGroup := api.Group("/admin")
		{
			adminGroup.GET("handlepost", admin.QueryUnhandlePost)
			adminGroup.POST("handlepost", admin.ReceivePost)
			adminGroup.DELETE("handlepost", admin.QuashHandle)
			adminGroup.PUT("handlepost", admin.ChangeResonse)
		}
	}
}