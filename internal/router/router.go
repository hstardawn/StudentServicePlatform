package router

import (
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
	}
}