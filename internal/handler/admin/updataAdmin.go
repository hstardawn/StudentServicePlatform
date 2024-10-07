package admin

import (
	"StudentServicePlatform/internal/apiException"
	"StudentServicePlatform/internal/service"
	"StudentServicePlatform/pkg/utils"

	"github.com/gin-gonic/gin"
)

type updateAdmin struct {
	AdminID  int `json:"admin_id" binding:"required"`
	UserID   int `json:"user_id" binding:"required"`
	UserType int `json:"user_type"`
}

func UpdateAdmin(c *gin.Context){
	var data updateAdmin
	err := c.ShouldBindJSON(&data)
	if err != nil {
		_ = c.AbortWithError(200, apiException.ParamError)
		return
	}

	// 检验用户存在
	user , err := service.GetUserByUserID(data.UserID)
	if err != nil {
		_ = c.AbortWithError(200, apiException.UserNotFind)
		return
	}

	// bug
	if data.AdminID == data.UserID{
		_ = c.AbortWithError(200 , apiException.DouM)
	}
	
	// 检验用户身份
	if user.UserType == 2 {
		_ = c.AbortWithError(200, apiException.Wantdie)
		return
	}
	// 检验管理员
	admin, err := service.GetUserByUserID(data.AdminID)
	if err != nil{
		_ = c.AbortWithError(200, apiException.AdminNotFind)
		return
	}

	// 检验管理员权限
	if admin.UserType != 2{
		_ = c.AbortWithError(200, apiException.LackRight)
		return
	}

	// 更改权限
	if data.UserType ==2 {
		_ = c.AbortWithError(200 ,apiException.LackRight)
		return
	}
	err = service.UpdateUserType(data.UserID, data.UserType)
	if err != nil {
		_ = c.AbortWithError(200, apiException.UpdateRightError)
		return
	}
	utils.JsonSuccess(c,nil)
}