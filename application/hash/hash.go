package passhash

import "golang.org/x/crypto/bcrypt"

func UserPassHash(pass string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(pass), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}

func UserPassMach(hash, pass string) bool {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(pass)) == nil
}
