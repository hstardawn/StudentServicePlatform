package router

import (
	"StudentServicePlatform/internal/handler/admin"
	"StudentServicePlatform/internal/handler/student"
	"StudentServicePlatform/internal/handler/upload"
	"StudentServicePlatform/internal/handler/user"
	"StudentServicePlatform/internal/middleware"

	// "StudentServicePlatform/internal/middleware"

	"github.com/gin-gonic/gin"
)

func Init(r *gin.Engine){
	const pre = "/api"
	api:=r.Group(pre)
	{	
		uploadGroup:= api.Group("/upload")
		{
			uploadGroup.Use(middleware.IsLogin)
			uploadGroup.POST("/post_image",upload.UploadPostImage)
			uploadGroup.POST("/user_image",upload.UploadUserImage)
		}
		userGroup := api.Group("/user")
		{
			userGroup.Use(middleware.IsLogin)
			userGroup.POST("/register",user.Register)
			userGroup.POST("/login",user.Login)
			userGroup.PUT("/update",user.UpdateUser)
			userGroup.POST("/send_code", user.SendCode)
		}

		studentGroup := api.Group("/student")
		{
			studentGroup.Use(middleware.IsLogin)
			studentGroup.POST("/post",student.CreatePost)
			studentGroup.PUT("/post",student.UpdatePost)
			studentGroup.DELETE("/post",student.DeletePost)
			studentGroup.GET("/post",student.GetPostList)
			studentGroup.GET("/response",student.GetResponse)
			studentGroup.POST("/response",student.CreateResponseRating)
		}

		adminGroup := api.Group("/admin")
		{
			
			adminGroup.GET("/handlepost",middleware.IsAdmin, admin.QueryUnhandlePost)
			adminGroup.POST("/handlepost", middleware.IsAdmin,admin.ReceivePost)
			adminGroup.DELETE("/handlepost", middleware.IsAdmin,admin.QuashHandle)
			adminGroup.PUT("/handlepost", middleware.IsAdmin,admin.ChangeResponse)
			adminGroup.GET("/handlepost/getall", admin.GetAllPost)

			adminGroup.GET("/superadmin", middleware.IsSU,admin.GetTrash)
			adminGroup.POST("/superadmin", middleware.IsSU,admin.HandleTrash)
			adminGroup.GET("/superadmin/queryadmin", middleware.IsSU,admin.QueryAdmin)
			adminGroup.PUT("/superadmin", middleware.IsSU,admin.UpdateAdmin)
		}
	}
}