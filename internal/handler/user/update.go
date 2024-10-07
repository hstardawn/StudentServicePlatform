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
	Sex      string `json:"sex" binding:"required"`
	PhoneNum int    `json:"phone_num" binding:"required"`
	Email    string `json:"email" binding:"required"`
	OriginPassword string `json:"origin_password" binding:"required"`
	NewPassword string `json:"new_password" binding:"required"`
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
	if len(data.NewPassword)<8||len(data.NewPassword)>16{
		_ = c.AbortWithError(200, apiException.PasswordError) //密码长度必须大于8且小于16位
		return
	}
	user, _ := service.GetUserPassword(data.Username)
	if user.Email!= data.Email {
		_ = c.AbortWithError(200, apiException.EmailError) //邮箱错误
		return
	}
	//加密
	_ = utils.CheckPassword(user.Password, data.OriginPassword)
	if user.Password != data.OriginPassword {
		_ = c.AbortWithError(200, apiException.NoThatPasswordOrWrong) //密码错误
		return
	}
	
	//解密
	hashpassword , err:= utils.HashPassword(data.NewPassword)
	if err != nil{
		_ = c.AbortWithError(200, apiException.EncryptionFailed)
		return
	}
	err = service.UpdateUser(data.Username, data.Name,data.Sex, data.PhoneNum,/* data.Email,*/ hashpassword)
	if err != nil {
		_ = c.AbortWithError(200, apiException.UpdateUserError) //修改用户信息失败
		return
	}
	utils.JsonSuccess(c, nil)
}