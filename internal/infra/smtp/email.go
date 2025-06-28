package smtp

import (
	"fmt"
	"net/smtp"
	"strings"
)

type Smtp struct {
	From     string
	Port     string
	Password string
	Host     string
}

type SmptSendEmail struct {
	To      []string
	Message string
	Subject string
}

func (s *Smtp) Send(email SmptSendEmail) error {
	auth := smtp.PlainAuth("", s.From, s.Password, s.Host)
	addr := s.Host + ":" + s.Port

	msg := []byte("To: " + strings.Join(email.To, ",") + "\r\n" +
		"Subject: " + email.Subject + "\r\n" +
		"\r\n" +
		email.Message + "\r\n")

	err := smtp.SendMail(addr, auth, s.From, email.To, msg)
	if err != nil {
		return fmt.Errorf("erro ao enviar e-mail: %w", err)
	}
	return nil
}
