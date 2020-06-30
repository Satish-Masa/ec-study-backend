package interfaces

import (
	"net/http"
	"time"

	"github.com/Satish-Masa/ec-backend/application/user"
	domainUser "github.com/Satish-Masa/ec-backend/domain/user"
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

type jwtCustomClaims struct {
	UID  int    `json: "uid"`
	Name string `json: "name"`
	jwt.StandardClaims
}

var signingKey = []byte("secret")

var Config = middleware.JWTConfig{
	Claims:     &jwtCustomClaims{},
	SigningKey: signingKey,
}

func createToken(u *domainUser.User) (user.UserCreateResponce, error) {
	if u.Name == "" {
		return user.UserCreateResponce{}, &echo.HTTPError{
			Code:    http.StatusUnauthorized,
			Message: "invalid name",
		}
	}

	claims := &jwtCustomClaims{
		u.ID,
		u.Name,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 72).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	t, err := token.SignedString(signingKey)
	if err != nil {
		return user.UserCreateResponce{}, &echo.HTTPError{
			Code:    http.StatusInternalServerError,
			Message: "failed to create the token",
		}
	}

	resp := new(user.UserCreateResponce)
	resp.Token = t

	return *resp, nil
}

func FetchToken(u *domainUser.User) (resp user.UserCreateResponce, err error) {
	resp, err = createToken(u)
	if err != nil {
		return user.UserCreateResponce{}, err
	}

	return resp, nil
}

func FindUserID(c echo.Context) int {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*jwtCustomClaims)
	uid := claims.UID
	return uid
}
