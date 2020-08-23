package item

import "github.com/Satish-Masa/ec-backend/domain/item"

type ItemRepository struct {
	Repository item.ItemRepository
}

type ItemRequest struct {
	ID     string `json: "id"`
	Number int    `json: "number"`
}

func (a ItemRepository) GetItemList() ([]item.Item, error) {
	return a.Repository.Get()
}

func (a ItemRepository) FindItem(id int) (item.Item, error) {
	return a.Repository.Find(id)
}
