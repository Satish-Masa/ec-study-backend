package mail

import (
	"log"

	"github.com/Satish-Masa/ec-backend/config"
	domainMail "github.com/Satish-Masa/ec-backend/domain/mail"
	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

type MailApplication struct {
	Repository domainMail.MailRepository
}

type MailCheckRequest struct {
	Token string `json: "token"`
}

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

func (a MailApplication) SaveMail(m *domainMail.Mail) error {
	return a.Repository.Save(m)
}

func (a MailApplication) UpdateMail(id int, token string) error {
	return a.Repository.Update(id, token)
}

func (a MailApplication) FindMail(id int) (domainMail.Mail, error) {
	return a.Repository.Find(id)
}

func (a MailApplication) CheckMail(token string, id int) error {
	return a.Repository.Check(token, id)
}

func (a MailApplication) Validation(id int) (bool, error) {
	return a.Repository.Validation(id)
}
