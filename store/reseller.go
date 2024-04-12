package store

import (
	"eps-backend/db"
	"eps-backend/model"
	"eps-backend/utils"
	"fmt"
	"time"
)

type ResellerConstruct struct {
	db db.DBConnection
}

func NewResellerStore(db db.DBConnection) *ResellerConstruct {
	return &ResellerConstruct{db}
}

func (c *ResellerConstruct) GetLaba(path, startDt, endDt, id string) ([]model.LabaReseller, error) {
	var labaReseller []model.LabaReseller

	if startDt == "" {
		startDt = time.Now().AddDate(0, 0, -1).Format("2006-01-02")
	}

	if endDt == "" {
		endDt = time.Now().Format("2006-01-02")
	}

	sql := "SELECT t.kode_produk, COUNT(1) trx, t.kode_reseller, nama, CONVERT(INT, SUM(harga-harga_beli)) laba"
	sql = fmt.Sprintf("%s FROM transaksi t JOIN reseller r ON t.kode_reseller = r.kode WHERE t.kode_reseller = '%s' AND CAST(tgl_entri AS DATE) BETWEEN '%s' AND '%s'", sql, id, startDt, endDt)
	sql = fmt.Sprintf("%s AND status = %d", sql, 20)
	sql = fmt.Sprintf("%s GROUP BY t.kode_reseller, nama, t.kode_produk ORDER BY t.kode_produk", sql)

	conn := utils.SelectConn(path, c.db)
	if err := conn.Debug().Raw(sql).Scan(&labaReseller).Error; err != nil {
		return nil, err
	}

	return labaReseller, nil
}
