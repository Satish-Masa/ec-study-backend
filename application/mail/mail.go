package mail

import (
	"net/smtp"
)

func SendMail(email string) error {
	auth := smtp.PlainAuth("", "user@example.com", "password", "mail.example.com")
	to := []string{email}
	msg := []byte("Subject: メールアドレスの確認" + "\r\n" + "\r\n" + "ログイン画面 http://localhost:8080/login\r\n")

	err := smtp.SendMail("mail.example.com:25", auth, "sender@example.net", to, msg)
	if err != nil {
		return err
	}

	return nil
}
