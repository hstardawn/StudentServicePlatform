package middleware

import (
	"StudentServicePlatform/internal/apiException"
	// "StudentServicePlatform/internal/model"
	"StudentServicePlatform/internal/service"
	"StudentServicePlatform/pkg/utils"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func IsLogin(c *gin.Context) {
	// 1. 从请求头中获取token
	tokenStr := c.Request.Header.Get("Authorization")
	if len(tokenStr) <= 7 {
		_ = c.AbortWithError(http.StatusOK, apiException.AuthExpired)
		return
	}
	tokenStr = tokenStr[7:]
	//utils.Log.Println(tokenStr)
	// 2. 解析token
	jwtUser, err := utils.ParseJwt(tokenStr)
	if err != nil {
		_ = c.AbortWithError(http.StatusOK, apiException.AuthExpired)
		return
	}

	// 3. 判断是否过期
	if time.Now().Unix() > jwtUser.ExpiresAt.Unix() {
		_ = c.AbortWithError(http.StatusOK, apiException.AuthExpired)
		return
	}

	// 4. 获取用户信息
	user, err := service.GetUserByUserID(jwtUser.UserID)
	if err != nil {
		_ = c.AbortWithError(http.StatusOK, apiException.ServerError)
		return
	}
	c.Set("user", user)

}

// func IsAdmin(c *gin.Context) {
// 	IsLogin(c)
// 	user := c.MustGet("user").(*model.User)
// 	if !(user.UserType == 1 || user.UserType == 2) {
// 		_ = c.AbortWithError(http.StatusOK, apiException.PermissionsNotAllowed)
// 		return
// 	}
// }