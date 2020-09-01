package ordered

import (
	"time"

	"github.com/Satish-Masa/ec-backend/domain/ordered"
)

type OrderedRepository struct {
	Repository ordered.OrderedRepository
}

type OrderedResponce struct {
	ID          int       `json: "id"`
	Created_at  time.Time `json: "created_at"`
	Name        string    `json: "name"`
	Description string    `json: "description"`
	Price       int       `json: "price"`
	Number      int       `json: "number"`
}

func (a OrderedRepository) AddOrdered(iid, uid, num int) error {
	return a.Repository.Add(iid, uid, num)
}

func (a OrderedRepository) GetOrdered(id int) ([]ordered.Ordered, error) {
	return a.Repository.Get(id)
}
