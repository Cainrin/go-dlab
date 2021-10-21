package tools

import (
	"bytes"
	"encoding/base64"
	"github.com/gogf/gf/frame/g"
	"net/smtp"
	"os"
	"strings"
	"text/template"
)

var Email = newService()

type emailService struct {
	auth                             smtp.Auth
	host, port, from, password, name string
}

func toBase64(target string) string {
	//需引入base64库
	return "=?UTF-8?B?" + base64.StdEncoding.EncodeToString([]byte(target)) + "?="
}

func (s *emailService) SendHtmlEmail(subject string, templatePath string, to []string, data interface{}) error {
	mimeHeaders := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
	tempString, err := s.parseTemplate(templatePath, data)
	if err != nil {
		return err
	}

	msg := []byte("Subject: " + toBase64(subject) + "\n" + "From: " + toBase64(s.name) + "<" + s.from + ">" + "\n" + "To: " + strings.Join(to, ";") + "\n" + mimeHeaders + "\n" + tempString)
	if err := smtp.SendMail(s.host+":"+s.port, s.auth, s.from, to, msg); err != nil {
		return err
	}

	return nil
}

func (s *emailService) parseTemplate(templateFileName string, data interface{}) (string, error) {
	dir, err := os.Getwd()
	if err != nil {
		return "", err
	}

	t, err := template.ParseFiles(dir + "/template/" + templateFileName)
	if err != nil {
		return "", err
	}
	buf := new(bytes.Buffer)
	if err = t.Execute(buf, data); err != nil {
		return "", err
	}
	return buf.String(), nil
}

func newService() *emailService {

	host := g.Cfg().GetString("email.host")
	port := g.Cfg().GetString("email.port")
	from := g.Cfg().GetString("email.from")
	password := g.Cfg().GetString("email.password")
	name := g.Cfg().GetString("email.name")
	return &emailService{
		host:     host,
		port:     port,
		from:     from,
		name:     name,
		password: password,
		auth:     smtp.PlainAuth("", from, password, host),
	}
}
