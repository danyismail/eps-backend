package store

import (
	"eps-backend/model"
	"fmt"

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

func (c *KpiConstruct) FindAll(page int, view int, mdn string) (*[]model.VKpis, int64, error) {
	var kpis []model.VKpis
	var total int64
	if page == 0 && view == 0 {
		if err := c.db.Debug().Limit(10).Find(&kpis).Order("tgl_entri DESC").Count(&total).Error; err != nil {
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
}

func (c *KpiConstruct) Test() (*[]model.VKpis, error) {
	modelKpi := []model.VKpis{}
	result := c.db.Raw("SELECT TOP 10 * FROM v_kpi").Scan(&modelKpi)
	if result.Error != nil {
		return nil, result.Error
	}
	return &modelKpi, nil
}
