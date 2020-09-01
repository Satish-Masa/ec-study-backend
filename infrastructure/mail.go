package infrastructure

import (
	domainMail "github.com/Satish-Masa/ec-backend/domain/mail"
	"github.com/jinzhu/gorm"
)

type mailRepository struct {
	conn *gorm.DB
}

func NewMailRepository(conn *gorm.DB) domainMail.MailRepository {
	return &mailRepository{conn: conn}
}

func (i *mailRepository) Save(m *domainMail.Mail) error {
	err := i.conn.Create(&m).Error
	if err != nil {
		return err
	}
	return nil
}

func (i *mailRepository) Update(id int, token string) error {
	var m domainMail.Mail
	m.Token = token
	m.UserID = id
	m.Validation = true
	i.conn.Where("user_id = ?", id).First(&domainMail.Mail{})
	err := i.conn.Model(&m).Updates(domainMail.Mail{Token: token, Validation: true}).Error
	if err != nil {
		return err
	}
	return err
}

func (i *mailRepository) Find(id int) (domainMail.Mail, error) {
	var m domainMail.Mail
	err := i.conn.First(&m, "user_id = ?", id).Error
	if err != nil {
		return domainMail.Mail{}, err
	}
	return m, nil
}

func (i *mailRepository) Check(token string, id int) error {
	var m domainMail.Mail
	err := i.conn.First(&m, "user_id = ? AND token = ?", id, token).Error
	if err != nil {
		return err
	}
	return nil
}

func (i *mailRepository) Validation(id int) (bool, error) {
	var m domainMail.Mail
	err := i.conn.First(&m, "user_id = ?", id).Error
	if err != nil {
		return false, err
	}
	if m.Validation {
		return true, nil
	}
	return false, nil
}
