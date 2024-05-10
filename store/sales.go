package store

import (
	"eps-backend/db"
	"eps-backend/model"
	"eps-backend/utils"
	"fmt"
)

type SalesConstruct struct {
	db db.DBConnection
}

func NewSalesStore(db db.DBConnection) *SalesConstruct {
	return &SalesConstruct{db}
}

func (c *SalesConstruct) GetSales(path string) ([]model.SalesReport, error) {
	salesToday := []model.SalesReport{}
	sql := `
	select
	count(1) as trx,
	sum(harga_beli) as pembelian,
	sum(harga) as penjualan,
	sum(harga) - sum(harga_beli)as laba
from
	transaksi t
WHERE
	status = 20
	and t.tgl_entri BETWEEN (CONVERT(DATETIME,
	CONVERT(DATE,
	GETDATE()))) AND (
	SELECT
		CAST(GETDATE() AS DATETIME));
	`
	if err := utils.SelectConn(path, c.db).Raw(sql).Scan(&salesToday).Error; err != nil {
		return nil, err
	}
	return salesToday, nil
}

func (c *SalesConstruct) GetPPH(path string) ([]model.SalesReport, error) {
	pphOfSalesToday := []model.SalesReport{}
	sql := `
		SELECT 
			COUNT(1) AS trx,
			SUM(harga_beli) AS pembelian,
			SUM(harga) AS penjualan,
			SUM(harga) - SUM(harga_beli) AS laba,
			SUM(harga) * 0.005 AS pph
		FROM
			transaksi t
		WHERE
				status = 20
			AND kode_reseller  in (
			'EPS0634',
			'EPS0695',
			'EPS0840',
			'EPS0921',
			'EPS0712',
			'EPS6890',
			'EPS6973',
			'EPS6957',
			'EPS6995',
			'EPS0935'
		) AND t.tgl_entri BETWEEN (CONVERT(DATETIME,
			CONVERT(DATE,
			GETDATE()))) AND (
			SELECT
				CAST(GETDATE() AS DATETIME));
	`
	if err := utils.SelectConn(path, c.db).Debug().Raw(sql).Scan(&pphOfSalesToday).Error; err != nil {
		return nil, err
	}
	return pphOfSalesToday, nil
}

func (c *SalesConstruct) GetSalesPeriode(path, from, to string) ([]model.SalesReport, error) {
	salesToday := []model.SalesReport{}
	sql := "SELECT COUNT(1) AS trx, SUM(t.harga_beli) AS pembelian, SUM(t.harga) AS penjualan ,SUM(t.harga) - SUM(t.harga_beli) AS laba"
	sql = fmt.Sprintf("%s FROM transaksi t", sql)
	sql = fmt.Sprintf("%s WHERE t.status = 20", sql)
	sql = fmt.Sprintf("%s AND cast(t.tgl_entri AS date) BETWEEN '%s' AND '%s';", sql, from, to)
	if err := utils.SelectConn(path, c.db).Raw(sql).Debug().Scan(&salesToday).Error; err != nil {
		return nil, err
	}
	return salesToday, nil
}
