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
	UserAlreadyExisted = NewError(http.StatusInternalServerError, 200502, "用户已存在")
	UsernameError      = NewError(http.StatusInternalServerError, 200503, "用户名长度必须为12位")
	PasswordError      = NewError(http.StatusInternalServerError, 200504, "密码长度必须大于8且小于16位")
	Register           = NewError(http.StatusInternalServerError, 200505, "注册失败")

	//登录
	UserNotFind           = NewError(http.StatusInternalServerError, 200506, "该用户不存在")
	NoThatPasswordOrWrong = NewError(http.StatusInternalServerError, 200507, "密码错误")

	//鉴权
	ServerError = NewError(http.StatusInternalServerError, 200508, "系统异常，请稍后重试!")
	AuthExpired = NewError(http.StatusInternalServerError, 200509, "登陆状态已过期，请重新登陆")

	//修改用户信息
	UpdateUserError = NewError(http.StatusInternalServerError, 200510, "修改用户信息失败")

	//提交反馈
	PostTypeError   = NewError(http.StatusInternalServerError, 200511, "反馈类型无效")
	CreatePostError = NewError(http.StatusInternalServerError, 200512, "提交反馈失败")

	//上传图片
	FileTypeError      = NewError(http.StatusInternalServerError, 200513, "文件不是图片类型")
	ParseFormDataError = NewError(http.StatusInternalServerError, 200514, "解析表单数据失败")
	FileExistedError   = NewError(http.StatusInternalServerError, 200515, "文件已存在")
	GetFileError       = NewError(http.StatusInternalServerError, 200516, "获取文件失败")

	//修改反馈
	PostNotFind          = NewError(http.StatusInternalServerError, 200517, "反馈不存在")
	UserConnotUpdatePost = NewError(http.StatusInternalServerError, 200518, "用户无权修改反馈")
	PostHasBeenHandled   = NewError(http.StatusInternalServerError, 200519, "反馈已被处理")
	EmailError           = NewError(http.StatusInternalServerError, 200520, "邮箱错误")
	UpdatePostError      = NewError(http.StatusInternalServerError, 200521, "修改反馈失败")

	//删除反馈
	UserConnotDeletePost = NewError(http.StatusInternalServerError, 200522, "用户无权删除反馈")
	DeletePostError      = NewError(http.StatusInternalServerError, 200523, "删除反馈失败")

	//获取反馈列表
	GetPostListError = NewError(http.StatusInternalServerError, 200524, "获取反馈列表失败")

	//查看回复
	GetPostError     = NewError(http.StatusInternalServerError, 200525, "用户没有提出反馈")
	GetResponseError = NewError(http.StatusInternalServerError, 200526, "查看回复失败")

	//做出评价
	TrashPost                 = NewError(http.StatusInternalServerError, 200527, "反馈被标记为垃圾信息")
	UserConnotRateResponse    = NewError(http.StatusInternalServerError, 200528, "用户无权做出评价")
	ResponseRatingError       = NewError(http.StatusInternalServerError, 200529, "评价类型无效")
	CreateResponseRatingError = NewError(http.StatusInternalServerError, 200530, "做出评价失败")

	LackRight                   = NewError(http.StatusInternalServerError, 200531, "权限不足")
	ReatHandle                  = NewError(http.StatusInternalServerError, 200532, "重复接单")
	SaveError                   = NewError(http.StatusInternalServerError, 200533, "保存失败")
	AdminUncompaired            = NewError(http.StatusInternalServerError, 200534, "处理人不一致")
	HandleError                 = NewError(http.StatusInternalServerError, 200535, "处理失败")
	GetPostResponseTimeError    = NewError(http.StatusInternalServerError, 200536, "获取回复时间失败")
	UpdatePostResponseTimeError = NewError(http.StatusInternalServerError, 200537, "更新回复时间失败")
	PostNotHandle               = NewError(http.StatusInternalServerError, 200538, "帖子未处理")
	GetAdminListError           = NewError(http.StatusInternalServerError, 200539, "获取管理员列表失败")
	AdminNotFind                = NewError(http.StatusInternalServerError, 200540, "管理员不存在")
	UpdateRightError            = NewError(http.StatusInternalServerError, 200541, "更权失败")
	GetUserError                = NewError(http.StatusInternalServerError, 200542, "获取用户失败")
	SendError =NewError(http.StatusInternalServerError, 200531, "发送验证码失败")
	VartiyError =NewError(http.StatusInternalServerError, 200532, "验证码不匹配")
	EncryptionFailed =NewError(http.StatusInternalServerError, 200533, "加密失败")
	ResponseNotExist =NewError(http.StatusInternalServerError, 200534,"回复不存在")
)

func OtherError(message string) *Error {
	return NewError(http.StatusForbidden, 100403, message)
}
