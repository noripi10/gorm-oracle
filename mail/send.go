package mail

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/jordan-wright/email"
)

func SendMail(message string) {
	// auth := smtp.PlainAuth("", "", "", "")
	address := os.Getenv("SMPT_ADDRESS")
	port, err := strconv.Atoi(os.Getenv("SMPT_PORT"))
	if err != nil {
		log.Fatal(err)
	}

	addr := fmt.Sprintf("%s:%d", address, port)
	from := fmt.Sprintf("%s <%s>", os.Getenv("FROM_NAME"), os.Getenv("FROM_ADDRESS"))
	to := []string{os.Getenv("TO_ADDRESS")}

	subject := "Go Email Sample"
	body := message

	email := email.NewEmail()
	email.From = from
	email.To = to
	// email.Cc =
	// email.Bcc =
	email.Subject = subject
	email.Text = []byte(body)
	sendErr := email.Send(addr, nil)
	if sendErr != nil {
		log.Fatal(sendErr)
	}
}
