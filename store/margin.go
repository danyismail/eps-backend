package store

import (
	"eps-backend/db"
	"eps-backend/model"
	"eps-backend/utils"
	"fmt"
)

type MarginConstruct struct {
	db db.DBConnection
}

func NewMarginStore(db db.DBConnection) *MarginConstruct {
	return &MarginConstruct{db}
}

func (c *MarginConstruct) ByReseller(path, startDt, endDt string) ([]model.MarginReseller, error) {
	var marginReseller []model.MarginReseller

	if startDt == "" {
		startDt = "2020-02-22"
	}

	if endDt == "" {
		endDt = "2020-02-29"
	}

	sql := "SELECT t.kode_reseller, r.nama, count(1) AS total_trx,  sum(t.harga_beli) AS pembelian,  sum(t.harga) AS penjualan, sum(t.harga) - sum(t.harga_beli) AS laba "
	sql = fmt.Sprintf("%s FROM transaksi t LEFT JOIN reseller r ON t.kode_reseller = r.kode ", sql)
	sql = fmt.Sprintf("%s WHERE status = 20 AND t.tgl_entri BETWEEN '%s' AND '%s'", sql, startDt, endDt)
	sql = fmt.Sprintf("%s GROUP BY t.kode_reseller, r.nama ORDER BY total_trx DESC", sql)

	conn := utils.SelectConn(path, c.db)
	if err := conn.Debug().Raw(sql).Scan(&marginReseller).Error; err != nil {
		return nil, err
	}

	return marginReseller, nil
}

func (c *MarginConstruct) BySupplier(path, startDt, endDt string) ([]model.MarginSupplier, error) {
	var marginSupplier []model.MarginSupplier

	if startDt == "" {
		startDt = "2020-02-22"
	}

	if endDt == "" {
		endDt = "2020-02-29"
	}

	sql := "SELECT m.label, count(1) AS total_trx,  sum(t.harga_beli) AS pembelian,  sum(t.harga) AS penjualan, sum(t.harga) - sum(t.harga_beli) AS laba "
	sql = fmt.Sprintf("%s FROM transaksi t LEFT JOIN modul m ON t.kode_modul = m.kode ", sql)
	sql = fmt.Sprintf("%s WHERE status = 20 AND t.tgl_entri BETWEEN '%s' AND '%s'", sql, startDt, endDt)
	sql = fmt.Sprintf("%s GROUP BY m.label ORDER BY total_trx DESC", sql)

	conn := utils.SelectConn(path, c.db)
	if err := conn.Debug().Raw(sql).Scan(&marginSupplier).Error; err != nil {
		return nil, err
	}

	return marginSupplier, nil
}

func (c *MarginConstruct) ByProvider(path, startDt, endDt string) ([]model.MarginProvider, error) {
	var marginProvider []model.MarginProvider

	if startDt == "" {
		startDt = "2020-02-22"
	}

	if endDt == "" {
		endDt = "2020-02-29"
	}

	sql := "SELECT mp.provider, count(1) AS total_trx,  sum(t.harga_beli) AS pembelian,  sum(t.harga) AS penjualan, sum(t.harga) - sum(t.harga_beli) AS laba "
	sql = fmt.Sprintf("%s FROM transaksi t LEFT JOIN produk p ON t.kode_produk = p.kode LEFT JOIN mapping_providers mp ON t.kode_produk = mp.kode", sql)
	sql = fmt.Sprintf("%s WHERE status = 20 AND t.tgl_entri BETWEEN '%s' AND '%s'", sql, startDt, endDt)
	sql = fmt.Sprintf("%s GROUP BY mp.provider ORDER BY total_trx DESC", sql)

	conn := utils.SelectConn(path, c.db)
	if err := conn.Debug().Raw(sql).Scan(&marginProvider).Error; err != nil {
		return nil, err
	}

	return marginProvider, nil
}
