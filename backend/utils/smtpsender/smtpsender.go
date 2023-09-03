package smtpsender

import (
	"crypto/tls"
	"fmt"
	"mime"
	"net/smtp"
	"strings"
	"sync"

	"implude.kr/VOAH-Backend-Core/configs"
	"implude.kr/VOAH-Backend-Core/utils/logger"
)

var SMTPServer SMTPServerInfo
var tlsConfig *tls.Config

type SMTPServerInfo struct {
	Host     string
	Port     int
	Username string
	Password string
}

func (s *SMTPServerInfo) Address() string {
	return fmt.Sprintf("%s:%d", s.Host, s.Port)
}

func (s *SMTPServerInfo) Connect() (*smtp.Client, error) {
	if configs.Env.SMTP.SSL {
		c, err := tls.Dial("tcp", s.Address(), tlsConfig)
		if err != nil {
			return nil, err
		}
		return smtp.NewClient(c, s.Host)
	} else if configs.Env.SMTP.STARTTLS {
		c, err := smtp.Dial(s.Address())
		if err != nil {
			return nil, err
		}
		if err = c.StartTLS(tlsConfig); err != nil {
			return nil, err
		}
		return c, nil
	} else {
		return smtp.Dial(s.Address())
	}
}

type Mail struct {
	From    string   `json:"from"`
	Tos     []string `json:"to"`
	Subject string   `json:"subject"`
	Body    string   `json:"body"`
}

func (mail *Mail) BuildMessage() string {
	message := ""
	message += fmt.Sprintf("From: %s\r\n", mail.From)
	if len(mail.Tos) > 0 {
		message += fmt.Sprintf("To: %s\r\n", strings.Join(mail.Tos, ";"))
	}

	message += fmt.Sprintf("Subject: %s\r\n", mail.Subject)
	message += "\r\n" + mail.Body

	return message
}
func (mail *Mail) cSubject() {
	mail.Subject = mime.QEncoding.Encode("utf-8", mail.Subject)
}
func (mail *Mail) Send(c *smtp.Client) error {
	smtpConf := configs.Env.SMTP
	log := logger.Logger
	if err := c.Auth(smtp.PlainAuth("", smtpConf.Username, smtpConf.Password, smtpConf.Host)); err != nil {
		log.Error(err.Error())
		return err
	}
	if err := c.Mail(mail.From); err != nil {
		log.Error(err.Error())
		return err
	}
	for _, k := range mail.Tos {
		if err := c.Rcpt(k); err != nil {
			log.Error(err.Error())
			return err
		}
	}
	mail.cSubject() // convert subject to utf-8
	plainMail := mail.BuildMessage()
	w, err := c.Data()
	if err != nil {
		log.Error(err.Error())
		return err
	}

	if _, err = w.Write([]byte(plainMail)); err != nil {
		log.Error(err.Error())
		return err
	}
	if err = w.Close(); err != nil {
		log.Error(err.Error())
		return err
	}
	return nil
}

func (mail *Mail) ConnectAndSend() error {
	smtpConn, err := SMTPServer.Connect()
	if err != nil {
		return err
	}
	if err := mail.Send(smtpConn); err != nil {
		return err
	}
	smtpConn.Quit()
	return nil
}

func InitSMTP(wait *sync.WaitGroup) {
	log := logger.Logger
	smtpConf := configs.Env.SMTP
	SMTPServer = SMTPServerInfo{
		Host:     smtpConf.Host,
		Port:     smtpConf.Port,
		Username: smtpConf.Username,
		Password: smtpConf.Password,
	}
	tlsConfig = &tls.Config{
		InsecureSkipVerify: true, // 이 옵션은 테스트 목적으로만 사용해야 합니다. 실제 환경에서는 false로 설정해야 합니다.
		ServerName:         smtpConf.Host,
	}
	c, err := SMTPServer.Connect()
	if err != nil {
		log.Error("SMTP Connection Error")
		log.Fatal(err)
	}
	if err = c.Auth(smtp.PlainAuth("", smtpConf.Username, smtpConf.Password, smtpConf.Host)); err != nil {
		log.Error("SMTP Auth Error")
		log.Fatal(err)
	}
	if err = c.Quit(); err != nil {
		log.Error("SMTP Quit Error")
		log.Fatal(err)
	}

	log.Info("Connected to SMTP server")
	defer wait.Done()
}
