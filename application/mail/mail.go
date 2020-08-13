package mail

import (
	"log"

	"github.com/Satish-Masa/ec-backend/config"
	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

func SendEmail(email, token string) error {
	from := mail.NewEmail("Check Email", "email@example.com")
	subject := "Check your Email"
	to := mail.NewEmail("Check Your Email", email)
	plainTextContent := "This is login from!! "
	htmlContent := "This is the Login Form & Mail Check!!(http://localhost:8081/mailcheck/?token=" + token + ")"
	message := mail.NewSingleEmail(from, subject, to, plainTextContent, htmlContent)
	client := sendgrid.NewSendClient(config.Config.APIKey)
	response, err := client.Send(message)
	if err != nil {
		return err
	} else {
		log.Println(response.StatusCode)
		log.Println(response.Body)
		log.Println(response.Headers)
	}
	return nil
}
