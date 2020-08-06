package cart

import (
	"github.com/Satish-Masa/ec-backend/domain/item"
	"github.com/Satish-Masa/ec-backend/domain/user"
)

type Cart struct {
	ID     int `json: "id" gorm: "praimaly_key"`
	ItemID int `json: "item_id"`
	UserID int `json: "user_id"`
}

func NewCart(i item.Item, u user.User) *Cart {
	return &Cart{
		ItemID: i.ID,
		UserID: u.ID,
	}
}
