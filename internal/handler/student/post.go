package student

import (
	"StudentServicePlatform/internal/apiException"
	"StudentServicePlatform/internal/service"
	"StudentServicePlatform/pkg/utils"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type CreatePostData struct {
	UserID      int    `json:"user_id"`
	Name        string `json:"name"`
	IsAnonymous int    `json:"is_anonymous"`
	IsUrgent    int    `json:"is_urgent"`
	PostType    int    `json:"post_type"`
	Title       string `json:"title"`
	Content     string `json:"content"`
}

func CreatePost(c *gin.Context) {
	var data CreatePostData
	err := c.ShouldBindJSON(&data)
	if err != nil {
		_ = c.AbortWithError(200, apiException.ParamError) //参数错误
		return
	}
	_, err = service.GetUserByUserID(data.UserID)
	if err != nil {
		_ = c.AbortWithError(200, apiException.UserNotFind) //用户不存在
		return
	}
	if data.PostType != 1 && data.PostType != 2 && data.PostType != 3 && data.PostType != 4 {
		_ = c.AbortWithError(200, apiException.PostTypeError) //反馈类型无效
		return
	}
	err = service.CreatePost(data.UserID, data.Name, data.IsAnonymous, data.IsUrgent, data.PostType, data.Title, data.Content)
	if err != nil {
		_ = c.AbortWithError(200, apiException.CreatePostError) //提交反馈失败
		return
	}
	utils.JsonSuccess(c, nil)
}

type UpdatePostData struct {
	UserID      int    `json:"user_id"`
	ID          int    `json:"post_id"`
	IsAnonymous int    `json:"is_anonymous"`
	IsUrgent    int    `json:"is_urgent"`
	PostType    int    `json:"post_type"`
	Title       string `json:"title"`
	Content     string `json:"content"`
}

func UpdatePost(c *gin.Context) {
	var data UpdatePostData
	err := c.ShouldBindJSON(&data)
	if err != nil {
		_ = c.AbortWithError(200, apiException.ParamError) //参数错误
		return
	}
	_, err = service.GetUserByUserID(data.UserID)
	if err != nil {
		_ = c.AbortWithError(200, apiException.UserNotFind) //用户不存在
		return
	}
	post, err := service.GetPostByID(data.ID)
	if err != nil {
		_ = c.AbortWithError(200, apiException.PostNotFind) //反馈不存在
		return
	}
	if post.UserID != data.UserID {
		_ = c.AbortWithError(200, apiException.UserConnotUpdatePost) //无权修改帖子
		return
	}
	if data.PostType != 1 && data.PostType != 2 && data.PostType != 3 && data.PostType != 4 {
		_ = c.AbortWithError(200, apiException.PostTypeError) //反馈类型无效
		return
	}
	err = service.UpdatePost(data.UserID, data.ID, data.IsAnonymous, data.IsUrgent, data.PostType, data.Title, data.Content)
	if err != nil {
		_ = c.AbortWithError(200, apiException.UpdatePostError) //修改反馈失败
		return
	}
	utils.JsonSuccess(c, nil)
}

type DeletePostData struct {
	UserID int `json:"user_id" binding:"required"`
	ID     int `json:"post_id" binding:"required"`
}

func DeletePost(c *gin.Context) {
	var data DeletePostData
	err := c.ShouldBindJSON(&data)
	if err != nil {
		_ = c.AbortWithError(200, apiException.ParamError) //参数错误
		return
	}
	_, err = service.GetUserByUserID(data.UserID)
	if err != nil {
		_ = c.AbortWithError(200, apiException.UserNotFind) //用户不存在
		return
	}
	post, err := service.GetPostByID(data.ID)
	if err != nil {
		_ = c.AbortWithError(200, apiException.PostNotFind) //反馈不存在
		return
	}
	if post.UserID != data.UserID {
		_ = c.AbortWithError(200, apiException.UserConnotDeletePost) //无权删除帖子
		return
	}
	err = service.DeletePost(data.UserID, data.ID)
	if err != nil {
		_ = c.AbortWithError(200, apiException.DeletePostError) //删除反馈失败
		return
	}
	utils.JsonSuccess(c, nil)
}

type PostResponse struct {
	ID             int       `json:"post_id"`
	UserID         int       `json:"user_id"`
	Name           string    `json:"name"`
	IsAnonymous    int       `json:"is_anonymous"`
	IsUrgent       int       `json:"is_urgent"`
	PostType       int       `json:"post_type"`
	Title          string    `json:"title"`
	Content        string    `json:"content"`
	Response       string    `json:"response"`
	ResponseRating int       `json:"response_rating"`
	Status         int       `json:"status"`
	CreateAt       time.Time `json:"post_time"`
	ResponseAt     time.Time `json:"response_time"`
}

func GetPostList(c *gin.Context) {

	postList, err := service.GetPostList()
	if err != nil {
		_ = c.AbortWithError(http.StatusOK, apiException.GetPostListError) //获取反馈列表失败
		return
	}

	postResponseList := make([]PostResponse, len(postList))
	for index, post := range postList {
		response, err := service.GetResponseByPostID(post.ID)
		if err != nil {
			_ = c.AbortWithError(http.StatusOK, apiException.ServerError) //系统异常，请稍后重试!
			return
		}
		postResponseList[index] = PostResponse{
			ID:             post.ID,
			UserID:         post.UserID,
			Name:           post.Name,
			IsAnonymous:    post.IsAnonymous,
			IsUrgent:       post.IsUrgent,
			PostType:       post.PostType,
			Title:          post.Title,
			Content:        post.Content,
			Response:       response.Response,
			ResponseRating: response.ResponseRating,
			Status:         post.Status,
			CreateAt:       post.CreateAt,
			ResponseAt:     response.CreateAt,
		}

	}
	//fmt.Println(postResponseList)
	utils.JsonSuccess(c, postResponseList)
}
