package mail

import "time"

type Mail struct {
	UserID     int       `json: "user_id"`
	Created_at time.Time `json: "created_at"`
	Token      string    `json: "token"`
	Validation bool      `json: "validation"`
}

func NewMail(id int, token string) *Mail {
	return &Mail{
		UserID: id,
		Token:  token,
	}
}
