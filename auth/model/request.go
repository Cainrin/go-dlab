package model

type UserLoginReq struct {
	Email           string `json:"email" v:"email@email|required#邮箱不合法|邮箱不能为空"`
	Password        string `json:"password" v:"password@required#密码不能为空"`
	CaptchaId       string `json:"captcha_id"`
	CaptchaResponse string `json:"captcha_response"`
}

type ChangePasswordReq struct {
	OldPassword  string `json:"old_password" v:"old_password@required#旧密码不能为空"`
	NewPassword1 string `json:"new_password1" v:"new_password1@required|different:OldPassword#新密码不能为空|新密码不能与旧密码相同"`
	NewPassword2 string `json:"new_password2" v:"new_password2@required|different:OldPassword|same:NewPassword1#新密码不能为空|新密码不能与旧密码相同|两次密码不同"`
}

type ForgetPasswordReq struct {
	Email           string `json:"email" v:"email@required#邮箱不能为空"`
	Role            string `json:"role" v:"role@required"`
	CaptchaId       string `json:"captcha_id" v:"captcha_id@required#CaptchaId不能为空"`
	CaptchaResponse string `json:"captcha_response" v:"captcha_response@required#验证码不能为空"`
}

type ResetPasswordReq struct {
	Code      string `json:"code" v:"code@required#Code不能为空"`
	Password1 string `json:"password1" v:"password1@required#Password1不能为空"`
	Password2 string `json:"password2" v:"password2@required|same:Password1#Password2不能为空|两个密码不相同"`
}
