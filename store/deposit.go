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

func (c *DepositConstruct) GetBalance(path string, periode ...string) ([]model.CurrentDeposit, error) {
	// Set default periode if not provided
	var sql string
	p := "today"
	if len(periode) > 0 {
		p = periode[0]
	}

	balance := []model.CurrentDeposit{}
	switch p {
	case "yesterday":
		sql = `
		SELECT 
			kode_modul,
			m.label,
			count(1) AS total_transaksi,
			sum(harga_beli) AS pemakaian_saldo,
			m.saldo AS saldo_sekarang
		FROM 
			transaksi t
		JOIN modul m ON
			t.kode_modul = m.kode
		WHERE 
			t.tgl_entri >= CAST(DATEADD(DAY, -1, GETDATE()) AS DATE)
			AND t.tgl_entri < CAST(GETDATE() AS DATE)
			AND t.status = 20
		GROUP BY
			t.kode_modul,
			m.label,
			m.saldo
		ORDER BY m.label ASC;`
	default: // default to "today"
		sql = `
		SELECT
			kode_modul,
			m.label,
			count(1) AS total_transaksi,
			sum(harga_beli) AS pemakaian_saldo,
			m.saldo AS saldo_sekarang
		FROM
			transaksi t
		JOIN modul m ON
			t.kode_modul = m.kode
		WHERE
			t.tgl_entri BETWEEN CONVERT(DATETIME, CONVERT(DATE, GETDATE()))
			AND CAST(GETDATE() AS DATETIME)
			AND t.status = 20
		GROUP BY
			t.kode_modul,
			m.label,
			m.saldo
		ORDER BY m.label ASC;`
	}

	if err := utils.SelectConn(path, c.db).Raw(sql).Scan(&balance).Error; err != nil {
		return nil, err
	}
	return balance, nil
}
