package admin

import (
	"StudentServicePlatform/internal/apiException"
	"StudentServicePlatform/internal/service"
	"StudentServicePlatform/pkg/utils"

	"github.com/gin-gonic/gin"
)

type quashHandle struct {
	AdminID int `form:"admin_id" binding:"required"`
	PostID  int `form:"post_id" binding:"required"`
}

func QuashHandle(c *gin.Context) {
	var data quashHandle
	err := c.ShouldBind(&data)
	if  err != nil {
		_ = c.AbortWithError(200, apiException.ParamError)
		return
	}
	// 检验用户存在
	_ , err = service.GetUserByUserid(data.AdminID)
	if err != nil {
		_ = c.AbortWithError(200, apiException.UserNotFound)
		return
	}

	// 检验反馈存在
	post, err := service.GetPostByPostId(data.PostID)
	if err != nil {
		_ = c.AbortWithError(200, apiException.PostNotFound)
		return
	}

	// 检验是否同一人处理
	if post.AdminID != data.AdminID{
		_ = c.AbortWithError(200, apiException.AdminUncompaired)
		return
	}

	// 撤销
	post.Status = 0
	post.AdminID = 0
	post.Response = ""
	err = service.QuashPost(post)
	if err!=nil{
		_ = c.AbortWithError(200, apiException.SaveError)
		return
	}

	utils.JsonSuccess(c, nil)
}
