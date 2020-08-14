package ordered

import "github.com/Satish-Masa/ec-backend/domain/ordered"

type OrderedRepository struct {
	Repository ordered.OrderedRepository
}

func (a OrderedRepository) AddOrdered(iid, uid, num int) error {
	return a.Repository.Add(iid, uid, num)
}
