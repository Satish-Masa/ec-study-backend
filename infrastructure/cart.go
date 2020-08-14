package infrastructure

import (
	domainCart "github.com/Satish-Masa/ec-backend/domain/cart"
	"github.com/jinzhu/gorm"
)

type cartRepository struct {
	conn *gorm.DB
}

func NewCartRepository(conn *gorm.DB) domainCart.CartRepository {
	return &cartRepository{conn: conn}
}

func (c *cartRepository) Add(iid, uid, num int) error {
	var cart domainCart.Cart
	cart.ItemID = iid
	cart.UserID = uid
	cart.Number = num
	err := c.conn.Create(&cart).Error
	if err != nil {
		return err
	}
	return nil
}

func (c *cartRepository) Get(uid int) ([]domainCart.Cart, error) {
	var carts []domainCart.Cart
	err := c.conn.Find(&carts).Where("user_id = ?", uid).Error
	if err != nil {
		return carts, err
	}
	return carts, nil
}
