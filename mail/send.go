package mail

import (
	"encoding/json"
	"fmt"
	"gorm-oracle/config"
	"log"

	"github.com/jordan-wright/email"
)

func SendMail(bodyData any) error {
	// setting info
	conf, _ := config.New()
	// auth := smtp.PlainAuth("", "", "", "")
	address := conf.SmtpAddress
	port := conf.SmtpPort
	fromName := conf.FromName
	fromAddress := conf.FromAddress
	toAddress := conf.ToAdress

	addr := fmt.Sprintf("%s:%d", address, port)
	from := fmt.Sprintf("%s <%s>", fromName, fromAddress)
	to := []string{toAddress}
	subject := "Go Email Sample"

	json, err := json.Marshal(bodyData)
	if err != nil {
		log.Fatal(err)
	}
	message := "WK_DATA: " + "\r\n" + string(json)
	body := message

	email := email.NewEmail()
	email.From = from
	email.To = to
	// email.Cc =
	// email.Bcc =
	email.Subject = subject
	email.Text = []byte(body)
	// 認証無し
	sendErr := email.Send(addr, nil)
	if sendErr != nil {
		return sendErr
	}
	return nil
}
