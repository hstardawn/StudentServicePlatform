package student

import (
	"StudentServicePlatform/internal/apiException"
	"StudentServicePlatform/internal/model"
	"StudentServicePlatform/internal/service"
	"StudentServicePlatform/pkg/utils"
	"time"

	"github.com/gin-gonic/gin"
)

// 查看回复
type GetResponseData struct {
	UserID int `form:"user_id"`
	// PostID int `form:"post_id"`
}

type Response struct {
	PostID   int       `json:"post_id"`
	Content  string    `json:"content"`
	Response string    `json:"response"`
	CreateAt time.Time `json:"response_time"`
}

func GetResponse(c *gin.Context) {
	var data GetResponseData
	err := c.ShouldBind(&data)
	if err != nil {
		_ = c.AbortWithError(200, apiException.ParamError) //参数错误
		return
	}
	_, err = service.GetUserByUserID(data.UserID)
	if err != nil {
		_ = c.AbortWithError(200, apiException.UserNotFind) //用户不存在
		return
	}
	Posts, err := service.GetPostByUserID(data.UserID)
	//fmt.Println(Posts)
	if err != nil {
		_ = c.AbortWithError(200, apiException.GetPostError) //用户没有提出反馈
		return
	}
	var response []model.Response
	var ResponseList []Response
	for _,Posts := range Posts {
		response, err = service.GetResponse(Posts.ID)
	    if err != nil {
		    _ = c.AbortWithError(200, apiException.GetResponseError) //查看回复失败
		    return
	    }
	    for _, response := range response {
		    post, err := service.GetPostByID(response.PostID)
		    if err!= nil {
			    _ = c.AbortWithError(200, apiException.GetPostError) //用户没有提出反馈
			    return
		    }
		    ResponseList = append(ResponseList, Response{
			    PostID:   response.PostID,
			    Content:  post.Content,
			    Response: response.Response,
			    CreateAt: response.CreateAt,
		    })
	    }
	}
	utils.JsonSuccess(c, gin.H{
		"response": ResponseList,
	})
}

type CreateResponseRatingData struct {
	UserID         int `form:"user_id"`
	PostID         int `form:"post_id"`
	ResponseRating int `form:"response_rating"`
}

func CreateResponseRating(c *gin.Context) {
	var data CreateResponseRatingData
	err := c.ShouldBind(&data)
	if err != nil {
		_ = c.AbortWithError(200, apiException.ParamError) //参数错误
		return
	}
	_, err = service.GetUserByUserID(data.UserID)
	if err != nil {
		_ = c.AbortWithError(200, apiException.UserNotFind) //用户不存在
		return
	}
	post, err := service.GetPostByID(data.PostID)
	if err != nil {
		_ = c.AbortWithError(200, apiException.PostNotFind) //反馈不存在
		return
	}
	if post.Status == 2 {
		_ = c.AbortWithError(200, apiException.TrashPost) //反馈被标记为垃圾信息
		return
	}
	if post.UserID != data.UserID {
		_ = c.AbortWithError(200, apiException.UserConnotRateResponse) //无权做出评价
		return
	}
	if data.ResponseRating != 1 && data.ResponseRating != 2 && data.ResponseRating != 3 && data.ResponseRating != 4 {
		_ = c.AbortWithError(200, apiException.ResponseRatingError) //评价类型无效
		return
	}
	_, err = service.GetResponse(data.UserID)
	if err != nil {
		_ = c.AbortWithError(200, apiException.CreateResponseRatingError) //查看回复失败
		return
	}
	utils.JsonSuccess(c, nil)
}
