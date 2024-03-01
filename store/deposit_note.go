package store

import (
	"eps-backend/model"
	"errors"

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

func (c *DepositNoteConstruct) GetById(id int) (*model.DepositNote, error) {
	notes := new(model.DepositNote)
	if err := c.dev.Debug().Where("id = ?", id).Find(notes).Error; err != nil {
		return nil, err
	}
	if notes.ID == 0 {
		return nil, nil
	}
	return notes, nil
}

func (c *DepositNoteConstruct) Update(notes model.DepositNote) error {
	result := c.dev.Debug().Updates(&notes)
	if result.RowsAffected == 0 {
		return errors.New("no record found")
	}
	if result.Error != nil {
		return result.Error
	}
	return nil
}
