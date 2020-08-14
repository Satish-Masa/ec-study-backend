package user

type User struct {
	ID         int    `json: "id" gorm: "praimaly_key"`
	Email      string `json: "email"`
	Password   string `json: "password"`
	Token      string `json: "token"`
	Validation bool   `json: "validation`
}

func NewUser(name, pass, token string) *User {
	return &User{
		Email:    name,
		Password: pass,
		Token:    token,
	}
}
