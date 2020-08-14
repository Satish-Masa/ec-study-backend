package item

type Item struct {
	ID          int    `json: "id" gorm: "praimaly_key"`
	Name        string `json: "name"`
	Description string `json: "description"`
	Price       int    `json: "price"`
	Stock       int    `json: "stock"`
}

func NewItem(name, description string, price, stock int) *Item {
	return &Item{
		Name:        name,
		Description: description,
		Price:       price,
		Stock:       stock,
	}
}
