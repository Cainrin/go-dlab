package service

import (
	"context"
	"github.com/Cainrin/go-dlab/auth/model"
)

// Session Session管理服务
var Session = sessionService{}

type sessionService struct{}

const (
	// 用户信息存放在Session中的Key
	sessionKeyUser = "sessionKeyUser"
)

// SetUser 设置用户Session.
func (s *sessionService) SetUser(ctx context.Context, user *model.BasicUserInfo) error {
	return Context.Get(ctx).Session.Set(sessionKeyUser, user)
}

// GetUser 获取当前登录的用户信息对象，如果用户未登录返回nil。
func (s *sessionService) GetUser(ctx context.Context) *model.BasicUserInfo {
	customCtx := Context.Get(ctx)
	if customCtx != nil {
		if v := customCtx.Session.GetVar(sessionKeyUser); v != nil {
			var user *model.BasicUserInfo
			_ = v.Struct(&user)
			return user
		}
	}
	return nil
}

// RemoveUser 删除用户Session。
func (s *sessionService) RemoveUser(ctx context.Context) error {
	customCtx := Context.Get(ctx)
	if customCtx != nil {
		return customCtx.Session.Remove(sessionKeyUser)
	}

	return nil
}
