package user

import domainUser "github.com/Satish-Masa/ec-backend/domain/user"

type UserApplication struct {
	Repository domainUser.UserRepository
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

func (a UserApplication) SaveUser(u *domainUser.User) error {
	return a.Repository.Save(u)
}

func (a UserApplication) FindUser(email string) (domainUser.User, error) {
	return a.Repository.Find(email)
}

func (a UserApplication) FindEmail(email string) bool {
	return a.Repository.FindEmail(email)
}
