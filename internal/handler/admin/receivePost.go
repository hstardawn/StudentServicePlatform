package admin

import (
	"StudentServicePlatform/internal/apiException"
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
		_ = c.AbortWithError(200, apiException.ParamError)
		return
	}

	// 检验用户存在
	user, err := service.GetUserByUserID(data.AdminID)
	if err != nil {
		_ = c.AbortWithError(200, apiException.AdminNotFind)
		return
	}
	
	//检验反馈存在
	post, err := service.GetPostByID(data.PostID)
	if err != nil {
		_ =c.AbortWithError(200, apiException.PostNotFind)
		return
	}

	// 检验用户权限
	if user.UserType==3 {
		_ =c.AbortWithError(200, apiException.LackRight)
		return
	}

	// 检查反馈状态
	if post.Status != 0{
		_ = c.AbortWithError(200, apiException.ReatHandle)
		return
	}

	//接单
	if data.Approval == 2{
		err := service.UpdatePostStatus(data.AdminID,data.PostID,2)
		if err != nil {
			_ = c.AbortWithError(200, apiException.UpdatePostError)
			return
		}
	}
	err = service.ReceivePost(data.AdminID, data.PostID, data.Response)
	if err!=nil{
		_ = c.AbortWithError(200, apiException.HandleError)
		return
	}

	err = service.UpdatePostStatus(data.AdminID,data.PostID, data.Approval)
	if err != nil{
		_ = c.AbortWithError(200, apiException.SaveError)
		return
	}

	utils.JsonSuccess(c, nil)
}