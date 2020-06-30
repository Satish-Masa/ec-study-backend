package infrastructure

import (
	domainUser "github.com/Satish-Masa/ec-backend/domain/user"
	"github.com/jinzhu/gorm"
)

type userRepository struct {
	conn *gorm.DB
}

func NewUserRepository(conn *gorm.DB) domainUser.UserRepository {
	return &userRepository{conn: conn}
}

func (i *userRepository) Save(u *domainUser.User) error {
	err := i.conn.Create(&u).Error
	if err != nil {
		return err
	}
	return nil
}

func (i *userRepository) Find(id int) (domainUser.User, error) {
	var user domainUser.User
	err := i.conn.First(&user, "id = ?", id).Error
	if err != nil {
		return domainUser.User{}, err
	}
	return user, nil
}
