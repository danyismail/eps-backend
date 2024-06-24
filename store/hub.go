package store

import (
	"eps-backend/db"
	"eps-backend/model"
	"fmt"
	"log"
	"sync"
)

type HubConstruct struct {
	db db.DBConnection
}

func NewHubStore(db db.DBConnection) *HubConstruct {
	return &HubConstruct{db}
}

func (c *HubConstruct) GetBrandRevenue(startDt, endDt string) ([]model.BrandRevenue, error) {
	var wg sync.WaitGroup
	//c.db.DigiAmazone, c.db.DigiEps, c.db.Amazone, c.db.Eps, c.db.Backup, c.db.Otodev
	listCnx := []model.Dbs{
		{
			Name: "Digipos Amazone",
			Cnx:  c.db.DigiAmazone,
		},
		{
			Name: "Digipos EPS",
			Cnx:  c.db.DigiEps,
		},
		{
			Name: "Replica Amazone",
			Cnx:  c.db.Amazone,
		},
		{
			Name: "Replica EPS",
			Cnx:  c.db.Eps,
		},
		{
			Name: "Amazone",
			Cnx:  c.db.Backup,
		},
		{
			Name: "Otodev",
			Cnx:  c.db.Otodev,
		},
	}

	results := make(chan *model.BrandRevenue, len(listCnx))
	for _, v := range listCnx {
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
	totalMa := 0
	var totalPenjualan float64
	var totalPembelian float64
	var totalTekor float64
	var totalBakar float64
	var totalLaba float64
	var totalKomisi float64
	var totalPpn float64
	var totalPph float64
	for _, v := range allResults {
		totalMa += v.Member
		totalPenjualan += v.Penjualan
		totalPembelian += v.Pembelian
		totalTekor += v.Tekor
		totalBakar += v.Bakar
		totalLaba += v.Laba
		totalKomisi += v.Komisi
		totalPpn += v.Ppn11
		totalPph += v.Pph22
	}
	allResults = append(allResults, model.BrandRevenue{
		Server:    "Total",
		Member:    totalMa,
		Penjualan: totalPenjualan,
		Pembelian: totalPembelian,
		Tekor:     totalTekor,
		Bakar:     totalBakar,
		Laba:      totalLaba,
		Komisi:    totalKomisi,
		Ppn11:     totalPpn,
		Pph22:     totalPph,
	})
	return allResults, nil
}

func getBrandRevenue(db model.Dbs, startDt, endDt string, result chan<- *model.BrandRevenue, wg *sync.WaitGroup) {
	defer wg.Done()

	// start a new transaction.
	tx := db.Cnx.Begin()
	if tx.Error != nil {
		result <- nil
		log.Printf("failed to begin transaction: %v", tx.Error)
		return
	}

	// ensure the transaction is rolled back if not committed.
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
			log.Printf("transaction rolled back due to panic: %v", r)
		} else {
			tx.Rollback() // this will be called if we reach the end without committing.
			log.Println("transaction rolled back due to error or completion.")
		}
	}()

	switch startDt {
	case "":
		startDt = " tgl_entri >= CAST(GETDATE() AS DATE) "
		endDt = " tgl_entri < DATEADD(DAY, 1, CAST(GETDATE() AS DATE)) "
	case "yesterday":
		startDt = " tgl_entri >= DATEADD(DAY, -1, CAST(GETDATE() AS DATE)) "
		endDt = " tgl_entri < CAST(GETDATE() AS DATE) "
	case "beforeYesterday":
		startDt = " tgl_entri >= DATEADD(DAY, -2, CAST(GETDATE() AS DATE)) "
		endDt = " tgl_entri < DATEADD(DAY, -1, CAST(GETDATE() AS DATE)) "
	}

	sql := "SELECT FORMAT(sum(t.harga),'0.######') penjualan, FORMAT(sum(t.harga_beli),'0.######') pembelian, "
	sql = fmt.Sprintf("%s CASE WHEN sum(t.harga) - sum(t.harga_beli) > 0 THEN 0 WHEN sum(t.harga) - sum(t.harga_beli) < 0 THEN FORMAT(sum(t.harga) - sum(t.harga_beli),'0.######')", sql)
	sql = fmt.Sprintf("%s END AS tekor, FORMAT(sum(t.harga - p.harga_jual),'0.######') bakar, FORMAT(sum(t.harga) - sum(t.harga_beli),'0.######') laba, FORMAT(sum(komisi) ,'0.######') komisi, 0 ppn11, 0 pph22", sql)
	sql = fmt.Sprintf("%s FROM transaksi t LEFT JOIN produk p on t.kode_produk = p.kode LEFT JOIN reseller r ON t.kode_reseller = r.kode ", sql)
	sql = fmt.Sprintf("%s WHERE status = 20 AND %s AND %s", sql, startDt, endDt)

	//1. get brand revenue
	var brandsRevenue *model.BrandRevenue
	if err := db.Cnx.Debug().Raw(sql).Scan(&brandsRevenue).Error; err != nil {
		result <- nil
		return
	}

	//2. get member active
	member_active := 0
	countActiveMember := fmt.Sprintf("SELECT COUNT(DISTINCT kode_reseller) AS member_active FROM transaksi WHERE status = 20 AND %s AND %s", startDt, endDt)
	if err := db.Cnx.Debug().Raw(countActiveMember).Scan(&member_active).Error; err != nil {
		log.Printf("failed to query data: %v", err)
		result <- nil
		return
	}

	//3. get pph 22
	getPPH12 := fmt.Sprintf(`
	SELECT
		SUM(pph) AS total_pph
	FROM
	(
		SELECT
			COUNT(1) AS trx,
			SUM(t.harga_beli) AS pembelian,
			SUM(t.harga) AS penjualan,
			SUM(t.harga) - SUM(t.harga_beli) AS laba,
			SUM(t.harga) * 0.005 AS pph
		FROM
			transaksi t
		JOIN produk p on
			t.kode_produk = p.kode
		WHERE
			status = 20
			AND (
				(t.kode_produk NOT LIKE '%%NF%%' AND kode_reseller IN ('EPS0695', 'EPS0634', 'EPS0840', 'EPS0935'))
				OR (t.kode_produk LIKE '%%FZNF%%' AND kode_reseller IN ('EPS6995', 'EPS6890', 'EPS6957', 'EPS6973', 'EPS0921'))
				OR (t.kode_produk NOT LIKE '%%NF%%' AND kode_reseller IN ('EPS6995', 'EPS6890', 'EPS6957', 'EPS6973', 'EPS0921'))
				OR (t.kode_produk LIKE '%%SBNONF%%' AND kode_reseller = 'EPS0712')
				OR (t.kode_produk NOT LIKE '%%NF%%' AND kode_reseller = 'EPS0712')
			)
			AND p.kode_operator NOT IN ('PLN', 'PLN01', 'PLN02', 'ECOMM', 'GAMES')
			AND %s AND %s
	) pph22;
	`, startDt, endDt)

	var pphCount model.PPH
	if err := db.Cnx.Raw(getPPH12).Debug().Scan(&pphCount).Error; err != nil {
		result <- nil
		return
	}

	// if everything is successful, commit the transaction.
	if err := tx.Commit().Error; err != nil {
		log.Printf("failed to commit transaction: %v", err)
		return
	}
	log.Println("transaction committed successfully.")

	brandsRevenue.Server = db.Name
	brandsRevenue.Member = member_active
	brandsRevenue.Pph22 = pphCount.TotalPph
	result <- brandsRevenue

}
