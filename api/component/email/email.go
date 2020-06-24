package email

import (
    "configurator/api/component/config"
    "configurator/api/component/logger"
    "github.com/jordan-wright/email"
    "net/smtp"
    "net/textproto"
    "time"
)

var (
	E *email.Pool
)

func init() {
	var err error
	E, err = email.NewPool(
		config.C.GetString("email.address"),
		config.C.GetInt("email.pool.maxSize"),
		smtp.PlainAuth("", config.C.GetString("email.username"), config.C.GetString("email.password"), config.C.GetString("email.host")),
	)
	if err != nil {
		logger.L.Fatalf("Fatal error email: %v\n", err)
	}
}

func Send(to []string, from string, subject string, html string) error {
	e := &email.Email{
		To:      to,
		From:    from,
		Subject: subject,
		HTML:    []byte(html),
		Headers: textproto.MIMEHeader{},
	}
	return E.Send(e, 10*time.Second)
}
