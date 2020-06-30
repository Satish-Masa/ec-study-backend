package user

import domainUser "github.com/Satish-Masa/ec-backend/domain/user"

type UserApplication struct {
	Repository domainUser.UserRepository
}

type UserCreateRequest struct {
	Name     string `json: "name"`
	Password string `json: "password"`
}

type UserCreateResponce struct {
	Token string `json: "token"`
}

type UserLoginRequest struct {
	Name     string `json: "name"`
	Password string `json: "password"`
}

type UserLoginResponce struct {
	Result string `json: "result"`
}

func (a UserApplication) SaveUser(u *domainUser.User) error {
	return a.Repository.Save(u)
}

func (a UserApplication) FindUser(id int) (domainUser.User, error) {
	return a.Repository.Find(id)
}
