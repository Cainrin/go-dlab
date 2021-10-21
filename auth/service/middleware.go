package service

import (
	"github.com/Cainrin/go-dlab/auth/model"
	"github.com/Cainrin/go-dlab/errors"
	"github.com/Cainrin/go-dlab/response"
	"github.com/gogf/gf/net/ghttp"
	"net/http"
)

// Middleware 中间件管理服务
func New() *middlewareService {
	return &middlewareService{}
}

type middlewareService struct{}

// Ctx 自定义上下文对象
func (s *middlewareService) Ctx(m map[string]string) func(r *ghttp.Request) {

	return func(r *ghttp.Request) {
		// 初始化，务必最开始执行
		customCtx := &model.Context{
			Session: r.Session,
		}
		Context.Init(r, customCtx)
		//action :=
		if user := Session.GetUser(r.Context()); user != nil {
			customCtx.User = &model.ContextUser{
				Id:       user.Id,
				Email:    user.Email,
				Username: user.Username,
			}
		}
		if err := r.GetError(); err != nil {
			response.FailedJsonExit(r, errors.ServiceError("服务器内部错误，请联系管理员", err))
		}
		// 执行下一步请求逻辑
		r.Middleware.Next()
		if r.Response.Status >= http.StatusInternalServerError {
			r.Response.ClearBuffer()
			response.FailedJsonExit(r, errors.ServiceError("服务器内部错误，请联系管理员"))
		}
	}

}

// Auth 鉴权中间件，只有登录成功之后才能通过
func (s *middlewareService) Auth(r *ghttp.Request) {
	if User.IsSignedIn(r.Context()) {
		r.Middleware.Next()
	} else {
		response.FailedJsonExit(r, errors.WithoutAuthorizationError("无鉴权信息"))
	}
}

// CORS 允许接口跨域请求
func (s *middlewareService) CORS(r *ghttp.Request) {
	r.Response.CORSDefault()
	r.Middleware.Next()
}

// ActiveRequired 必须为active用户
func (s *middlewareService) ActiveRequired(f func(user *model.BasicUserInfo) (bool, error)) func(r *ghttp.Request) {
	return func(r *ghttp.Request) {
		if user := User.GetProfile(r.GetCtx()); user != nil {
			ok, err := f(user)
			if ok {
				r.Middleware.Next()
			} else {
				if err == nil {
					response.FailedJsonExit(r, errors.AccountUnavailableError("账号已被停用"))
				} else {
					response.FailedJsonExit(r, err)
				}
			}
		} else {
			response.FailedJsonExit(r, errors.WithoutAuthorizationError("无鉴权信息"))
		}
	}
}

// DefaultUserCtx 自定义上下文对象
func (s *middlewareService) DefaultUserCtx(r *ghttp.Request) {
	// 初始化，务必最开始执行
	customCtx := &model.Context{
		Session: r.Session,
	}
	Context.Init(r, customCtx)
	if user := Session.GetUser(r.Context()); user != nil {
		customCtx.User = &model.ContextUser{
			Id:       2,
			Email:    "example@ex.com",
			Username: "defalutuser",
		}
		//log.ActionLogger.Log(user, r)
	}

	if err := r.GetError(); err != nil {
		response.FailedJsonExit(r, errors.ServiceError("服务器内部错误，请联系管理员", err))
	}

	// 执行下一步请求逻辑
	r.Middleware.Next()

	if r.Response.Status >= http.StatusInternalServerError {
		r.Response.ClearBuffer()
		response.FailedJsonExit(r, errors.ServiceError("服务器内部错误，请联系管理员"))
	}
}
