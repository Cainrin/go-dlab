package errors

const (
	CUserOrPasswordFailed = 1000 // 账号密码错误
	CWithoutAuthorization = 1001 // 没有鉴权
	CAuthorizationFailed  = 1002 // 鉴权错误
	CCaptchaMissed        = 1003 // 没有验证码
	CCaptchaIncorrect     = 1004 // 验证码错误
	CAccountUnavailable   = 1005 // 账号不可用
	CHasNoPermission      = 1006 // 没有权限
	CFormInvalid          = 1007 // 表单参数格式错误
	CInvalid              = 1008 // 无效请求
	CConflict             = 1009 // 与现有数据冲突
	CNotFound             = 1010 // 没找到所查找Object
	CORMError             = 1011 // ORM错误
	CServiceError         = 1012 // 业务逻辑错误
	CLimitError           = 1013 // 操作被限制
	CUnknownError         = 1999 // 未知的错误
)

type ServerError struct {
	Code    int
	Message string
	Err     error
}

func (s *ServerError) Error() string {
	if s.Err != nil {
		return s.Message + ":" + s.Err.Error()
	}
	return s.Message
}

func UserOrPasswordFailed(message string, errs ...error) *ServerError {
	return &ServerError{
		Code:    CUserOrPasswordFailed,
		Message: message,
		Err:     getError(errs...),
	}
}

func InvalidError(message string, errs ...error) *ServerError {
	return &ServerError{
		Code:    CInvalid,
		Message: message,
		Err:     getError(errs...),
	}
}

func NotFoundError(message string, errs ...error) *ServerError {
	return &ServerError{
		Code:    CNotFound,
		Message: message,
		Err:     getError(errs...),
	}
}

func ConflictError(message string, errs ...error) *ServerError {
	return &ServerError{
		Code:    CConflict,
		Message: message,
		Err:     getError(errs...),
	}
}

func NotPermissionError(message string, errs ...error) *ServerError {
	return &ServerError{
		Code:    CHasNoPermission,
		Message: message,
		Err:     getError(errs...),
	}
}

func WithoutAuthorizationError(message string, errs ...error) *ServerError {
	return &ServerError{
		Code:    CWithoutAuthorization,
		Message: message,
		Err:     getError(errs...),
	}
}

func CaptchaMissedError(message string, errs ...error) *ServerError {
	return &ServerError{
		Code:    CCaptchaMissed,
		Message: message,
		Err:     getError(errs...),
	}
}

func CaptchaIncorrectError(message string, errs ...error) *ServerError {
	return &ServerError{
		Code:    CCaptchaIncorrect,
		Message: message,
		Err:     getError(errs...),
	}
}

func AccountUnavailableError(message string, errs ...error) *ServerError {
	return &ServerError{
		Code:    CAccountUnavailable,
		Message: message,
		Err:     getError(errs...),
	}
}

func AuthorizationFailedError(message string, errs ...error) *ServerError {
	return &ServerError{
		Code:    CAuthorizationFailed,
		Message: message,
		Err:     getError(errs...),
	}
}

func FormFormatInvalidError(message string, errs ...error) *ServerError {
	return &ServerError{
		Code:    CFormInvalid,
		Message: message,
		Err:     getError(errs...),
	}
}

func ORMError(message string, errs ...error) *ServerError {
	return &ServerError{
		Code:    CORMError,
		Message: message,
		Err:     getError(errs...),
	}
}

func ServiceError(message string, errs ...error) *ServerError {
	return &ServerError{
		Code:    CServiceError,
		Message: message,
		Err:     getError(errs...),
	}
}

func getError(err ...error) error {
	if err != nil {
		return err[0]
	}
	return nil
}

func LimitError(message string, errs ...error) *ServerError {
	return &ServerError{
		Code:    CLimitError,
		Message: message,
		Err:     getError(errs...),
	}
}
