package store

import (
	"eps-backend/db"
	"eps-backend/model"
	"eps-backend/utils"
	"fmt"
	"strings"
	"time"
)

type KpiConstruct struct {
	db db.DBConnection
}

func NewKpiStore(db db.DBConnection) *KpiConstruct {
	return &KpiConstruct{
		db: db,
	}
}

func (c *KpiConstruct) GetAll(path string, startDt string, endDt string, pageNumber int, pageSize int, mdn string, status int, shift string) (*[]model.VKpis, model.AttributeKPI, error) {
	var kpis []model.VKpis
	var attr model.AttributeKPI

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
		pageSize = 10
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

	//count all data
	countQuery := strings.Replace(sql, "*", "COUNT(1)", -1)
	if err := c.db.DigiAmazone.Raw(countQuery).Scan(&attr.Total).Error; err != nil {
		return nil, attr, err
	}
	fmt.Println("all kpi : ", attr.Total)

	querySuccess := fmt.Sprintf("%s %s", sql, " AND kpi <= 180")
	countSuccess := strings.Replace(querySuccess, "*", "COUNT(1)", -1)
	if err := utils.SelectConn(path, c.db).Raw(countSuccess).Scan(&attr.Success).Error; err != nil {
		return nil, attr, err
	}
	fmt.Println("success kpi : ", attr.Success)

	queryFailed := fmt.Sprintf("%s %s", sql, " AND kpi > 180")
	countFailed := strings.Replace(queryFailed, "*", "COUNT(1)", -1)
	if err := utils.SelectConn(path, c.db).Raw(countFailed).Scan(&attr.Failed).Error; err != nil {
		return nil, attr, err
	}
	fmt.Println("failed kpi : ", attr.Failed)

	//get data with limit
	if pageNumber > 0 && pageSize > 0 {
		sql = fmt.Sprintf("%s ORDER BY (tgl_entri) DESC OFFSET %d ROWS FETCH NEXT %d ROW ONLY", sql, offset, fetch)
	}

	conn := utils.SelectConn(path, c.db)
	if err := conn.Raw(sql).Scan(&kpis).Error; err != nil {
		return nil, attr, err
	}

	attr.View = int64(len(kpis))
	if attr.View <= 0 {
		return nil, attr, nil
	}
	return &kpis, attr, nil
}
