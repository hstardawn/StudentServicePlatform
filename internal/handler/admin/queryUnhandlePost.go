package admin

import (
	"StudentServicePlatform/internal/apiException"
	"StudentServicePlatform/internal/service"
	"StudentServicePlatform/pkg/utils"

	"github.com/gin-gonic/gin"
)

type queryUnhandlePost struct {
	AdminID int `form:"admin_id" binding:"required"`
}

func QueryUnhandlePost(c *gin.Context) {
	var data queryUnhandlePost
	err := c.ShouldBind(&data)
	if err != nil{
		_ = c.AbortWithError(200, apiException.ParamError)
		return
	}

	//检验用户存在
	user, err := service.GetUserByUserid(data.AdminID)
	if err != nil {
		_ = c.AbortWithError(200, apiException.UserNotFound)
		return
	}

	//检验用户权限
	if user.UserType==3 {
		_ = c.AbortWithError(200, apiException.NotAdmin)
		return
	}

	postList, err := service.QueryUnhandlePost()
	if err != nil {
		_ = c.AbortWithError(200, apiException.SearchError)
		return
	}
	utils.JsonSuccess(c, gin.H{
		"post_list": postList,
	})
}