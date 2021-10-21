package tools

import (
	"github.com/gogf/gf/database/gredis"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/util/gconv"
	"github.com/mojocn/base64Captcha"
)

var Captcha = new(captchaService)

type captchaService struct {
}

type redisStore struct {
	client *gredis.Redis
}

func (s *redisStore) Set(id string, value string) {
	_, err := s.client.Do("SET", id, value)
	_, err = s.client.Do("EXPIRE", id, 5*60)
	if err != nil {
		panic(err)
	}
}

func (s *redisStore) Get(id string, clear bool) string {
	val, err := g.Redis().DoVar("GET", id)
	if err != nil {
		return ""
	}
	if clear {
		_, err := s.client.Do("DEL", id)
		if err != nil {
			return ""
		}

	}
	return gconv.String(val)
}

func (s *redisStore) Verify(id, answer string, clear bool) bool {
	v := s.Get(id, clear)
	return v == answer
}

func newRedisStore() base64Captcha.Store {
	return &redisStore{
		client: g.Redis(),
	}
}

var store = newRedisStore()

func (s *captchaService) Generate() (string, string, error) {
	driver := base64Captcha.DefaultDriverDigit
	c := base64Captcha.NewCaptcha(driver, store)
	return c.Generate()
}

// Verify 验证验证码
func (s *captchaService) Verify(id string, val string) bool {
	if id == "" || val == "" {
		return false
	}
	return store.Verify(id, val, true)
}
