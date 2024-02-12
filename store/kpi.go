package store

import (
	"eps-backend/model"
	"fmt"
	"strings"
	"time"

	"gorm.io/gorm"
)

type KpiConstruct struct {
	dev  *gorm.DB
	prod *gorm.DB
}

func NewKpiStore(dev *gorm.DB, prod *gorm.DB) *KpiConstruct {
	return &KpiConstruct{
		dev,
		prod,
	}
}

func (c *KpiConstruct) FindAll(startDt string, endDt string, pageNumber int, pageSize int, mdn string, status int, shift string) (*[]model.VKpis, int64, int64, error) {
	var kpis []model.VKpis
	var totalRow int64

	if startDt == "" {
		startDt = time.Now().AddDate(0, 0, -1).Format("2006-01-02")
	}

	if endDt == "" {
		endDt = time.Now().Format("2006-01-02")
	}

	if pageNumber == 0 {
		pageNumber = 1
	}
	if pageSize == 0 {
		pageSize = 200
	}

	offset := (pageNumber - 1) * pageSize
	fetch := pageSize

	sql := fmt.Sprintf("SELECT * FROM v_kpis WHERE cast(tgl_entri as date) BETWEEN '%s' AND '%s'", startDt, endDt)
	var whereQuery string
	if mdn != "" {
		whereQuery += fmt.Sprintf(" AND tujuan =  '%s'", mdn)
	}

	if status != 0 {
		whereQuery += fmt.Sprintf(" AND status =  %d", status)
	}

	if shift != "" {
		whereQuery += fmt.Sprintf(" AND shift = '%s'", shift)
	}

	if whereQuery != "" {
		sql = fmt.Sprintf("%s %s", sql, whereQuery)
	}

	//count here
	var totalData int64
	countQuery := strings.Replace(sql, "*", "COUNT(1)", -1)
	if err := c.dev.Debug().Raw(countQuery).Scan(&totalData).Error; err != nil {
		return nil, totalData, totalRow, err
	}

	if pageNumber > 0 && pageSize > 0 {
		sql = fmt.Sprintf("%s ORDER BY (tgl_entri) DESC OFFSET %d ROWS FETCH NEXT %d ROW ONLY", sql, offset, fetch)
	}

	if err := c.dev.Debug().Raw(sql).Scan(&kpis).Error; err != nil {
		return nil, totalData, totalRow, err
	}
	totalRow = int64(len(kpis))
	if len(kpis) <= 0 {
		return nil, totalData, totalRow, nil
	}
	return &kpis, totalData, totalRow, nil
}

func (c *KpiConstruct) FindAllProd(startDt string, endDt string, pageNumber int, pageSize int, mdn string, status int, shift string) (*[]model.VKpis, int64, int64, error) {
	var kpis []model.VKpis
	var totalRow int64

	if startDt == "" {
		startDt = time.Now().AddDate(0, 0, -1).Format("2006-01-02")
	}

	if endDt == "" {
		endDt = time.Now().Format("2006-01-02")
	}

	if pageNumber == 0 {
		pageNumber = 1
	}
	if pageSize == 0 {
		pageSize = 200
	}

	offset := (pageNumber - 1) * pageSize
	fetch := pageSize

	sql := fmt.Sprintf("SELECT * FROM v_kpis WHERE cast(tgl_entri as date) BETWEEN '%s' AND '%s'", startDt, endDt)
	var whereQuery string
	if mdn != "" {
		whereQuery += fmt.Sprintf(" AND tujuan =  '%s'", mdn)
	}

	if status != 0 {
		whereQuery += fmt.Sprintf(" AND status =  %d", status)
	}

	if shift != "" {
		whereQuery += fmt.Sprintf(" AND shift = '%s'", shift)
	}

	if whereQuery != "" {
		sql = fmt.Sprintf("%s %s", sql, whereQuery)
	}

	//count here
	var totalData int64
	countQuery := strings.Replace(sql, "*", "COUNT(1)", -1)
	if err := c.prod.Debug().Raw(countQuery).Scan(&totalData).Error; err != nil {
		return nil, totalData, totalRow, err
	}

	if pageNumber > 0 && pageSize > 0 {
		sql = fmt.Sprintf("%s ORDER BY (tgl_entri) DESC OFFSET %d ROWS FETCH NEXT %d ROW ONLY", sql, offset, fetch)
	}

	if err := c.prod.Debug().Raw(sql).Scan(&kpis).Error; err != nil {
		return nil, totalData, totalRow, err
	}
	totalRow = int64(len(kpis))
	if len(kpis) <= 0 {
		return nil, totalData, totalRow, nil
	}
	return &kpis, totalData, totalRow, nil
}

func (c *KpiConstruct) Test() (*[]model.VKpis, error) {
	modelKpi := []model.VKpis{}
	result := c.dev.Raw("SELECT TOP 10 * FROM v_kpi").Scan(&modelKpi)
	if result.Error != nil {
		return nil, result.Error
	}
	return &modelKpi, nil
}
