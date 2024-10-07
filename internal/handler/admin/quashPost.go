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
	err := c.ShouldBindQuery(&data)
	if  err != nil {
		_ = c.AbortWithError(200, apiException.ParamError)
		return
	}
	// 检验用户存在
	_ , err = service.GetUserByUserID(data.AdminID)
	if err != nil {
		_ = c.AbortWithError(200, apiException.AdminNotFind)
		return
	}

	// 检验反馈存在
	post, err := service.GetPostByID(data.PostID)
	if err != nil {
		_ = c.AbortWithError(200, apiException.PostNotFind)
		return
	}

	// 检验举报是否存在
	_, err = service.GetResponseByPID(data.PostID)
	if err != nil {
		_ = c.AbortWithError(200, apiException.ResponseNotExist)
		return
	}

	// 检验是否同一人处理
	if post.AdminID != data.AdminID{
		_ = c.AbortWithError(200, apiException.AdminUncompaired)
		return
	}

	// 撤销
	err = service.QuashPost(data.PostID)
	if err!=nil{
		_ = c.AbortWithError(200, apiException.SaveError)
		return
	}
	err = service.UpdatePostStatus(data.AdminID, data.PostID, 0)
	if err != nil{
		_ = c.AbortWithError(200, apiException.UpdatePostError)
		return 
	}

	utils.JsonSuccess(c, nil)
}
