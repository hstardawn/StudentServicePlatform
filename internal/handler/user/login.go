package user

import (
	"StudentServicePlatform/internal/apiException"
	"StudentServicePlatform/internal/service"
	"StudentServicePlatform/pkg/utils"
	"encoding/json"

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
	user, err := service.GetUserByUsername(data.Username)
	if err != nil {
		_ = c.AbortWithError(200, apiException.UserNotFind) //用户不存在
		return
	}
	// 生成token
	token, err := utils.GenerateJWT(user.ID)
	if err != nil {
		_ = c.AbortWithError(200, apiException.ServerError)
		return
	}
	user, _ = service.GetUserPassword(data.Username)
	hashpassword , err:= utils.HashPassword(data.Password)
	if err != nil{
		_ = c.AbortWithError(200, apiException.EncryptionFailed)
		return
	}
	err = utils.CheckPassword(hashpassword, data.Password)
	if err != nil {
		_ = c.AbortWithError(200, apiException.NoThatPasswordOrWrong)
		return
	}
	//解析图片
	var pictures []string
	picturesBytes := []byte(user.Pictures)
	_ = json.Unmarshal(picturesBytes, &pictures)
	utils.JsonSuccess(c, gin.H{
		"token": token,
		"user":  user,
	})
}
