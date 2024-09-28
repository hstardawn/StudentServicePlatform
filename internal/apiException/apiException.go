package apiexception

import "net/http"

type Error struct {
	StatusCode int    `json:"-"`
	Code       int    `json:"code"`
	Msg        string `json:"msg"`
}

var (
	ParamError = NewError(http.StatusInternalServerError, 200500, "参数错误")
	UserNotFound = NewError(http.StatusInternalServerError, 200501, "用户不存在" )
	NotAdmin =NewError(http.StatusInternalServerError, 200502, "权限不足")
	SearchError =NewError(http.StatusInternalServerError, 200503, "查询失败")
	PostNotFound =NewError(http.StatusInternalServerError, 200504, "反馈不存在")
	ReatHandle =NewError(http.StatusInternalServerError, 200505, "重复接单")
	SaveError =NewError(http.StatusInternalServerError, 200506, "保存失败")
	AdminUncompaired =NewError(http.StatusInternalServerError, 200507, "处理人不一致")
)
func OtherError(message string) *Error {
	return NewError(http.StatusForbidden, 100403, message)
}

func (e *Error) Error() string {
	return e.Msg
}

func NewError(statusCode, Code int, msg string) *Error {
	return &Error{
		StatusCode: statusCode,
		Code:       Code,
		Msg:        msg,
	}
}