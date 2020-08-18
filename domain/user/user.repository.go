package user

type UserRepository interface {
	Save(*User) error
	Find(string) (User, error)
	CheckEmail(string) error
	Validation(int) (bool, error)
}
