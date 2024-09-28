package admin

import (
	apiexception "StudentServicePlatform/internal/apiException"
	"StudentServicePlatform/internal/service"
	"StudentServicePlatform/pkg/utils"

	"github.com/gin-gonic/gin"
)

type receivePost struct {
	AdminID  int    `json:"admin_id" binding:"required"`
	PostID   int    `json:"post_id" binding:"required"`
	Approval int    `json:"approval" binding:"required"`
	Response string `json:"response" binding:"required"`
}

func ReceivePost(c *gin.Context){
	var data receivePost
	err := c.ShouldBindJSON(&data)
	if err != nil {
		_ = c.AbortWithError(200, apiexception.ParamError)
		return
	}

	// 检验用户存在
	user, err := service.GetUserByUserid(data.AdminID)
	if err != nil {
		_ = c.AbortWithError(200, apiexception.UserNotFound)
		return
	}
	
	//检验反馈存在
	post, err := service.GetPostByPostId(data.PostID)
	if err != nil {
		_ =c.AbortWithError(200, apiexception.PostNotFound)
		return
	}

	// 检验用户权限
	if user.UserType==3 {
		_ =c.AbortWithError(200, apiexception.NotAdmin)
	}

	// 检查反馈状态
	if post.Status != 0{
		_ = c.AbortWithError(200, apiexception.ReatHandle)
	}

	//接单
	post.Status = data.Approval
	post.Response = data.Response
	err = service.SavePost(*post)
	if err!=nil{
		_ = c.AbortWithError(200, apiexception.SaveError)
		return
	}

	utils.JsonSuccess(c, nil)
}