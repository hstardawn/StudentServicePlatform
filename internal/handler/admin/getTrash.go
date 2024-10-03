package admin

import (
	"StudentServicePlatform/internal/apiException"
	"StudentServicePlatform/internal/service"
	"StudentServicePlatform/pkg/utils"

	"github.com/gin-gonic/gin"
)

type getTrash struct {
	AdminID int `form:"admin_id" binding:"required"`
}

func GetTrash(c *gin.Context) {
	var data getTrash
	err := c.ShouldBindQuery(&data)
	if  err != nil {
		_ = c.AbortWithError(200, apiException.ParamError)
		return
	}
	// 检验用户存在
	user, err := service.GetUserByUserID(data.AdminID)
	if err != nil {
		_ = c.AbortWithError(200, apiException.AdminNotFind)
		return
	}

	//检验用户权限
	if user.UserType != 2 {
		_ = c.AbortWithError(200, apiException.LackRight)
		return
	}

	// 获取垃圾信息
	trashList, err := service.QueryTrash()
	if err != nil {
		_ = c.AbortWithError(200, apiException.GetPostListError)
		return
	}
	utils.JsonSuccess(c, gin.H{
		"post_list": trashList,
	})
}