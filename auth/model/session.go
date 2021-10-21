package model

type BasicUserInfo struct {
	Role       uint   `json:"role"`
	Id         int64  `json:"id"`
	CompanyId  int64  `json:"company_id"`
	Email      string `json:"email"`
	Username   string `json:"username"`
	OpenId     string `json:"open_id"`
	SessionKey string `json:"session_key"`
	IsStaff    bool   `json:"is_staff"`
	IsSuper    bool   `json:"is_super"`
}
