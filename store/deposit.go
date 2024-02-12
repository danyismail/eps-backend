package store

import (
	"eps-backend/model"

	"gorm.io/gorm"
)

type DepositConstruct struct {
	dev  *gorm.DB
	prod *gorm.DB
}

func NewDepositStore(dev *gorm.DB, prod *gorm.DB) *DepositConstruct {
	return &DepositConstruct{
		dev,
		prod,
	}
}

func (c *DepositConstruct) GetBalanceToday() ([]model.CurrentDeposit, error) {
	depositToday := []model.CurrentDeposit{}
	sql := "SELECT * FROM v_today_supplier_balances"
	if err := c.dev.Debug().Raw(sql).Scan(&depositToday).Error; err != nil {
		return nil, err
	}
	return depositToday, nil
}

func (c *DepositConstruct) GetBalanceTodayProd() ([]model.CurrentDeposit, error) {
	depositToday := []model.CurrentDeposit{}
	sql := "SELECT * FROM v_today_supplier_balances"
	if err := c.prod.Debug().Raw(sql).Scan(&depositToday).Error; err != nil {
		return nil, err
	}
	return depositToday, nil
}
