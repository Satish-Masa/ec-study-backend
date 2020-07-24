package item

import domainItem "github.com/Satish-Masa/ec-backend/domain/item"

type ItemApplication struct {
	Repository domainItem.ItemRepository
}

func (a ItemApplication) GetItemList() ([]domainItem.Item, error) {
	return a.Repository.Get()
}
