package user

type UserRepository interface {
	Save(*User) error
	Find(string) (User, error)
}
