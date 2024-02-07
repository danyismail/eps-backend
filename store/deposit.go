package store

import (
	"eps-backend/model"

	"gorm.io/gorm"
)

type DepositConstruct struct {
	db *gorm.DB
}

func NewDepositStore(db *gorm.DB) *DepositConstruct {
	return &DepositConstruct{
		db: db,
	}
}

func (c *DepositConstruct) GetBalanceToday() ([]model.CurrentDeposit, error) {
	depositToday := []model.CurrentDeposit{}
	sql := "SELECT * FROM v_today_supplier_balances"
	if err := c.db.Debug().Raw(sql).Scan(&depositToday).Error; err != nil {
		return nil, err
	}
	return depositToday, nil
}
