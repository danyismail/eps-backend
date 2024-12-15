package store

import (
	"eps-backend/db"
	"eps-backend/model"
)

type ValueListConstruct struct {
	db db.DBConnection
}

func NewValueListStore(db db.DBConnection) *ValueListConstruct {
	return &ValueListConstruct{
		db: db,
	}
}

func (c *ValueListConstruct) Create(vl *model.ValueList) error {
	if err := c.db.DigiEps.Debug().Table("value_list").Create(vl).Error; err != nil {
		return err
	}
	return nil
}

func (c *ValueListConstruct) GetAll(page, view int, shortCode string) (vl []model.ValueList, total int, err error) {
	if page <= 0 {
		page = 1
	}
	if view <= 0 {
		view = 10
	}

	offset := (page - 1) * view
	query := c.db.DigiEps.Debug().Table("value_list")

	if shortCode != "" {
		query = query.Where("short_code = ?", shortCode)
	}

	var count int64
	if err := query.Count(&count).Error; err != nil {
		return nil, 0, err
	}

	if err := query.Limit(view).Offset(offset).Find(&vl).Error; err != nil {
		return nil, 0, err
	}

	return vl, int(count), nil
}

func (c *ValueListConstruct) GetOne(id int) (*model.ValueList, error) {
	var vl model.ValueList

	//check if row exist
	err := c.db.DigiEps.Debug().Table("value_list").Where("id = ?", id).First(&model.ValueList{}).Error
	if err != nil {
		return nil, err
	}

	if err := c.db.DigiEps.Debug().Table("value_list").Where("id = ?", id).First(&vl).Error; err != nil {
		return nil, err
	}
	return &vl, nil
}

func (c *ValueListConstruct) Update(id int, vl *model.ValueList) error {
	if err := c.db.DigiEps.Debug().Table("value_list").Where("id = ?", id).Updates(vl).Error; err != nil {
		return err
	}
	return nil
}

func (c *ValueListConstruct) Delete(id int) error {
	//check if row exist
	err := c.db.DigiEps.Debug().Table("value_list").Where("id = ?", id).First(&model.ValueList{}).Error
	if err != nil {
		return err
	}

	if err := c.db.DigiEps.Debug().Table("value_list").Where("id = ?", id).Delete(&model.ValueList{}).Error; err != nil {
		return err
	}
	return nil
}
