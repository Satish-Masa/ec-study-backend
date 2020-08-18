package cart

import (
	"github.com/Satish-Masa/ec-backend/domain/cart"
)

type CartRepository struct {
	Repository cart.CartRepository
}

type CartResponce struct {
	ID          int    `json: "id"`
	Name        string `json: "name"`
	Description string `json: "description"`
	Price       int    `json: "price"`
	Stock       int    `json: "stock"`
	Number      int    `json: "number"`
}

type CartDeleteRequest struct {
	ID int `json: "id"`
}

func (a CartRepository) AddCart(iid, uid, num int) error {
	return a.Repository.Add(iid, uid, num)
}

func (a CartRepository) GetCart(uid int) ([]cart.Cart, error) {
	return a.Repository.Get(uid)
}

func (a CartRepository) DeleteCart(uid, iid int) error {
	return a.Repository.Delete(uid, iid)
}
