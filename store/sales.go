package store

import (
	"eps-backend/db"
	"eps-backend/model"
	"eps-backend/utils"
	"fmt"

	"gorm.io/gorm"
)

type SalesConstruct struct {
	db db.DBConnection
}

func NewSalesStore(db db.DBConnection) *SalesConstruct {
	return &SalesConstruct{db}
}

func (c *SalesConstruct) GetSalesToday() ([]model.SalesReport, error) {
	salesToday := []model.SalesReport{}
	sql := "SELECT * FROM v_today_sales"
	if err := c.db.DigiEps.Raw(sql).Scan(&salesToday).Error; err != nil {
		return nil, err
	}
	return salesToday, nil
}

func (c *SalesConstruct) GetSalesTodayProd() ([]model.SalesReport, error) {
	salesToday := []model.SalesReport{}
	sql := "SELECT * FROM v_today_sales"
	if err := c.db.DigiAmazone.Raw(sql).Scan(&salesToday).Error; err != nil {
		return nil, err
	}
	return salesToday, nil
}

func (c *SalesConstruct) GetSalesReplica(conn string) ([]model.SalesReport, error) {
	salesToday := []model.SalesReport{}
	sql := `
	select
		count(distinct(kode_reseller)) as ma ,
		count(1) as trx,
		sum(harga_beli) as pembelian ,
		sum(harga) as penjualan,
		sum(harga) - sum(harga_beli)as laba
	from
		transaksi t
	WHERE
		status = 20
	AND t.tgl_entri BETWEEN (CONVERT(DATETIME,
	CONVERT(DATE,
	GETDATE()))) AND (
	SELECT
		CAST(GETDATE() AS DATETIME))
GROUP BY
	t.kode_reseller
	`
	if err := c.SelectConn(conn).Raw(sql).Scan(&salesToday).Error; err != nil {
		return nil, err
	}
	return salesToday, nil
}

func (c *SalesConstruct) Sales(from, to string) ([]model.SalesReport, error) {

	salesToday := []model.SalesReport{}
	sql := "SELECT COUNT(DISTINCT(t.kode_reseller)) AS ma ,count(1) AS trx, SUM(harga) AS pembelian, SUM(t.harga_beli) AS penjualan ,SUM(t.harga) - SUM(t.harga_beli) AS laba"
	sql = fmt.Sprintf("%s FROM transaksi t", sql)
	sql = fmt.Sprintf("%s WHERE t.status = 20", sql)
	sql = fmt.Sprintf("%s AND cast(t.tgl_entri AS date) BETWEEN '%s' AND '%s'", sql, from, to)
	sql = fmt.Sprintf("%s GROUP BY t.kode_reseller;", sql)
	if err := c.db.DigiEps.Raw(sql).Scan(&salesToday).Error; err != nil {
		return nil, err
	}
	return salesToday, nil
}

func (c *SalesConstruct) SalesProd(from, to string) ([]model.SalesReport, error) {

	salesToday := []model.SalesReport{}
	sql := "SELECT COUNT(DISTINCT(t.kode_reseller)) AS ma ,count(1) AS trx, SUM(harga) AS pembelian, SUM(t.harga_beli) AS penjualan ,SUM(t.harga) - SUM(t.harga_beli) AS laba"
	sql = fmt.Sprintf("%s FROM transaksi t", sql)
	sql = fmt.Sprintf("%s WHERE t.status = 20", sql)
	sql = fmt.Sprintf("%s AND cast(t.tgl_entri AS date) BETWEEN '%s' AND '%s'", sql, from, to)
	sql = fmt.Sprintf("%s GROUP BY t.kode_reseller;", sql)
	if err := c.db.DigiAmazone.Raw(sql).Scan(&salesToday).Error; err != nil {
		return nil, err
	}
	return salesToday, nil
}

func (c *SalesConstruct) SalesReplica(path, from, to string) ([]model.SalesReport, error) {
	salesToday := []model.SalesReport{}
	sql := "SELECT COUNT(DISTINCT(t.kode_reseller)) AS ma ,count(1) AS trx, SUM(harga) AS pembelian, SUM(t.harga_beli) AS penjualan ,SUM(t.harga) - SUM(t.harga_beli) AS laba"
	sql = fmt.Sprintf("%s FROM transaksi t", sql)
	sql = fmt.Sprintf("%s WHERE t.status = 20", sql)
	sql = fmt.Sprintf("%s AND cast(t.tgl_entri AS date) BETWEEN '%s' AND '%s'", sql, from, to)
	sql = fmt.Sprintf("%s GROUP BY t.kode_reseller;", sql)
	if err := c.SelectConn(path).Raw(sql).Scan(&salesToday).Error; err != nil {
		return nil, err
	}
	return salesToday, nil
}

func (c *SalesConstruct) SelectConn(path string) *gorm.DB {
	switch path {
	case utils.AMAZONE:
		return c.db.Amazone
	case utils.EPS:
		return c.db.Eps
	default:
		return c.db.Eps
	}
}
