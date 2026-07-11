package utils

import (
	"crypto/tls"
	"fmt"
	"net/smtp"
	"strings"

	"wvp-go/server/global"
)

// ErrorToEmail 发送错误信息到邮箱
func ErrorToEmail(subject string, body string) error {
	cfg := global.GVA_CONFIG.Email
	if cfg.Host == "" || cfg.From == "" {
		return fmt.Errorf("邮件配置不完整")
	}

	to := strings.Split(cfg.To, ",")
	msg := []byte(fmt.Sprintf("Subject: %s\r\n"+
		"From: %s <%s>\r\n"+
		"To: %s\r\n"+
		"MIME-Version: 1.0\r\n"+
		"Content-Type: text/plain; charset=UTF-8\r\n"+
		"\r\n"+
		"%s", subject, cfg.Nickname, cfg.From, cfg.To, body))

	addr := fmt.Sprintf("%s:%d", cfg.Host, cfg.Port)
	auth := smtp.PlainAuth("", cfg.From, cfg.Secret, cfg.Host)

	if cfg.IsSSL {
		return sendWithSSL(addr, auth, cfg.From, to, msg)
	}

	return smtp.SendMail(addr, auth, cfg.From, to, msg)
}

func sendWithSSL(addr string, auth smtp.Auth, from string, to []string, msg []byte) error {
	conn, err := tls.Dial("tcp", addr, &tls.Config{
		InsecureSkipVerify: true,
	})
	if err != nil {
		return err
	}
	defer conn.Close()

	c, err := smtp.NewClient(conn, strings.Split(addr, ":")[0])
	if err != nil {
		return err
	}
	defer c.Close()

	if err = c.Auth(auth); err != nil {
		return err
	}
	if err = c.Mail(from); err != nil {
		return err
	}
	for _, addr := range to {
		if err = c.Rcpt(addr); err != nil {
			return err
		}
	}
	w, err := c.Data()
	if err != nil {
		return err
	}
	if _, err = w.Write(msg); err != nil {
		return err
	}
	return w.Close()
}
