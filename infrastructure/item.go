package infrastructure

import (
	domainItem "github.com/Satish-Masa/ec-backend/domain/item"
	"github.com/jinzhu/gorm"
)

type itemRepository struct {
	conn *gorm.DB
}

func NewItemRepository(conn *gorm.DB) domainItem.ItemRepository {
	return &itemRepository{conn: conn}
}

func (i *itemRepository) Get() ([]domainItem.Item, error) {
	var items []domainItem.Item
	err := i.conn.Find(&items).Error
	if err != nil {
		return items, err
	}
	return items, nil
}

func (i *itemRepository) Find(id int) (domainItem.Item, error) {
	var item domainItem.Item
	err := i.conn.First(&item, "id = ?", id).Error
	if err != nil {
		return domainItem.Item{}, err
	}
	return item, nil
}
