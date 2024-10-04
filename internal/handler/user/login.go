package user

import (
	"StudentServicePlatform/internal/apiException"
	// "StudentServicePlatform/internal/model"
	// "StudentServicePlatform/internal/model"
	"StudentServicePlatform/internal/service"
	"StudentServicePlatform/pkg/utils"

	"github.com/gin-gonic/gin"
)

type LoginData struct {
	Username int    `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func Login(c *gin.Context) {
	var data LoginData
	if err := c.ShouldBindJSON(&data); err != nil {
		_ = c.AbortWithError(200, apiException.ParamError) //参数错误
		return
	}
	_, err := service.GetUserByUsername(data.Username)
	if err != nil {
		_ = c.AbortWithError(200, apiException.UserNotFind) //用户不存在
		return
	}
	user, _ := service.GetUserPassword(data.Username)
	err = utils.CheckPassword(data.Password, user.Password)
	if err != nil {
		_ = c.AbortWithError(200, apiException.NoThatPasswordOrWrong)
		return
	}
	utils.JsonSuccess(c, user)
	// utils.JsonSuccess(c, gin.H{
	// 	"user_id":   user.ID,
	// 	"name":      user.Name,
	// 	"user_type": user.UserType,
	// })
}
