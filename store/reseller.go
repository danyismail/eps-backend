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

	sql := "SELECT t.kode_produk, COUNT(1) trx, t.kode_reseller, CONVERT(INT, SUM(harga-harga_beli)) laba"
	sql = fmt.Sprintf("%s FROM transaksi t JOIN reseller r ON t.kode_reseller = r.kode WHERE t.kode_reseller = '%s' AND CAST(tgl_entri AS DATE) BETWEEN '%s' AND '%s'", sql, id, startDt, endDt)
	sql = fmt.Sprintf("%s AND status = %d", sql, 20)
	sql = fmt.Sprintf("%s GROUP BY t.kode_reseller, t.kode_produk ORDER BY t.kode_produk", sql)

	conn := utils.SelectConn(path, c.db)
	if err := conn.Raw(sql).Scan(&labaReseller).Error; err != nil {
		return nil, err
	}

	return labaReseller, nil
}

func (c *ResellerConstruct) GetSum(path, startDt, endDt, id string) (*model.SumLabaReseller, error) {
	var labaReseller model.SumLabaReseller

	if startDt == "" {
		startDt = time.Now().AddDate(0, 0, -1).Format("2006-01-02")
	}

	if endDt == "" {
		endDt = time.Now().Format("2006-01-02")
	}

	sql := "SELECT r.nama,COUNT(1) trx,CONVERT(INT, sum(harga - harga_beli)) laba"
	sql = fmt.Sprintf("%s FROM transaksi t JOIN reseller r ON t.kode_reseller = r.kode WHERE t.kode_reseller = '%s' AND CAST(tgl_entri AS DATE) BETWEEN '%s' AND '%s'", sql, id, startDt, endDt)
	sql = fmt.Sprintf("%s AND status = %d", sql, 20)
	sql = fmt.Sprintf("%s GROUP BY r.nama", sql)

	conn := utils.SelectConn(path, c.db)
	if err := conn.Raw(sql).Scan(&labaReseller).Error; err != nil {
		return nil, err
	}

	return &labaReseller, nil
}

func (c *ResellerConstruct) GetList(path, arg string) ([]model.Reseller, error) {
	list := []model.Reseller{}
	likeStart := "'%"
	likeEnd := "%'"
	sql := fmt.Sprintf("select kode,nama from reseller r where r.kode like %s%s%s or r.nama like %s%s%s ;", likeStart, arg, likeEnd, likeStart, arg, likeEnd)
	conn := utils.SelectConn(path, c.db)
	if err := conn.Raw(sql).Debug().Scan(&list).Error; err != nil {
		return nil, err
	}
	return list, nil
}

func (c *ResellerConstruct) GetLabaHourly(path string) (*model.ResponseLabaPerJam, error) {
	cl := []model.CekLabaHourly{}
	sql := `
	SELECT 
		CAST(tgl_entri AS DATE) as tanggal,
		CASE 
			WHEN DATEPART(HOUR, tgl_entri) + 1 = 24 THEN 0
			ELSE DATEPART(HOUR, tgl_entri) + 1
		END AS jam,
		COUNT(1)  AS trx,
		CAST(SUM(harga - harga_beli) AS INT)  AS laba
	FROM 
		transaksi t
	WHERE 
		tgl_entri >= CONVERT(datetime, CONVERT(date, DATEADD(day, -2, GETDATE()))) AND status = 20
	GROUP BY 
		CAST(tgl_entri AS DATE),
		DATEPART(HOUR, tgl_entri)
	ORDER BY 
		CAST(tgl_entri AS DATE),
		DATEPART(HOUR, tgl_entri);
	`
	conn := utils.SelectConn(path, c.db)
	if err := conn.Raw(sql).Debug().Scan(&cl).Error; err != nil {
		return nil, err
	}

	// Map to store aggregated totals
	aggregatedData := make(map[string]int)
	// // Aggregate data based on "day" key
	mapData := []model.CekLabaHourly{}
	var sumTrx, sumLaba int
	for _, item := range cl {
		aggregatedData[item.Tanggal.Local().UTC().Format(utils.DateOnly)] += item.Trx
		sumTrx += item.Trx
		sumLaba += item.Laba
		item.Trx = sumTrx
		item.Laba = sumLaba
		mapData = append(mapData, item)
	}
	result := model.ResponseLabaPerJam{
		HourlyData: mapData,
		Aggregate:  aggregatedData,
	}
	// Convert aggregated data to slice of DayTotal structs
	// var result []model.CekLabaHourly
	// for tanggal, total := range aggregatedData {
	// 	result = append(result, model.CekLabaHourly{Tanggal: tanggal, Trx: total})
	// }
	// return result, nil
	return &result, nil
}
