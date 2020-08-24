package ordered

import "time"

type Ordered struct {
	Created_at time.Time `json: "created_at"`
	ItemID     int       `json: "item_id"`
	UserID     int       `json: "user_id"`
	Number     int       `json: "number"`
}

func NewOrdered(i, u, num int) *Ordered {
	return &Ordered{
		ItemID: i,
		UserID: u,
		Number: num,
	}
}
