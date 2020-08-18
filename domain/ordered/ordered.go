package ordered

type Ordered struct {
	ItemID int `json: "item_id"`
	UserID int `json: "user_id"`
	Number int `json: "number"`
}

func NewOrdered(i, u, num int) *Ordered {
	return &Ordered{
		ItemID: i,
		UserID: u,
		Number: num,
	}
}
