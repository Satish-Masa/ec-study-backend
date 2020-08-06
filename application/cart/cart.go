package cart

import (
	domainCart "github.com/Satish-Masa/ec-backend/domain/cart"
	"github.com/Satish-Masa/ec-backend/domain/item"
	"github.com/Satish-Masa/ec-backend/domain/user"
)

type CartRepository struct {
	Repository domainCart.CartRepository
}

func (a CartRepository) AddCart(item item.Item, user user.User) error {
	return a.Repository.Add(item, user)
}

func (a CartRepository) GetCart(uid int) ([]domainCart.Cart, error) {
	return a.Repository.Get(uid)
}
