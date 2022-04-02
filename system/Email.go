package system

import (
	"app/config"
	"fmt"
	"net/smtp"
	"os"
)

type mail struct {
	From    string
	To      string
	Subject string
	Body    string
}

func SendMail(to string, msg string) {

	from := config.G_MAIL.From
	username := config.G_MAIL.Username
	password := config.G_MAIL.Password

	// toList is list of email address that email is to be sent.
	toList := []string{to}

	// host is address of server that the
	// sender's email address belongs,
	// in this case its gmail.
	// For e.g if your are using yahoo
	// mail change the address as smtp.mail.yahoo.com
	host := config.G_MAIL.Host

	// Its the default port of smtp server
	port := config.G_MAIL.Port

	mail := mail{
		From:    from,
		To:      to,
		Subject: "New subscription",
		Body:    msg,
	}

	// We can't send strings directly in mail,
	// strings need to be converted into slice bytes
	body := []byte(bodyConstruct(mail))

	// PlainAuth uses the given username and password to
	// authenticate to host and act as identity.
	// Usually identity should be the empty string,
	// to act as username.
	auth := smtp.PlainAuth("", username, password, host)

	// SendMail uses TLS connection to send the mail
	// The email is sent to all address in the toList,
	// the body should be of type bytes, not strings
	// This returns error if any occured.
	err := smtp.SendMail(host+":"+port, auth, from, toList, body)

	// handling the errors
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println("Email sent successfully")
}

func bodyConstruct(m mail) string {

	msg := fmt.Sprintf("From: %s\r\n", m.From)
	msg += fmt.Sprintf("To: %s\r\n", m.To)
	msg += fmt.Sprintf("Subject: %s\r\n", m.Subject)
	msg += fmt.Sprintf("\r\n%s\r\n", m.Body)

	return msg
}
