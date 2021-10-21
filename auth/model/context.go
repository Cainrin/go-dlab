package model

import "github.com/gogf/gf/net/ghttp"

const (
	// ContextKey 上下文变量存储键名
	ContextKey = "ContextKey"
)

// Context 请求上下文结构
type Context struct {
	Session *ghttp.Session // 当前Session管理对象
	User    *ContextUser   // 上下文用户信息
}

// ContextUser 请求上下文中的用户信息
type ContextUser struct {
	Id       int64  // 用户ID
	Username string // 用户名称
	Email    string // 邮箱
}
