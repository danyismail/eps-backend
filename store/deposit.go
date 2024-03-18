package store

import (
	"eps-backend/db"
	"eps-backend/model"
	"eps-backend/utils"

	"gorm.io/gorm"
)

type DepositConstruct struct {
	db db.DBConnection
}

func NewDepositStore(db db.DBConnection) *DepositConstruct {
	return &DepositConstruct{
		db,
	}
}

func (c *DepositConstruct) GetBalanceToday() ([]model.CurrentDeposit, error) {
	depositToday := []model.CurrentDeposit{}
	sql := "SELECT * FROM v_today_supplier_balances"
	if err := c.db.DigiEps.Debug().Raw(sql).Scan(&depositToday).Error; err != nil {
		return nil, err
	}
	return depositToday, nil
}

func (c *DepositConstruct) GetBalanceTodayProd() ([]model.CurrentDeposit, error) {
	depositToday := []model.CurrentDeposit{}
	sql := "SELECT * FROM v_today_supplier_balances"
	if err := c.db.DigiAmazone.Debug().Raw(sql).Scan(&depositToday).Error; err != nil {
		return nil, err
	}
	return depositToday, nil
}

func (c *DepositConstruct) GetBalance(conn string) ([]model.CurrentDeposit, error) {
	depositToday := []model.CurrentDeposit{}
	sql := `
	select
		kode_modul,
		m.label,
		count(1) as total_transaksi,
		sum(harga_beli) as pemakaian_saldo ,
		m.saldo as saldo_sekarang
	from
		transaksi t
	join modul m on
		t.kode_modul = m.kode
	WHERE
		t.tgl_entri BETWEEN (CONVERT(DATETIME,
		CONVERT(DATE,
		GETDATE()))) AND (
		SELECT
			CAST(GETDATE() AS DATETIME))
	group by
		t.kode_modul,
		m.label,
		m.saldo
	`
	if err := c.SwitchDB(conn).Debug().Raw(sql).Scan(&depositToday).Error; err != nil {
		return nil, err
	}
	return depositToday, nil
}

func (c *DepositConstruct) SwitchDB(path string) *gorm.DB {
	switch path {
	case utils.AMAZONE:
		return c.db.Amazone
	case utils.EPS:
		return c.db.Eps
	default:
		return c.db.Eps
	}
}
