package infrastructure

import (
	domainCart "github.com/Satish-Masa/ec-backend/domain/cart"
	"github.com/Satish-Masa/ec-backend/domain/item"
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

	ok := c.conn.Where("user_id = ? AND item_id = ?", uid, iid).First(&domainCart.Cart{}).Error
	if ok != nil {
		err := c.conn.Create(&cart).Error
		if err != nil {
			return err
		}
	}

	err := c.conn.Model(&cart).Update("number", num).Error
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

func (c *cartRepository) Delete(uid, iid int) error {
	err := c.conn.Delete(&domainCart.Cart{}, "user_id = ? AND item_id = ?", uid, iid).Error
	if err != nil {
		return err
	}
	return nil
}

func (c *cartRepository) Set(uid int) error {
	var carts []domainCart.Cart
	err := c.conn.Find(&carts).Where("user_id = ?", uid).Error
	if err != nil {
		return err
	}
	for _, cart := range carts {
		ok := c.conn.First(&item.Item{}, "id = ? AND stock = ?", cart.ItemID, 0).Error
		if ok == nil {
			c.conn.Delete(&domainCart.Cart{}, "user_id = ? AND item_id = ?", uid, cart.ItemID)
		}
	}
	return nil
}
