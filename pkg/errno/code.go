package errno

var (
	// 常量错误
	OK                  = &Errno{Code: 0, Message: "OK"}
	InternalServerError = &Errno{Code: 10001, Message: "内部错误"}
	ErrBind             = &Errno{Code: 10002, Message: "将请求体绑定到结构时发生错误"}

	ErrValidation = &Errno{Code: 20001, Message: "无效验证."}
	ErrDatabase   = &Errno{Code: 20002, Message: "数据库出现错误."}
	ErrToken      = &Errno{Code: 20003, Message: "令牌签证发生错误."}

	// 用户错误
	ErrEncrypt           = &Errno{Code: 20101, Message: "加密过程发生错误."}
	ErrUserNotFound      = &Errno{Code: 20102, Message: "没有找到用户."}
	ErrTokenInvalid      = &Errno{Code: 20103, Message: "token无效"}
	ErrPasswordIncorrect = &Errno{Code: 20104, Message: "密码错误."}
)
