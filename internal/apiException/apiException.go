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
	Register           = NewError(http.StatusInternalServerError, 200504, "注册失败")

	//登录
	UserNotFind           = NewError(http.StatusInternalServerError, 200505, "该用户不存在")
	NoThatPasswordOrWrong = NewError(http.StatusInternalServerError, 200506, "密码错误")

	//鉴权
	ServerError = NewError(http.StatusInternalServerError, 200507, "系统异常，请稍后重试!")
	AuthExpired = NewError(http.StatusInternalServerError, 200508, "登陆状态已过期，请重新登陆")

	//修改用户信息
	UpdateUserError = NewError(http.StatusInternalServerError, 200509, "修改用户信息失败")

	//提交反馈
	PostTypeError   = NewError(http.StatusInternalServerError, 200510, "反馈类型无效")
	CreatePostError = NewError(http.StatusInternalServerError, 200511, "提交反馈失败")

	//上传图片
	FileTypeError      = NewError(http.StatusInternalServerError, 200512, "文件不是图片类型")
	ParseFormDataError = NewError(http.StatusInternalServerError, 200513, "解析表单数据失败")
	FileExistedError    = NewError(http.StatusInternalServerError, 200514, "文件已存在")
	GetFileError       = NewError(http.StatusInternalServerError, 200514, "获取文件失败")

	//修改反馈
	PostNotFind          = NewError(http.StatusInternalServerError, 200512, "反馈不存在")
	UserConnotUpdatePost = NewError(http.StatusInternalServerError, 200513, "用户无权修改反馈")
	UpdatePostError      = NewError(http.StatusInternalServerError, 200514, "修改反馈失败")

	//删除反馈
	UserConnotDeletePost = NewError(http.StatusInternalServerError, 200515, "用户无权删除反馈")
	DeletePostError      = NewError(http.StatusInternalServerError, 200516, "删除反馈失败")

	//获取反馈列表
	GetPostListError = NewError(http.StatusInternalServerError, 200517, "获取反馈列表失败")

	//查看回复
	GetResponseError = NewError(http.StatusInternalServerError, 200518, "查看回复失败")

	//做出评价
	UserConnotRateResponse    = NewError(http.StatusInternalServerError, 200519, "用户无权做出评价")
	ResponseRatingError       = NewError(http.StatusInternalServerError, 200520, "评价类型无效")
	CreateResponseRatingError = NewError(http.StatusInternalServerError, 200521, "做出评价失败")

	LackRight =NewError(http.StatusInternalServerError, 200521, "权限不足")
	ReatHandle =NewError(http.StatusInternalServerError, 200522, "重复接单")
	SaveError =NewError(http.StatusInternalServerError, 200523, "保存失败")
	AdminUncompaired =NewError(http.StatusInternalServerError, 200524, "处理人不一致")
	HandleError =NewError(http.StatusInternalServerError, 200525, "处理失败")
	PostNotHandle =NewError(http.StatusInternalServerError, 200526, "帖子未处理")
	GetAdminListError = NewError(http.StatusInternalServerError, 200527, "获取管理员列表失败")
	AdminNotFind =NewError(http.StatusInternalServerError, 200528, "管理员不存在")
	UpdateRightError =NewError(http.StatusInternalServerError, 200529, "更权失败")
	GetUserError =NewError(http.StatusInternalServerError, 200530, "获取用户失败")
	SendError =NewError(http.StatusInternalServerError, 200531, "发送验证码失败")
	VartiyError =NewError(http.StatusInternalServerError, 200532, "验证码不匹配")
)

func OtherError(message string) *Error {
	return NewError(http.StatusForbidden, 100403, message)
}
