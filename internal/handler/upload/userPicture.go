package upload

import (
	"StudentServicePlatform/internal/apiException"
	"StudentServicePlatform/internal/global"
	"StudentServicePlatform/internal/service"
	"StudentServicePlatform/pkg/utils"
	"os"

	"github.com/gin-gonic/gin"
)

type UploadUserPictureData struct {
	UserID   int    `form:"user_id" binding:"required"`
	Filename string `json:"filename"`
}

func UploadUserImage(c *gin.Context) {
	var data UploadUserPictureData
	err := c.ShouldBind(&data)
	if err != nil {
		_ = c.AbortWithError(200, apiException.ParamError) //参数错误
		return
	}
	user, err := service.GetUserByUserID(data.UserID)
	if err != nil {
		_ = c.AbortWithError(200, apiException.UserNotFind) //用户不存在
		return
	}

	//判断文件是否为图片
	wrong:=service.IsImage(data.Filename)
	if wrong {
		_ = c.AbortWithError(200, apiException.FileTypeError)//文件不是图片类型
		return
	}

	// 解析表单数据
	err = c.Request.ParseForm() 
	if err != nil {
		_ = c.AbortWithError(200, apiException.ParseFormDataError)//解析表单数据失败
		return
	}

	//获取上传文件
	file, err := c.FormFile("image")
	if err != nil {
		_ = c.AbortWithError(200, apiException.GetFileError)//获取文件失败
		return
	}

	//保存文件
	picType := service.GetFileType(file.Filename)
	dst := global.Config.GetString("file.imagePath") + "/" + service.GetUUID() + picType
	_, err = os.Stat(dst);
	if err == nil {
		_ = c.AbortWithError(200, apiException.FileExistedError)
		return
	}
	err = c.SaveUploadedFile(file, "."+dst);
	if err != nil {
		_ = c.AbortWithError(200, apiException.ServerError)
		return
	}

	//存储到数据库
	err = service.StoreUserPicture(user.ID,dst)
	if err!= nil {
		_ = c.AbortWithError(200, apiException.StorePictureError)//存储图片失败
		return
	}
	utils.JsonSuccess(c, gin.H{"picture_url": dst})
}