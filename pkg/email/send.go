package email

import (
	"goapi/pkg/logger"
	"gopkg.in/gomail.v2"
)

func SendOneEmail(params Params) error {
	if len(params.ContentType) == 0 {
		params.ContentType = "text/plain"
	}
	m := gomail.NewMessage()
	m.SetHeader("From", params.From) // 发件人
	//m.SetHeader("From", fmt.Sprintf("%s<%s>", params.Subject, params.From)) // 发件人
	m.SetHeader("To", params.To)               // 收件人
	m.SetHeader("Subject", params.Subject)     // 标题
	m.SetBody(params.ContentType, params.Body) // 内容
	d := gomail.NewDialer(params.Host, params.Port, params.Username, params.Password)
	if err := d.DialAndSend(m); err != nil {
		logger.Error(err)
		return err
	}
	return nil
}

func SendBatchEmail(params Params) error {
	if len(params.ContentType) == 0 {
		params.ContentType = "text/plain"
	}
	m := gomail.NewMessage()
	m.SetHeader("From", params.From)           // 发件人
	m.SetHeader("To", params.Tos...)           // 收件人
	m.SetHeader("Subject", params.Subject)     // 标题
	m.SetBody(params.ContentType, params.Body) // 内容
	d := gomail.NewDialer(params.Host, params.Port, params.Username, params.Password)
	if err := d.DialAndSend(m); err != nil {
		logger.Error(err)
		return err
	}
	return nil
}
