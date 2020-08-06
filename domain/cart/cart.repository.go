package cart

import (
	"github.com/Satish-Masa/ec-backend/domain/item"
	"github.com/Satish-Masa/ec-backend/domain/user"
)

type CartRepository interface {
	Add(item.Item, user.User) error
	Get(int) ([]Cart, error)
}
