package service

import (
	"goNotificationService/dtos"
	"log"

	"github.com/spf13/viper"
	"gopkg.in/mail.v2"
)

type Service struct {
}

func NewService() *Service {
	return &Service{}
}

func (s *Service) SendMessage(input dtos.SendMessageRequest) error {
	username := viper.GetString("app_username")
	password := viper.GetString("app_password")

	log.Print(password)

	smtpHost := "smtp.mail.ru"
	smtpPort := 465

	for _, email := range input.Emails {
		m := mail.NewMessage()
		m.SetHeader("From", username)
		m.SetHeader("To", email)
		m.SetHeader("Subject", "Notification Email")
		m.SetBody("text/plain", "Добавлен новый раздел: "+input.Title)

		d := mail.NewDialer(smtpHost, smtpPort, username, password)

		err := d.DialAndSend(m)
		if err != nil {
			return err
		}
	}

	return nil
}
