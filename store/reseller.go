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

func (c *ResellerConstruct) GetLabaHourly(path string) ([][]model.CekLabaHourly, error) {
	sql := `
	SELECT 
		CAST(tgl_entri AS DATE) as tanggal,
		DATEPART(DAY, tgl_entri) as tgl,
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
		DATEPART(DAY, tgl_entri),
		DATEPART(HOUR, tgl_entri)
	ORDER BY 
		DATEPART(DAY, tgl_entri),
		DATEPART(HOUR, tgl_entri);
	`
	data := []model.CekLabaHourly{}
	conn := utils.SelectConn(path, c.db)
	if err := conn.Raw(sql).Debug().Scan(&data).Error; err != nil {
		return nil, err
	}

	aggregatedData := make(map[int]int)
	for _, item := range data {
		aggregatedData[item.Tgl] += item.Laba
	}
	var arr []int //array list of date
	for key := range aggregatedData {
		arr = append(arr, key)
	}

	var result [][]model.CekLabaHourly
	var sumTrx, sumLaba int
	for i := 0; i < len(arr); i++ {
		sumTrx, sumLaba = 0, 0
		innerArr := []model.CekLabaHourly{}
		for j := 0; j < len(data); j++ {
			if data[j].Tgl == arr[i] {
				sumTrx += data[j].Trx
				sumLaba += data[j].Laba
				data[j].Trx = sumTrx
				data[j].Laba = sumLaba
				innerArr = append(innerArr, data[j])
			}
		}
		result = append(result, innerArr)
	}
	return result, nil
}
