package tools

import (
	"context"
	"encoding/hex"
	"github.com/Cainrin/go-dlab/errors"
	"github.com/gogf/gf/database/gdb"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/os/gtime"
	"github.com/google/uuid"
	"time"
)

var ResetPasswordCodeStore = resetPasswordCodeStore{g.DB(), "reset_password_store"}

type resetPasswordCodeStore struct {
	db        gdb.DB
	tableName string
}

type resetPasswordCode struct {
	Id          int64       `orm:"id"`
	UserRole    uint        `orm:"user_role"`
	UserEmail   string      `orm:"user_email"`
	Code        string      `orm:"code"`
	ExpiredTime *gtime.Time `orm:"expired_time"`
}

func (s *resetPasswordCodeStore) GenerateCode(ctx context.Context, userEmail string, userRole int) (string, error) {
	dur, err := time.ParseDuration("1h")
	if err != nil {
		return "", err
	}
	code := hex.EncodeToString([]byte(uuid.New().String()))

	if _, err := s.db.Model(s.tableName).Ctx(ctx).Data(g.Map{
		"user_email":   userEmail,
		"user_role":    userRole,
		"code":         code,
		"expired_time": time.Now().Add(dur),
	}).Insert(); err != nil {
		return "", errors.ORMError("生成重置密码的code错误", err)
	}
	return code, nil
}

func (s *resetPasswordCodeStore) IsExistByEmail(email string, role int) bool {
	s.removeExpired()
	if count, err := s.db.Model(s.tableName).FindCount(gdb.Map{"user_email": email, "user_role": role}); err != nil {
		return false
	} else {
		return count != 0
	}
}

func (s *resetPasswordCodeStore) VerifyCode(code string) (uint, string, error) {
	s.removeExpired()

	var target *resetPasswordCode
	if err := s.db.Model(s.tableName).Where("code", code).Scan(&target); err != nil {
		return 0, "", err
	}

	if target == nil {
		return 0, "", errors.NotFoundError("数据库中无此code")
	}

	return target.UserRole, target.UserEmail, nil
}

func (s *resetPasswordCodeStore) removeExpired() {
	_, _ = s.db.Model(s.tableName).Delete("expired_time<=", time.Now())
}

func (s *resetPasswordCodeStore) RemoveCode(ctx context.Context, code string) error {
	if _, err := s.db.Model(s.tableName).Ctx(ctx).Delete("code", code); err != nil {
		return err
	}

	return nil
}
