package mail

type MailRepository interface {
	Save(*Mail) error
	Update(int, string) error
	Find(int) (Mail, error)
	Check(string, int) error
	Validation(int) (bool, error)
}
