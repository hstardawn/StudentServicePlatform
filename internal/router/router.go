package router

import (
	"StudentServicePlatform/internal/handler/admin"

	"github.com/gin-gonic/gin"
)

func Init(r *gin.Engine){
	const pre = "/api"
	api:=r.Group(pre)
	{
		adminGroup := api.Group("/admin")
		{
			adminGroup.GET("handlepost", admin.QueryUnhandlePost)
			adminGroup.POST("handlepost", admin.ReceivePost)
			adminGroup.DELETE("handlepost", admin.QuashHandle)
			adminGroup.PUT("handlepost", admin.ChangeResonse)
		}
	}
}