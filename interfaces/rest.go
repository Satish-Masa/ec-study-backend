package interfaces

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	UserAuth "github.com/Satish-Masa/ec-backend/application/auth"
	AppCart "github.com/Satish-Masa/ec-backend/application/cart"
	passhash "github.com/Satish-Masa/ec-backend/application/hash"
	AppItem "github.com/Satish-Masa/ec-backend/application/item"
	"github.com/Satish-Masa/ec-backend/application/mail"
	AppUser "github.com/Satish-Masa/ec-backend/application/user"
	"github.com/Satish-Masa/ec-backend/config"
	domainCart "github.com/Satish-Masa/ec-backend/domain/cart"
	domainItem "github.com/Satish-Masa/ec-backend/domain/item"
	domainUser "github.com/Satish-Masa/ec-backend/domain/user"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

type Rest struct {
	UserRepository domainUser.UserRepository
	ItemRepository domainItem.ItemRepository
	CartRepository domainCart.CartRepository
}

func (r Rest) signupHandler(c echo.Context) error {
	req := new(AppUser.UserCreateRequest)
	if err := c.Bind(req); err != nil {
		log.Fatal(err)
		return &echo.HTTPError{
			Code:    http.StatusInternalServerError,
			Message: "failed bind request",
		}
	}

	pass, err := passhash.UserPassHash(req.Password)
	if err != nil {
		log.Fatal(err)
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
		return c.JSON(http.StatusBadRequest, resp)
	}

	err = application.SaveUser(u)
	if err != nil {
		log.Fatal(err)
		resp.Result = "failed to save user"
		return c.JSON(http.StatusInternalServerError, resp)
	}

	err = mail.SendEmail(req.Email)
	if err != nil {
		log.Println(err)
		resp.Result = "Not Send Email"
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
		return c.JSON(http.StatusBadRequest, resp)
	}

	ok := passhash.UserPassMach(u.Password, req.Password)
	if !ok {
		resp.Token = "not correct password"
	}

	token, err := UserAuth.FetchToken(&u)
	resp.Token = token

	return c.JSON(http.StatusOK, resp)
}

func (r Rest) getItemsHandler(c echo.Context) error {
	application := AppItem.ItemApplication{
		Repository: r.ItemRepository,
	}

	resp, err := application.GetItemList()
	if err != nil {
		log.Fatal(err)
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, resp)
}

func (r Rest) getItemHandler(c echo.Context) error {
	req := new(AppItem.ItemRequest)
	if err := c.Bind(req); err != nil {
		return err
	}

	id, err := strconv.Atoi(req.ID)
	if err != nil {
		return err
	}

	application := AppItem.ItemApplication{
		Repository: r.ItemRepository,
	}

	resp, err := application.FindItem(id)
	if err != nil {
		log.Println("not find id")
		return err
	}

	return c.JSON(http.StatusOK, resp)
}

func (r Rest) addCartHandler(c echo.Context) error {
	user := UserAuth.Check(c)

	item := new(domainItem.Item)
	req := new(AppItem.ItemRequest)
	if err := c.Bind(req); err != nil {
		return err
	}

	item.ID, _ = strconv.Atoi(req.ID)

	capp := AppCart.CartRepository{
		Repository: r.CartRepository,
	}

	err := capp.AddCart(*item, user)
	if err != nil {
		return &echo.HTTPError{
			Code:    http.StatusInternalServerError,
			Message: "Failed to add cart",
		}
	}

	return c.JSON(http.StatusOK, echo.Map{
		"Create": "OK",
	})
}

func (r Rest) CartListHandler(c echo.Context) error {
	user := UserAuth.Check(c)

	application := AppCart.CartRepository{
		Repository: r.CartRepository,
	}

	carts, err := application.GetCart(user.ID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, nil)
	}

	iapplication := AppItem.ItemApplication{
		Repository: r.ItemRepository,
	}
	var resp []domainItem.Item
	for _, cart := range carts {
		item, err := iapplication.FindItem(cart.ItemID)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, nil)
		}
		resp = append(resp, item)
	}

	log.Println(resp)

	return c.JSON(http.StatusOK, resp)
}

func (r Rest) Start() {
	e := echo.New()

	e.Use(middleware.CORS())
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.POST("/auth/signup", r.signupHandler)
	e.POST("/auth/login", r.loginHandler)
	e.GET("/items/list", r.getItemsHandler)
	e.POST("/item", r.getItemHandler)

	auth := e.Group("/auth")
	auth.Use(middleware.JWTWithConfig(UserAuth.Config))
	auth.POST("/item/add", r.addCartHandler)
	auth.POST("/cart", r.CartListHandler)

	e.Logger.Fatal(e.Start(fmt.Sprintf(":%d", config.Config.Port)))
}
