package email

import (
	"crypto/tls"
	"fmt"
	"net/smtp"
	"strings"

	"wvp-go/server/global"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type emailPlug struct {
	to          string
	from        string
	host        string
	secret      string
	nickname    string
	port        int
	isSSL       bool
	isLoginAuth bool
}

// CreateEmailPlug 创建邮件插件
func CreateEmailPlug(to, from, host, secret, nickname string, port int, isSSL, isLoginAuth bool) *emailPlug {
	return &emailPlug{
		to:          to,
		from:        from,
		host:        host,
		secret:      secret,
		nickname:    nickname,
		port:        port,
		isSSL:       isSSL,
		isLoginAuth: isLoginAuth,
	}
}

func (e *emailPlug) Register(group *gin.RouterGroup) {
	group.POST("/email/send", e.SendEmail)
}

func (e *emailPlug) RouterPath() string {
	return "/email"
}

type emailReq struct {
	To      string `json:"to" binding:"required"`
	Subject string `json:"subject" binding:"required"`
	Body    string `json:"body" binding:"required"`
}

func (e *emailPlug) SendEmail(c *gin.Context) {
	var req emailReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"code": 400, "msg": "参数错误: " + err.Error()})
		return
	}

	msg := []byte(fmt.Sprintf("Subject: %s\r\n"+
		"From: %s <%s>\r\n"+
		"To: %s\r\n"+
		"MIME-Version: 1.0\r\n"+
		"Content-Type: text/plain; charset=UTF-8\r\n"+
		"\r\n"+
		"%s", req.Subject, e.nickname, e.from, req.To, req.Body))

	to := strings.Split(req.To, ",")
	addr := fmt.Sprintf("%s:%d", e.host, e.port)
	auth := smtp.PlainAuth("", e.from, e.secret, e.host)

	var err error
	if e.isSSL {
		err = sendWithSSL(addr, auth, e.from, to, msg)
	} else {
		err = smtp.SendMail(addr, auth, e.from, to, msg)
	}

	if err != nil {
		global.GVA_LOG.Error("发送邮件失败", zap.Error(err))
		c.JSON(500, gin.H{"code": 500, "msg": "发送邮件失败: " + err.Error()})
		return
	}
	c.JSON(200, gin.H{"code": 200, "msg": "发送成功"})
}

func sendWithSSL(addr string, auth smtp.Auth, from string, to []string, msg []byte) error {
	host := strings.Split(addr, ":")[0]
	conn, err := tls.Dial("tcp", addr, &tls.Config{
		InsecureSkipVerify: true,
	})
	if err != nil {
		return err
	}
	defer conn.Close()

	c, err := smtp.NewClient(conn, host)
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
