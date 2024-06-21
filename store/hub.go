package store

import (
	"eps-backend/db"
	"eps-backend/model"
	"fmt"
	"sync"

	"gorm.io/gorm"
)

type HubConstruct struct {
	db db.DBConnection
}

func NewHubStore(db db.DBConnection) *HubConstruct {
	return &HubConstruct{db}
}

func (c *HubConstruct) GetBrandRevenue(startDt, endDt string) ([]model.BrandRevenue, error) {
	var wg sync.WaitGroup
	dbs := []*gorm.DB{c.db.DigiAmazone, c.db.DigiEps, c.db.Amazone, c.db.Eps, c.db.Backup, c.db.Otodev}

	results := make(chan *model.BrandRevenue, len(dbs))
	for _, v := range dbs {
		wg.Add(1)
		go getBrandRevenue(v, startDt, endDt, results, &wg)
	}

	wg.Wait()
	close(results)

	allResults := make([]model.BrandRevenue, 0)
	for result := range results {
		if result != nil {
			allResults = append(allResults, *result)
		}
	}
	return allResults, nil
}

func getBrandRevenue(db *gorm.DB, startDt, endDt string, result chan<- *model.BrandRevenue, wg *sync.WaitGroup) {
	defer wg.Done()

	switch startDt {
	case "":
		startDt = " tgl_entri >= CAST(GETDATE() AS DATE) AND tgl_entri < DATEADD(DAY, 1, CAST(GETDATE() AS DATE)) "
	case "yesterday":
		startDt = " tgl_entri >= DATEADD(DAY, -1, CAST(GETDATE() AS DATE)) AND tgl_entri < CAST(GETDATE() AS DATE) "
	case "beforeYesterday":
		startDt = " tgl_entri >= DATEADD(DAY, -2, CAST(GETDATE() AS DATE)) AND tgl_entri < DATEADD(DAY, -1, CAST(GETDATE() AS DATE)) "
	}

	sql := "SELECT t.kode_reseller, r.nama, count(1) AS total_trx,  sum(t.harga_beli) AS pembelian,  sum(t.harga) AS penjualan, sum(t.harga) - sum(t.harga_beli) AS laba "
	sql = fmt.Sprintf("%s FROM transaksi t LEFT JOIN reseller r ON t.kode_reseller = r.kode ", sql)
	sql = fmt.Sprintf("%s WHERE status = 20 AND t.tgl_entri BETWEEN '%s' AND '%s'", sql, startDt, endDt)
	sql = fmt.Sprintf("%s GROUP BY t.kode_reseller, r.nama ORDER BY total_trx DESC", sql)

	var brandsRevenue *model.BrandRevenue
	if err := db.Debug().Raw(sql).Scan(&brandsRevenue).Error; err != nil {
		result <- nil
		return
	}
	result <- brandsRevenue
}
