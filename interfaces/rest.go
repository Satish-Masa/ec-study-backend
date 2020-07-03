package interfaces

import (
	"fmt"
	"net/http"

	UserAuth "github.com/Satish-Masa/ec-backend/application/auth"
	passhash "github.com/Satish-Masa/ec-backend/application/hash"
	AppUser "github.com/Satish-Masa/ec-backend/application/user"
	"github.com/Satish-Masa/ec-backend/config"
	domainUser "github.com/Satish-Masa/ec-backend/domain/user"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

type Rest struct {
	UserRepository domainUser.UserRepository
}

func (r Rest) signupHandler(c echo.Context) error {
	req := new(AppUser.UserCreateRequest)
	if err := c.Bind(req); err != nil {
		return &echo.HTTPError{
			Code:    http.StatusInternalServerError,
			Message: "failed bind request",
		}
	}

	pass, err := passhash.UserPassHash(req.Password)
	if err != nil {
		return err
	}

	u := domainUser.NewUser(req.Email, pass)

	application := AppUser.UserApplication{
		Repository: r.UserRepository,
	}

	resp := new(AppUser.UserCreateResponce)

	ok := application.FindEmail(req.Email)
	if ok {
		resp.Result = "find user email"
		return c.JSON(http.StatusInternalServerError, resp)
	}

	err = application.SaveUser(u)
	if err != nil {
		resp.Result = "failed to save user"
		return c.JSON(http.StatusInternalServerError, resp)
	}

	resp.Result = "success"

	return c.JSON(http.StatusOK, resp)
}

func (r Rest) loginHandler(c echo.Context) error {
	req := new(AppUser.UserLoginRequest)
	resp := new(AppUser.UserLoginResponce)
	if err := c.Bind(req); err != nil {
		return err
	}

	application := AppUser.UserApplication{
		Repository: r.UserRepository,
	}

	u, err := application.FindUser(req.Email)
	if err != nil {
		resp.Token = "not correct email"
		return c.JSON(http.StatusInternalServerError, resp)
	}

	ok := passhash.UserPassMach(u.Password, req.Password)
	if !ok {
		resp.Token = "not correct password"
	}

	token, err := UserAuth.FetchToken(&u)
	resp.Token = token

	return c.JSON(http.StatusOK, resp)
}

func (r Rest) Start() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.POST("/auth/signup", r.signupHandler)

	auth := e.Group("/auth")
	auth.Use(middleware.JWTWithConfig(UserAuth.Config))
	auth.POST("/login", r.loginHandler)
	e.Logger.Fatal(e.Start(fmt.Sprintf(":%d", config.Config.Port)))
}
