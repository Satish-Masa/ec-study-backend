package interfaces

import (
	"net/http"

	AppUser "github.com/Satish-Masa/ec-backend/application/user"
	domainUser "github.com/Satish-Masa/ec-backend/domain/user"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

type Rest struct {
	UserRepository domainUser.UserRepository
}

func (r Rest) createHandler(c echo.Context) error {
	req := new(AppUser.UserCreateRequest)
	if err := c.Bind(req); err != nil {
		return &echo.HTTPError{
			Code:    http.StatusInternalServerError,
			Message: "failed bind request",
		}
	}

	pass, err := UserPassHash(req.Password)
	if err != nil {
		return err
	}

	u := domainUser.NewUser(req.Name, pass)

	application := AppUser.UserApplication{
		Repository: r.UserRepository,
	}

	err = application.SaveUser(u)

	resp, err := FetchToken(u)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, resp)
}

func (r Rest) findHandler(c echo.Context) error {
	req := new(AppUser.UserLoginRequest)
	if err := c.Bind(req); err != nil {
		return nil
	}

	uid := FindUserID(c)

	application := AppUser.UserApplication{
		Repository: r.UserRepository,
	}

	u, err := application.FindUser(uid)
	if err != nil {
		return err
	}

	ok := UserPassMach(u.Password, req.Password)

	resp := new(AppUser.UserLoginResponce)
	if ok {
		resp.Result = "succese"
	} else {
		resp.Result = "failed"
	}

	return c.JSON(http.StatusOK, resp)
}

func (r Rest) Start() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.POST("/user/create", r.createHandler)

	auth := e.Group("/auth")
	auth.Use(middleware.JWTWithConfig(Config))
	auth.POST("/login", r.findHandler)
}
