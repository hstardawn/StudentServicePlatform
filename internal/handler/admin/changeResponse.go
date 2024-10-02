package admin

import (
	"StudentServicePlatform/internal/apiException"
	"StudentServicePlatform/internal/service"
	"StudentServicePlatform/pkg/utils"

	"github.com/gin-gonic/gin"
)

type changeResonse struct {
	AdminID  int    `json:"admin_id" binding:"required"`
	PostID   int    `json:"post_id" binding:"required"`
	Response string `json:"resonse" binding:"required"`
}

func ChangeResonse(c *gin.Context){
	var data changeResonse
	err := c.ShouldBindJSON(&data)
	if err != nil {
		_ = c.AbortWithError(200, apiException.ParamError)
		return
	}

	// 检验用户存在
	_, err = service.GetUserByUserID(data.AdminID)
	if err != nil {
		_ = c.AbortWithError(200, apiException.UserNotFind)
		return
	}
	
	//检验反馈存在
	post, err := service.GetPostByID(data.PostID)
	if err != nil {
		_ = c.AbortWithError(200, apiException.PostNotFind)
		return
	}

	// 检验用户权限
	if post.AdminID != data.AdminID{
		_ = c.AbortWithError(200, apiException.AdminUncompaired)
	}

	err = service.ChangeResonse(data.PostID, data.Response)
	if err != nil{
		_ = c.AbortWithError(200, apiException.SaveError)
		return
	}
	utils.JsonSuccess(c, nil)
}