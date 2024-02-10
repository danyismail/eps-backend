package store

import (
	"eps-backend/model"
	"fmt"

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

func (c *SalesConstruct) Sales(from, to string) ([]model.SalesReport, error) {

	salesToday := []model.SalesReport{}
	sql := "SELECT COUNT(DISTINCT(t.kode_reseller)) AS ma ,count(1) AS trx, SUM(harga) AS pembelian, SUM(t.harga_beli) AS penjualan ,SUM(t.harga) - SUM(t.harga_beli) AS laba"
	sql = fmt.Sprintf("%s FROM transaksi t", sql)
	sql = fmt.Sprintf("%s WHERE t.status = 20", sql)
	sql = fmt.Sprintf("%s AND cast(t.tgl_entri AS date) BETWEEN '%s' AND '%s'", sql, from, to)
	sql = fmt.Sprintf("%s GROUP BY t.kode_reseller;", sql)
	if err := c.db.Debug().Raw(sql).Scan(&salesToday).Error; err != nil {
		return nil, err
	}
	return salesToday, nil
}
