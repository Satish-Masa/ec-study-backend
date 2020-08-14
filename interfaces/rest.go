package interfaces

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	UserAuth "github.com/Satish-Masa/ec-backend/application/auth"
	"github.com/Satish-Masa/ec-backend/application/cart"
	AppCart "github.com/Satish-Masa/ec-backend/application/cart"
	passhash "github.com/Satish-Masa/ec-backend/application/hash"
	AppItem "github.com/Satish-Masa/ec-backend/application/item"
	"github.com/Satish-Masa/ec-backend/application/mail"
	"github.com/Satish-Masa/ec-backend/application/token"
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

	token := token.MakeToken()

	u := domainUser.NewUser(req.Email, pass, token)

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

	err = mail.SendEmail(req.Email, token)
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

func (r Rest) checkMailHandler(c echo.Context) error {
	req := new(AppUser.UserMailCheck)
	if err := c.Bind(req); err != nil {
		return err
	}

	application := AppUser.UserApplication{
		Repository: r.UserRepository,
	}

	err := application.CheckEmail(req.Token)
	if err != nil {
		return c.NoContent(http.StatusInternalServerError)
	}

	return c.NoContent(http.StatusOK)
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

	req := new(AppItem.ItemRequest)
	if err := c.Bind(req); err != nil {
		return err
	}

	iid, _ := strconv.Atoi(req.ID)

	capp := AppCart.CartRepository{
		Repository: r.CartRepository,
	}

	err := capp.AddCart(iid, user.ID, req.Number)
	if err != nil {
		return &echo.HTTPError{
			Code:    http.StatusInternalServerError,
			Message: "Failed to add cart",
		}
	}

	return c.NoContent(http.StatusOK)
}

func (r Rest) cartListHandler(c echo.Context) error {
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
	var resp []cart.CartResponce
	for _, cl := range carts {
		var r cart.CartResponce
		item, err := iapplication.FindItem(cl.ItemID)
		if err != nil {
			return c.NoContent(http.StatusInternalServerError)
		}
		r.Name = item.Name
		r.Description = item.Description
		r.Price = item.Price
		r.Stock = item.Stock
		r.Number = cl.Number
		resp = append(resp, r)
	}

	return c.JSON(http.StatusOK, resp)
}

func (r Rest) sendMailHandler(c echo.Context) error {
	req := new(AppUser.UserSendMail)
	if err := c.Bind(req); err != nil {
		return err
	}

	application := AppUser.UserApplication{
		Repository: r.UserRepository,
	}

	u, err := application.FindUser(req.Email)
	if err != nil {
		return c.NoContent(http.StatusInternalServerError)
	}

	ok := passhash.UserPassMach(u.Password, req.Password)
	if !ok {
		return c.NoContent(http.StatusInternalServerError)
	}

	err = mail.SendEmail(u.Email, u.Token)
	if err != nil {
		return c.NoContent(http.StatusInternalServerError)
	}

	return c.NoContent(http.StatusOK)
}

func (r Rest) Start() {
	e := echo.New()

	e.Use(middleware.CORS())
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.POST("/auth/signup", r.signupHandler)
	e.POST("/auth/login", r.loginHandler)
	e.POST("/auth/mailcheck", r.checkMailHandler)

	e.GET("/items/list", r.getItemsHandler)
	e.POST("/item", r.getItemHandler)

	e.POST("/send/mail", r.sendMailHandler)

	auth := e.Group("/auth")
	auth.Use(middleware.JWTWithConfig(UserAuth.Config))
	auth.POST("/item/add", r.addCartHandler)
	auth.POST("/cart", r.cartListHandler)

	e.Logger.Fatal(e.Start(fmt.Sprintf(":%d", config.Config.Port)))
}
