package item

import domainItem "github.com/Satish-Masa/ec-backend/domain/item"

type ItemApplication struct {
	Repository domainItem.ItemRepository
}

type ItemRequest struct {
	ID string `json: "id"`
}

func (a ItemApplication) GetItemList() ([]domainItem.Item, error) {
	return a.Repository.Get()
}

func (a ItemApplication) FindItem(id int) (domainItem.Item, error) {
	return a.Repository.Find(id)
}
