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
			userGroup.POST("/register",user.Register)
			userGroup.POST("/login",user.Login)
			userGroup.PUT("/update",user.UpdateUser)
		}

		studentGroup := api.Group("/student")
		{
			studentGroup.POST("/post",student.CreatePost)
			studentGroup.PUT("/post",student.UpdatePost)
			studentGroup.DELETE("/post",student.DeletePost)
			studentGroup.GET("/post",student.GetPostList)
			studentGroup.GET("/response",student.GetResponse)
			studentGroup.POST("/response",student.CreateResponseRating)
		}

		adminGroup := api.Group("/admin")
		{
			adminGroup.GET("/handlepost", admin.QueryUnhandlePost)
			adminGroup.POST("/handlepost", admin.ReceivePost)
			adminGroup.DELETE("/handlepost", admin.QuashHandle)
			adminGroup.PUT("/handlepost", admin.ChangeResponse)

			adminGroup.GET("/superadmin", admin.GetTrash)
			adminGroup.POST("/superadmin", admin.HandleTrash)
			adminGroup.GET("/superadmin/queryadmin", admin.QueryAdmin)
			adminGroup.PUT("/superadmin", admin.UpdateAdmin)
		}
	}
}