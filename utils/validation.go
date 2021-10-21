package utils

import (
	le "github.com/Cainrin/go-dlab/errors"
	"github.com/Cainrin/go-dlab/response"
	"github.com/gogf/gf/net/ghttp"
)

func Validation(r *ghttp.Request, parser interface{}) {
	if err := r.Parse(parser); err != nil {
		response.FailedJsonExit(r, le.FormFormatInvalidError("表单验证错误", err))
	}
}
