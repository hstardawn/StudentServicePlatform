package user

import (
	"StudentServicePlatform/internal/apiException"
	"StudentServicePlatform/internal/service"
	"StudentServicePlatform/pkg/utils"

	"github.com/gin-gonic/gin"
)

type UpdateUserData struct {
	Username int    `json:"username" binding:"required"`
	Name     string `json:"name" binding:"required"`
	Sex      string `json:"sex" bingding:"required"`
	PhoneNum int    `json:"phone_num" binding:"required"`
	Email    string `json:"email" binding:"required"`
	// UserType int    `json:"user_type" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func UpdateUser(c *gin.Context) {
	var data UpdateUserData
	err := c.ShouldBindJSON(&data)
	if err != nil {
		_ = c.AbortWithError(200, apiException.ParamError) //参数错误
		return
	}
	_, err = service.GetUserByUsername(data.Username)
	if err != nil {
		_ = c.AbortWithError(200, apiException.UserNotFind) //用户不存在
		return
	}
	if len(data.Password)<8||len(data.Password)>16{
		_ = c.AbortWithError(200, apiException.PasswordError) //密码长度必须大于8且小于16位
		return
	}
	// if data.UserType!=3&&data.UserType!=1&&data.UserType!=2{
	// 	_ = c.AbortWithError(200, apiException.UserTypeError) //用户类型无效
	// 	return
	// }
	user, _ := service.GetUserPassword(data.Username)
	if user.Password != data.Password {
		_ = c.AbortWithError(200, apiException.NoThatPasswordOrWrong) //密码错误
		return
	}
	err = service.UpdateUser(data.Username, data.Name, data.Sex, data.PhoneNum, data.Email, data.Password)
	if err != nil {
		_ = c.AbortWithError(200, apiException.UpdateUserError) //修改用户信息失败
		return
	}
	utils.JsonSuccess(c, nil)
}