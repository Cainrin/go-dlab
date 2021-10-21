package log

import (
	"github.com/Cainrin/go-dlab/auth/model"
	"github.com/gogf/gf/database/gdb"
	"github.com/gogf/gf/errors/gerror"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/os/gtime"
	"strings"
)

func New(tableName string) *actionLogDao {
	return &actionLogDao{
		g:         g.DB(),
		tableName: tableName,
	}
}

type actionLogDao struct {
	g         gdb.DB
	tableName string
}

type UserLog struct {
	Id           int64       `orm:"id" json:"id"`
	Role         uint        `orm:"role" json:"role"`
	UserId       int64       `orm:"user_id" json:"user_id"`
	CompanyId    int64       `orm:"company_id" json:"company_id"`
	Category     string      `orm:"category" json:"category"`
	Username     string      `orm:"username" json:"username"`
	CompanyName  string      `orm:"company_name" json:"company_name"`
	Ip           string      `orm:"ip" json:"ip"`
	UserAgent    string      `orm:"user_agent" json:"user_agent"`
	Uri          string      `orm:"uri" json:"uri"`
	Action       string      `orm:"action" json:"action"`
	OccurredTime *gtime.Time `orm:"occurred_time" json:"occurred_time"`
}

func (a *actionLogDao) Log(user *model.BasicUserInfo, r *ghttp.Request, detail string) {
	if len(detail) == 0 {
		return
	}

	l := strings.Split(detail, ";")
	if len(l) == 1 {
		l = []string{l[0], ""}
	}

	g.Log().Debugf("%s %s %s %s", r.GetUrl(), r.GetClientIp(), detail, r.URL.Path)

	if _, err := a.g.Model(a.tableName).Data(g.Map{
		"ip":         r.GetClientIp(),
		"user_agent": r.GetHeader("user-agent"),
		"user_id":    user.Id,
		"username":   user.Username,
		"company_id": user.CompanyId,
		"role":       user.Role,
		"uri":        r.GetUrl(),
		"category":   l[0],
		"action":     l[1],
		"body":       r.GetBodyString(),
	}).Insert(); err != nil {
		//panic(err)
		g.Log().Errorf("中间件日志记录出错 请检查 %s", err.Error())
	}
}

func (a *actionLogDao) FetchKocLog(role uint, query map[string]interface{}, offset int, count int, parser interface{}) (int, error) {
	var table string
	switch role {
	case 0:
		table = "koc_log"
	case 1:
		table = "brand_log"
	default:
		return 0, gerror.New("角色错误请检查传输字段！")
	}
	record := a.g.Model(table).Where(query)
	total, err := record.Count()
	err = record.Offset(offset).Limit(count).OrderDesc("occurred_time").Scan(parser)
	if err != nil {
		g.Log().Errorf("查询日志模块出现严重错误请检查！", err)
		return 0, gerror.New("查询日志出现错误请检查")
	}
	return total, nil
}
