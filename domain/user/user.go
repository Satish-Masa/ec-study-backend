package user

type User struct {
	ID       int    `json: "id" gorm: "praimaly_key"`
	Name     string `json: "name"`
	Password string `json: "password"`
}

func NewUser(name, pass string) *User {
	return &User{
		Name:     name,
		Password: pass,
	}
}
