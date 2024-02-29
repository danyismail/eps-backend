package store

import (
	"eps-backend/model"

	"gorm.io/gorm"
)

type DepositNoteConstruct struct {
	dev  *gorm.DB
	prod *gorm.DB
}

func NewDepositNoteStore(dev *gorm.DB, prod *gorm.DB) *DepositNoteConstruct {
	return &DepositNoteConstruct{
		dev,
		prod,
	}
}

func (c *DepositNoteConstruct) Create(notes model.DepositNote) error {
	if err := c.dev.Debug().Create(&notes).Error; err != nil {
		return err
	}
	return nil
}

func (c *DepositNoteConstruct) GetById(id int) (model.DepositNote, error) {
	notes := model.DepositNote{}
	if err := c.dev.Debug().Where("id = ?", id).Find(&notes).Error; err != nil {
		return notes, err
	}
	return notes, nil
}

func (c *DepositNoteConstruct) Update(notes model.DepositNote) error {
	if err := c.dev.Debug().Updates(&notes).Error; err != nil {
		return err
	}
	return nil
}
