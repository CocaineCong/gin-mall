package email

import (
	"gopkg.in/mail.v2"

	"mall/conf"
)

// TODO 封装成一个对象

// SendEmail 发送邮件
func SendEmail(data, email string) error {
	m := mail.NewMessage()
	m.SetHeader("From", conf.SmtpEmail)
	m.SetHeader("To", email)
	m.SetHeader("Subject", "FanOne")
	m.SetBody("text/html", data)
	d := mail.NewDialer(conf.SmtpHost, 465, conf.SmtpEmail, conf.SmtpPass)
	d.StartTLSPolicy = mail.MandatoryStartTLS
	if err := d.DialAndSend(m); err != nil {
		return err
	}

	return nil
}
