package cart

type Cart struct {
	ID     int `json: "id" gorm: "praimaly_key"`
	ItemID int `json: "item_id"`
	UserID int `json: "user_id"`
	Number int `json: "number"`
}

func NewCart(i, u, num int) *Cart {
	return &Cart{
		ItemID: i,
		UserID: u,
		Number: num,
	}
}
