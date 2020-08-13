package user

type UserRepository interface {
	Save(*User) error
	Find(string) (User, error)
	FindEmail(string) bool
	CheckEmail(string) error
}
