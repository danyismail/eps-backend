package store

import (
	"eps-backend/model"

	"gorm.io/gorm"
)

type SalesConstruct struct {
	db *gorm.DB
}

func NewSalesStore(db *gorm.DB) *SalesConstruct {
	return &SalesConstruct{
		db: db,
	}
}

func (c *SalesConstruct) GetSalesToday() ([]model.SalesReport, error) {
	salesToday := []model.SalesReport{}
	sql := "SELECT * FROM v_today_sales"
	if err := c.db.Debug().Raw(sql).Scan(&salesToday).Error; err != nil {
		return nil, err
	}
	return salesToday, nil
}
