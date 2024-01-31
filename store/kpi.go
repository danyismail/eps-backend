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

func (c *KpiConstruct) FindAll(page int, view int) (*[]model.VKpi, error) {
	var kpis []model.VKpi
	var total int64
	if page == -1 && view == -1 {
		if err := c.db.Find(&kpis).Count(&total).Error; err != nil {
			if err == gorm.ErrRecordNotFound {
				return nil, nil
			}
			return nil, err
		}
	} else {
		if err := c.db.Limit(view).Offset(view * (page - 1)).Find(&kpis).Error; err != nil {
			if err == gorm.ErrRecordNotFound {
				return nil, nil
			}
			return nil, err
		}
	}
	fmt.Println(total)
	return &kpis, nil
}

func (c *KpiConstruct) Test() (*[]model.VKpi, error) {
	modelKpi := []model.VKpi{}
	result := c.db.Raw("SELECT TOP 10 * FROM v_kpi").Scan(&modelKpi)
	if result.Error != nil {
		return nil, result.Error
	}
	return &modelKpi, nil
}
