package token

import (
	"math/rand"
	"time"
)

func MakeToken() string {
	rand.Seed(time.Now().UnixNano())

	rs1Letters := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
	b := make([]rune, 100)

	for i := range b {
		b[i] = rs1Letters[rand.Intn(len(rs1Letters))]
	}

	return string(b)
}
