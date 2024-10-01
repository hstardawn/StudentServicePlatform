package apiException

import "net/http"

func (e *Error) Error() string {
	return e.Msg
}

type Error struct {
	StatusCode int    `json:"-"`
	Code       int    `json:"code"`
	Msg        string `json:"msg"`
}

func NewError(statusCode, Code int, msg string) *Error {
	return &Error{
		StatusCode: statusCode,
		Code:       Code,
		Msg:        msg,
	}
}

var (
	//注册
	ParamError         = NewError(http.StatusInternalServerError, 200501, "参数错误")
	UserAlreadyExisted = NewError(http.StatusInternalServerError, 200502, "用户名已存在")
	PasswordError      = NewError(http.StatusInternalServerError, 200503, "密码长度必须大于8且小于16位")
	UserTypeError      = NewError(http.StatusInternalServerError, 200504, "用户类型无效")
	Register           = NewError(http.StatusInternalServerError, 200505, "注册失败")

	//登录
	UserNotFind           = NewError(http.StatusInternalServerError, 200506, "该用户不存在")
	NoThatPasswordOrWrong = NewError(http.StatusInternalServerError, 200507, "密码错误")

	//修改用户信息
	UpdateUserError = NewError(http.StatusInternalServerError, 200508, "修改用户信息失败")

	//提交反馈
	PostTypeError   = NewError(http.StatusInternalServerError, 200509, "反馈类型无效")
	CreatePostError = NewError(http.StatusInternalServerError, 200510, "提交反馈失败")

	//修改反馈
	PostNotFind          = NewError(http.StatusInternalServerError, 200511, "反馈不存在")
	UserConnotUpdatePost = NewError(http.StatusInternalServerError, 200512, "用户无权修改反馈")
	UpdatePostError      = NewError(http.StatusInternalServerError, 200513, "修改反馈失败")

	//删除反馈
	UserConnotDeletePost = NewError(http.StatusInternalServerError, 200514, "用户无权删除反馈")
	DeletePostError      = NewError(http.StatusInternalServerError, 200515, "删除反馈失败")

	//获取反馈列表
	GetPostListError = NewError(http.StatusInternalServerError, 200516, "获取反馈列表失败")

	//查看回复
	GetResponseError = NewError(http.StatusInternalServerError, 200517, "查看回复失败")

	//做出评价
	UserConnotRateResponse    = NewError(http.StatusInternalServerError, 200518, "用户无权做出评价")
	CreateResponseRatingError = NewError(http.StatusInternalServerError, 200519, "做出评价失败")
)
