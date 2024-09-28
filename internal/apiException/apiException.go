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
	ParamError         = NewError(http.StatusInternalServerError, 200501, "参数错误")
	UserAlreadyExisted = NewError(http.StatusInternalServerError, 200502, "用户名已存在")
	PasswordError      = NewError(http.StatusInternalServerError, 200503, "密码长度必须大于8且小于16位")
	UserTypeError      = NewError(http.StatusInternalServerError, 200504, "用户类型无效")
	Register           = NewError(http.StatusInternalServerError, 200505, "注册失败")

	UserNotFind           = NewError(http.StatusInternalServerError, 200506, "该用户不存在")
	NoThatPasswordOrWrong = NewError(http.StatusInternalServerError, 200507, "密码错误")

	UpdateUserError = NewError(http.StatusInternalServerError, 200508, "修改用户信息失败")
)
