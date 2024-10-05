package admin

import (
	"StudentServicePlatform/internal/apiException"
	"StudentServicePlatform/internal/service"
	"StudentServicePlatform/pkg/utils"

	"github.com/gin-gonic/gin"
)

type queryAdmin struct {
	AdminID int `form:"admin_id" binding:"required"`
}

type GetAdmin struct {
	ID       int    `json:"user_id"`
	Username int    `json:"username"`
	Name     string `json:"name"`
	Sex      string `json:"sex"`
	PhoneNum int    `json:"phone_num"`
	Email    string `json:"email"`
	UserType int    `json:"user_type"`
}

func QueryAdmin(c *gin.Context) {
	var data queryAdmin
	err := c.ShouldBindQuery(&data)
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

	// 检验用户权限
	if user.UserType != 2 {
		_ = c.AbortWithError(200, apiException.LackRight)
		return
	}

	// 获取管理员
	adminList, err := service.QueryAdmin()
	if err != nil {
		_ = c.AbortWithError(200, apiException.GetAdminListError)
		return
	}
	var admin_list []GetAdmin
	for _, admin := range adminList {
		// 2.获取帖子内容
		admin, err := service.GetUserByUserID(admin.ID)
		if err != nil {
			_ = c.AbortWithError(200, apiException.GetUserError)
			return
		}
		
		// 3.返回帖子内容
		admin_list = append(admin_list, GetAdmin{
			ID:       admin.ID,
			Username: admin.Username,
			Name:     admin.Name,
			Sex:      admin.Sex,
			PhoneNum: admin.PhoneNum,
			Email:    admin.Email,
			UserType: admin.UserType,
		})
	}

	utils.JsonSuccess(c, gin.H{
		"admin_list": admin_list,
	})
}
