package interfaces

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/Satish-Masa/ec-backend/application/auth"
	"github.com/Satish-Masa/ec-backend/application/cart"
	AppCart "github.com/Satish-Masa/ec-backend/application/cart"
	hash "github.com/Satish-Masa/ec-backend/application/hash"
	AppItem "github.com/Satish-Masa/ec-backend/application/item"
	"github.com/Satish-Masa/ec-backend/application/mail"
	AppOrdered "github.com/Satish-Masa/ec-backend/application/ordered"
	"github.com/Satish-Masa/ec-backend/application/token"
	AppUser "github.com/Satish-Masa/ec-backend/application/user"
	"github.com/Satish-Masa/ec-backend/config"
	domainCart "github.com/Satish-Masa/ec-backend/domain/cart"
	domainItem "github.com/Satish-Masa/ec-backend/domain/item"
	domainOrdered "github.com/Satish-Masa/ec-backend/domain/ordered"
	domainUser "github.com/Satish-Masa/ec-backend/domain/user"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

type Rest struct {
	UserRepository  domainUser.UserRepository
	ItemRepository  domainItem.ItemRepository
	CartRepository  domainCart.CartRepository
	OrderRepository domainOrdered.OrderedRepository
}

/* -------------------------------------------------------

						User

------------------------------------------------------- */
func (r Rest) signupHandler(c echo.Context) error {
	req := new(AppUser.UserCreateRequest)
	if err := c.Bind(req); err != nil {
		log.Fatal(err)
		return &echo.HTTPError{
			Code:    http.StatusInternalServerError,
			Message: "failed bind request",
		}
	}

	pass, err := hash.UserPassHash(req.Password)
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

	_, err = application.FindUser(req.Email)
	if err == nil {
		resp.Result = "find user email"
		log.Println(err)
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

	ok := hash.UserPassMach(u.Password, req.Password)
	if !ok {
		resp.Token = "not correct password"
	}

	token, err := auth.FetchToken(&u)
	resp.Token = token

	return c.JSON(http.StatusOK, resp)
}

/* -------------------------------------------------------

						Item

------------------------------------------------------- */
func (r Rest) getItemsHandler(c echo.Context) error {
	application := AppItem.ItemRepository{
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

	application := AppItem.ItemRepository{
		Repository: r.ItemRepository,
	}

	resp, err := application.FindItem(id)
	if err != nil {
		log.Println("not find id")
		return err
	}

	return c.JSON(http.StatusOK, resp)
}

/* -------------------------------------------------------

						Cart

------------------------------------------------------- */
func (r Rest) addCartHandler(c echo.Context) error {
	user := auth.Check(c)

	req := new(AppItem.ItemRequest)
	if err := c.Bind(req); err != nil {
		return err
	}

	iid, _ := strconv.Atoi(req.ID)

	cart_application := AppCart.CartRepository{
		Repository: r.CartRepository,
	}

	err := cart_application.AddCart(iid, user.ID, req.Number)
	if err != nil {
		return &echo.HTTPError{
			Code:    http.StatusInternalServerError,
			Message: "Failed to add cart",
		}
	}

	return c.NoContent(http.StatusOK)
}

func (r Rest) cartListHandler(c echo.Context) error {
	user := auth.Check(c)

	application := AppCart.CartRepository{
		Repository: r.CartRepository,
	}

	carts, err := application.GetCart(user.ID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, nil)
	}

	if len(carts) == 0 {
		log.Println("hello")
		return c.NoContent(http.StatusNoContent)
	}

	item_application := AppItem.ItemRepository{
		Repository: r.ItemRepository,
	}
	var resp []cart.CartResponce
	for _, cl := range carts {
		var r cart.CartResponce
		item, err := item_application.FindItem(cl.ItemID)
		if err != nil {
			return c.NoContent(http.StatusInternalServerError)
		}
		r.ID = item.ID
		r.Name = item.Name
		r.Description = item.Description
		r.Price = item.Price
		r.Stock = item.Stock
		r.Number = cl.Number
		resp = append(resp, r)
	}

	return c.JSON(http.StatusOK, resp)
}

func (r Rest) deleteCartHandler(c echo.Context) error {
	user := auth.Check(c)

	req := new(AppCart.CartDeleteRequest)
	if err := c.Bind(req); err != nil {
		return c.NoContent(http.StatusInternalServerError)
	}

	cart_application := AppCart.CartRepository{
		Repository: r.CartRepository,
	}

	err := cart_application.DeleteCart(user.ID, req.ID)
	if err != nil {
		return c.NoContent(http.StatusInternalServerError)
	}

	return c.NoContent(http.StatusOK)
}

func (r Rest) orderedHandler(c echo.Context) error {
	user := auth.Check(c)

	cart_application := AppCart.CartRepository{
		Repository: r.CartRepository,
	}

	err := cart_application.SetCart(user.ID)
	if err != nil {
		return c.NoContent(http.StatusInternalServerError)
	}

	return c.NoContent(http.StatusOK)
}

/* -------------------------------------------------------

						buy

------------------------------------------------------- */
func (r Rest) buyHandler(c echo.Context) error {
	user := auth.Check(c)

	cart_application := AppCart.CartRepository{
		Repository: r.CartRepository,
	}
	ordered_application := AppOrdered.OrderedRepository{
		Repository: r.OrderRepository,
	}

	carts, err := cart_application.GetCart(user.ID)
	if err != nil {
		return c.NoContent(http.StatusInternalServerError)
	}

	for _, cart := range carts {
		err := ordered_application.AddOrdered(cart.ItemID, cart.UserID, cart.Number)
		if err != nil {
			return c.NoContent(http.StatusInternalServerError)
		}
	}

	return c.NoContent(http.StatusOK)

}

func (r Rest) buyListHandler(c echo.Context) error {
	user := auth.Check(c)

	ordered_application := AppOrdered.OrderedRepository{
		Repository: r.OrderRepository,
	}
	item_application := AppItem.ItemRepository{
		Repository: r.ItemRepository,
	}

	orders, err := ordered_application.GetOrdered(user.ID)
	if err != nil {
		return c.NoContent(http.StatusInternalServerError)
	}

	var resp []AppOrdered.OrderedResponce
	for _, cl := range orders {
		var r AppOrdered.OrderedResponce
		item, err := item_application.FindItem(cl.ItemID)
		if err != nil {
			return c.NoContent(http.StatusInternalServerError)
		}
		r.ID = item.ID
		r.Name = item.Name
		r.Description = item.Description
		r.Price = item.Price
		r.Number = cl.Number
		r.Created_at = cl.Created_at
		resp = append(resp, r)
	}

	return c.JSON(http.StatusOK, resp)
}

/* -------------------------------------------------------

						Others

------------------------------------------------------- */
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

	ok := hash.UserPassMach(u.Password, req.Password)
	if !ok {
		return c.NoContent(http.StatusInternalServerError)
	}

	err = mail.SendEmail(u.Email, u.Token)
	if err != nil {
		return c.NoContent(http.StatusInternalServerError)
	}

	return c.NoContent(http.StatusOK)
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

func (r Rest) validationMailHandler(c echo.Context) error {
	user := auth.Check(c)

	user_application := AppUser.UserApplication{
		Repository: r.UserRepository,
	}

	ok, err := user_application.Validation(user.ID)
	if err != nil {
		return c.NoContent(http.StatusInternalServerError)
	}

	if ok {
		return c.String(http.StatusOK, "true")
	}

	return c.String(http.StatusOK, "false")
}

func (r Rest) Start() {
	e := echo.New()
	a := e.Group("/auth")
	a.Use(middleware.JWTWithConfig(auth.Config))

	e.Use(middleware.CORS())
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	/* -----------------user------------------ */
	e.POST("/auth/signup", r.signupHandler)
	e.POST("/auth/login", r.loginHandler)

	/* -----------------item------------------ */
	e.GET("/items/list", r.getItemsHandler)
	e.POST("/item", r.getItemHandler)

	/* -----------------cart------------------ */
	a.POST("/item/add", r.addCartHandler)
	a.POST("/cart", r.cartListHandler)
	a.POST("/cart/delete", r.deleteCartHandler)
	a.POST("/cart/ordered", r.orderedHandler)

	/* -----------------buy------------------ */
	a.POST("/buy", r.buyHandler)
	a.POST("/ordered", r.buyListHandler)

	/* -----------------other------------------ */
	e.POST("/auth/mailcheck", r.checkMailHandler)
	e.POST("/send/mail", r.sendMailHandler)
	a.POST("/varidation", r.validationMailHandler)

	e.Logger.Fatal(e.Start(fmt.Sprintf(":%d", config.Config.Port)))
}
