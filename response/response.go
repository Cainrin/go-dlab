package response

import (
	"encoding/xml"
	le "github.com/Cainrin/go-dlab/errors"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/util/gvalid"
)

// JsonResponse 数据返回通用JSON数据结构
type JsonResponse struct {
	Code    int         `json:"code"`    // 错误码((0:成功, 1:失败, >1:错误码))
	Message string      `json:"message"` // 提示信息
	Data    interface{} `json:"data"`    // 返回数据(业务接口定义具体数据结构)
}

// Json 标准返回结果数据结构封装。
func Json(r *ghttp.Request, code int, message string, data interface{}) {
	err := r.Response.WriteJson(JsonResponse{
		Code:    code,
		Message: message,
		Data:    data,
	})
	if err != nil {
		g.Log().Errorf("返回区域致命错误 %s", err.Error())
	}
}

// JsonExit 返回JSON数据并退出当前HTTP执行函数。
func JsonExit(r *ghttp.Request, err int, msg string, data ...interface{}) {
	if data != nil {
		Json(r, err, msg, data[0])
	} else {
		Json(r, err, msg, nil)
	}
	r.Exit()
}

// XmlExit 返回xml数据
func XmlExit(r *ghttp.Request, data interface{}) {
	if data, err := xml.Marshal(data); err != nil {
		g.Log().Errorf("xml 解析错误 %s", err.Error())
	} else {
		if err := r.Response.WriteXmlExit(data, "xml"); err != nil {
			g.Log().Errorf("返回数据致命错误 %s", err.Error())
		}
	}
}

func SuccessJsonExit(r *ghttp.Request, data ...interface{}) {
	JsonExit(r, 0, "ok", data...)
}

func FailedJsonExit(r *ghttp.Request, err error) {
	if se, ok := err.(*le.ServerError); ok {
		if se.Code == le.CFormInvalid {
			JsonExit(r, se.Code, se.Message, le.GValidError2MapData(se.Err.(gvalid.Error)))
		} else {
			g.Log().Errorf("code: %s, err: %+v", se.Code, se.Error())
			JsonExit(r, se.Code, se.Error(), nil)
		}
	} else {
		g.Log().Errorf("未知错误:" + err.Error())
		JsonExit(r, le.CUnknownError, "未知错误："+err.Error(), nil)
	}
}
