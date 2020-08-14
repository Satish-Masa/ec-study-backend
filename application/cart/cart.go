package cart

import (
	domainCart "github.com/Satish-Masa/ec-backend/domain/cart"
)

type CartRepository struct {
	Repository domainCart.CartRepository
}

type CartResponce struct {
	Name        string `json: "name"`
	Description string `json: "description"`
	Price       int    `json: "price"`
	Stock       int    `json: "stock"`
	Number      int    `json: "number"`
}

func (a CartRepository) AddCart(iid, uid, num int) error {
	return a.Repository.Add(iid, uid, num)
}

func (a CartRepository) GetCart(uid int) ([]domainCart.Cart, error) {
	return a.Repository.Get(uid)
}
