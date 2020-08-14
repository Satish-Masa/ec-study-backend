package infrastructure

import (
	"github.com/Satish-Masa/ec-backend/domain/item"
	"github.com/jinzhu/gorm"
)

type itemRepository struct {
	conn *gorm.DB
}

func NewItemRepository(conn *gorm.DB) item.ItemRepository {
	return &itemRepository{conn: conn}
}

func (i *itemRepository) Get() ([]item.Item, error) {
	var units []item.Item
	err := i.conn.Find(&units).Error
	if err != nil {
		return units, err
	}
	return units, nil
}

func (i *itemRepository) Find(id int) (item.Item, error) {
	var unit item.Item
	err := i.conn.First(&unit, "id = ?", id).Error
	if err != nil {
		return item.Item{}, err
	}
	return unit, nil
}
