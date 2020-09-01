package user

import "github.com/Satish-Masa/ec-backend/domain/user"

type UserApplication struct {
	Repository user.UserRepository
}

type UserCreateRequest struct {
	Email    string `json: "email"`
	Password string `json: "password"`
}

type UserCreateResponce struct {
	Result string `json: "result"`
}

type UserLoginRequest struct {
	Email    string `json: "email"`
	Password string `json: "password"`
}

type UserLoginResponce struct {
	Token string `json: "token"`
}

type UserSendMail struct {
	Email    string `json: "email"`
	Password string `json: "password"`
}

func (a UserApplication) SaveUser(u *user.User) error {
	return a.Repository.Save(u)
}

func (a UserApplication) FindUser(email string) (user.User, error) {
	return a.Repository.Find(email)
}
