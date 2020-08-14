package infrastructure

import (
	"github.com/Satish-Masa/ec-backend/domain/item"
	"github.com/Satish-Masa/ec-backend/domain/ordered"
	"github.com/jinzhu/gorm"
)

type orderedRepository struct {
	conn *gorm.DB
}

func NewOrderedRepository(conn *gorm.DB) ordered.OrderedRepository {
	return &orderedRepository{conn: conn}
}

func (i *orderedRepository) Add(iid, uid, num int) error {
	var o ordered.Ordered
	o.ItemID = iid
	o.UserID = uid
	o.Number = num

	tx := i.conn.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Create(&o).Error; err != nil {
		tx.Rollback()
		return err
	}

	var unit item.Item
	if err := tx.First(&unit, "id=?", o.ItemID).Error; err != nil {
		tx.Rollback()
		return err
	}

	stock := unit.Stock - o.Number
	if err := tx.Where(&unit).Update("stock", stock).Error; err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}
