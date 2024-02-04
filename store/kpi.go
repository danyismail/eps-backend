package store

import (
	"eps-backend/model"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type KpiConstruct struct {
	db *gorm.DB
}

func NewKpiStore(db *gorm.DB) *KpiConstruct {
	return &KpiConstruct{
		db: db,
	}
}

func (c *KpiConstruct) FindAll(startDt string, endDt string, pageNumber int, pageSize int, mdn string, status int, shift string) (*[]model.VKpis, int64, error) {
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

	if pageNumber > 0 && pageSize > 0 {
		sql = fmt.Sprintf("%s ORDER BY (tgl_entri) DESC OFFSET %d ROWS FETCH NEXT %d ROW ONLY", sql, offset, fetch)
	}

	if err := c.db.Debug().Raw(sql).Scan(&kpis).Count(&totalRow).Error; err != nil {
		return nil, totalRow, err
	}
	return &kpis, totalRow, nil
	/*
		var kpis []model.VKpis
		var total int64
		if page == 0 && view == 0 {
			if err := c.db.Debug().Limit(1000).Find(&kpis).Order("tgl_entri DESC").Count(&total).Error; err != nil {
				if err == gorm.ErrRecordNotFound {
					return nil, total, nil
				}
				return nil, total, err
			}
		} else {
			if err := c.db.Debug().Limit(view).Offset(view * (page - 1)).Find(&kpis).Order("tgl_entri DESC").Count(&total).Error; err != nil {
				if err == gorm.ErrRecordNotFound {
					return nil, total, nil
				}
				return nil, total, err
			}
		}
		fmt.Println(total)
		if page-1 != 0 {
			var ttl int64
			if e := c.db.Raw("SELECT count(*) FROM v_kpis as ttl").Scan(&ttl).Error; e != nil {
				return nil, total, e
			}
			total = ttl
		}
		return &kpis, total, nil
	*/
}

func (c *KpiConstruct) Test() (*[]model.VKpis, error) {
	modelKpi := []model.VKpis{}
	result := c.db.Raw("SELECT TOP 10 * FROM v_kpi").Scan(&modelKpi)
	if result.Error != nil {
		return nil, result.Error
	}
	return &modelKpi, nil
}
