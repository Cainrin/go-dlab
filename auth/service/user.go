package service

import (
	"context"
	"github.com/Cainrin/go-dlab/auth/model"
)

// User 中间件管理服务
var User = userService{}

type userService struct{}

// IsSignedIn 判断用户是否已经登录
func (s *userService) IsSignedIn(ctx context.Context) bool {
	if v := Context.Get(ctx); v != nil && v.User != nil {
		return true
	}
	return false
}

// SignOut 用户注销
func (s *userService) SignOut(ctx context.Context) error {
	return Session.RemoveUser(ctx)
}

// GetProfile 获得用户信息详情
func (s *userService) GetProfile(ctx context.Context) *model.BasicUserInfo {
	return Session.GetUser(ctx)
}
