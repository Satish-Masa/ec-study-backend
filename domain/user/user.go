package user

type User struct {
	ID       int    `json: "id" gorm: "praimaly_key"`
	Email    string `json: "email"`
	Password string `json: "password"`
}

func NewUser(name, pass string) *User {
	return &User{
		Email:    name,
		Password: pass,
	}
}
