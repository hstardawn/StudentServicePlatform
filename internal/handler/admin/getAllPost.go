package admin

import (
	"StudentServicePlatform/internal/apiException"
	"StudentServicePlatform/internal/service"
	"StudentServicePlatform/pkg/utils"

	"github.com/gin-gonic/gin"
)

type getAllPost struct {
	AdminID int `form:"admin_id" binding:"required"`
}

func GetAllPost(c *gin.Context) {
	var data getAllPost
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
	if user.UserType == 0 {
		_ = c.AbortWithError(200, apiException.LackRight)
		return
	}

	// 获取已处理的帖子
	postList, err := service.GetPostByAdminID(data.AdminID)
	if err != nil {
		_ = c.AbortWithError(200, apiException.GetPostListError)
	}
	utils.JsonSuccess(c,postList)

}