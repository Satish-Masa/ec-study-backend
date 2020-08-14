package item

import "github.com/Satish-Masa/ec-backend/domain/item"

type ItemApplication struct {
	Repository item.ItemRepository
}

type ItemRequest struct {
	ID     string `json: "id"`
	Number int    `json: "number"`
}

func (a ItemApplication) GetItemList() ([]item.Item, error) {
	return a.Repository.Get()
}

func (a ItemApplication) FindItem(id int) (item.Item, error) {
	return a.Repository.Find(id)
}
