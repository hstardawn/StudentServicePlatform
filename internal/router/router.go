package router

import (
	"StudentServicePlatform/internal/handler/admin"
	"StudentServicePlatform/internal/handler/student"
	"StudentServicePlatform/internal/handler/upload"
	"StudentServicePlatform/internal/handler/user"
	// "StudentServicePlatform/internal/middleware"

	"github.com/gin-gonic/gin"
)

func Init(r *gin.Engine){
	const pre = "/api"
	api:=r.Group(pre)
	{	
		uploadGroup:= api.Group("/upload")
		{
			// uploadGroup.Use(middleware.IsLogin)
			uploadGroup.POST("/post_image",upload.UploadPostImage)
			uploadGroup.POST("/user_image",upload.UploadUserImage)
		}
		userGroup := api.Group("/user")
		{
			userGroup.POST("/register",user.Register)
			userGroup.POST("/login",user.Login)
			userGroup.PUT("/update",user.UpdateUser)
			userGroup.POST("/send_code", user.SendCode)
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
			adminGroup.GET("/handlepost/getall", admin.GetAllPost)

			adminGroup.GET("/superadmin", admin.GetTrash)
			adminGroup.POST("/superadmin", admin.HandleTrash)
			adminGroup.GET("/superadmin/queryadmin", admin.QueryAdmin)
			adminGroup.PUT("/superadmin", admin.UpdateAdmin)
		}
	}
}