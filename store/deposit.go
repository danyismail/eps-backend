package store

import (
	"eps-backend/db"
	"eps-backend/model"
	"eps-backend/utils"
)

type DepositConstruct struct {
	db db.DBConnection
}

func NewDepositStore(db db.DBConnection) *DepositConstruct {
	return &DepositConstruct{
		db,
	}
}

func (c *DepositConstruct) GetBalance(path string) ([]model.CurrentDeposit, error) {
	balance := []model.CurrentDeposit{}
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
	order by m.label asc;
	`
	if err := utils.SelectConn(path, c.db).Raw(sql).Scan(&balance).Error; err != nil {
		return nil, err
	}
	return balance, nil
}
