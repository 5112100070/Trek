package utils

import (
	"crypto/tls"
	"fmt"
	"log"
	"net/smtp"
	"strings"
)

type Mail struct {
	senderId string
	toIds    []string
	subject  string
	body     string
}

type SmtpServer struct {
	host string
	port string
}

func (s *SmtpServer) ServerName() string {
	return s.host + ":" + s.port
}

func NewMail() Mail {
	mail := Mail{}
	mail.senderId = "treknoreply@gmail.com"
	return mail
}

func (mail *Mail) BuildReciever(recievers ...string) {
	mail.toIds = recievers

	//mail.toIds = []string{"ghazi.binarandi@gmail.com", "hanifnarendra@gmail.com", "russians.wahyu@gmail.com"}
}

func (mail *Mail) BuildMessage(message string) {
	mail.subject = "[TREK] PERMINTAAN SEWA"
	mail.body = message
}

func (mail *Mail) Build() string {
	message := ""
	message += fmt.Sprintf("From: %s\r\n", mail.senderId)
	if len(mail.toIds) > 0 {
		message += fmt.Sprintf("To: %s\r\n", strings.Join(mail.toIds, ";"))
	}

	message += fmt.Sprintf("Subject: %s\r\n", mail.subject)
	message += "\r\n" + mail.body

	return message
}

func (mail *Mail) SendMail() {
	smtpServer := SmtpServer{host: "smtp.gmail.com", port: "465"}

	// Build an auth
	auth := smtp.PlainAuth("", mail.senderId, "trek1234", smtpServer.host)

	// Gmail will reject connection if it's not secure
	// TLS config
	tlsconfig := &tls.Config{
		InsecureSkipVerify: true,
		ServerName:         smtpServer.host,
	}

	conn, err := tls.Dial("tcp", smtpServer.ServerName(), tlsconfig)
	if err != nil {
		log.Panic(err)
	}

	client, err := smtp.NewClient(conn, smtpServer.host)
	if err != nil {
		log.Panic(err)
	}

	// step 1: Use Auth
	if err = client.Auth(auth); err != nil {
		log.Panic(err)
	}

	// step 2: add all from and to
	if err = client.Mail(mail.senderId); err != nil {
		log.Panic(err)
	}
	for _, k := range mail.toIds {
		if err = client.Rcpt(k); err != nil {
			log.Panic(err)
		}
	}

	// Data
	w, err := client.Data()
	if err != nil {
		log.Panic(err)
	}

	_, err = w.Write([]byte(mail.Build()))
	if err != nil {
		log.Panic(err)
	}

	err = w.Close()
	if err != nil {
		log.Panic(err)
	}

	client.Quit()

	log.Println("Message Print Successfully")
}

// func main() {
// 	mail := Mail{}
// 	mail.senderId = "wahyu.kukuh.herlambang@gmail.com"
// 	mail.toIds = []string{"ghazi.binarandi@gmail.com", "hanifnarendra@gmail.com", "russians.wahyu@gmail.com"}
// 	mail.subject = "This is the email subject"
// 	mail.body = "We Support For The Justice"

// 	messageBody := mail.BuildMessage()

// 	smtpServer := SmtpServer{host: "smtp.gmail.com", port: "465"}

// 	log.Println(smtpServer.host)
// 	//build an auth
// 	auth := smtp.PlainAuth("", mail.senderId, "SOVIET085740378580", smtpServer.host)

// 	// Gmail will reject connection if it's not secure
// 	// TLS config
// 	tlsconfig := &tls.Config{
// 		InsecureSkipVerify: true,
// 		ServerName:         smtpServer.host,
// 	}

// 	conn, err := tls.Dial("tcp", smtpServer.ServerName(), tlsconfig)
// 	if err != nil {
// 		log.Panic(err)
// 	}

// 	client, err := smtp.NewClient(conn, smtpServer.host)
// 	if err != nil {
// 		log.Panic(err)
// 	}

// 	// step 1: Use Auth
// 	if err = client.Auth(auth); err != nil {
// 		log.Panic(err)
// 	}

// 	// step 2: add all from and to
// 	if err = client.Mail(mail.senderId); err != nil {
// 		log.Panic(err)
// 	}
// 	for _, k := range mail.toIds {
// 		if err = client.Rcpt(k); err != nil {
// 			log.Panic(err)
// 		}
// 	}

// 	// Data
// 	w, err := client.Data()
// 	if err != nil {
// 		log.Panic(err)
// 	}

// 	_, err = w.Write([]byte(messageBody))
// 	if err != nil {
// 		log.Panic(err)
// 	}

// 	err = w.Close()
// 	if err != nil {
// 		log.Panic(err)
// 	}

// 	client.Quit()

// 	log.Println("Mail sent successfully")

// }
