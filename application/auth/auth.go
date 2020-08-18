package auth

import (
	"time"

	"github.com/Satish-Masa/ec-backend/domain/user"
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

var secret = "2FMd5FNSqS/nW2wWJy5S3ppjSHhUnLt8HuwBkTD6HqfPfBBDlykwLA=="

type jwtCustomClaims struct {
	ID    int    `json:"id"`
	Email string `json:"email"`
	jwt.StandardClaims
}

var Config = middleware.JWTConfig{
	Claims:     &jwtCustomClaims{},
	SigningKey: []byte(secret),
}

func createToken(u *user.User) (string, error) {
	claims := &jwtCustomClaims{
		u.ID,
		u.Email,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 72).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func FetchToken(u *user.User) (resp string, err error) {
	resp, err = createToken(u)
	if err != nil {
		return "", err
	}

	return resp, nil
}

func Check(c echo.Context) user.User {
	var user user.User
	u := c.Get("user").(*jwt.Token)
	claims := u.Claims.(*jwtCustomClaims)
	user.ID = claims.ID
	user.Email = claims.Email
	return user
}
