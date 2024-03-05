package store

import (
	"eps-backend/model"
	"eps-backend/utils"
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

func (c *DepositNoteConstruct) Create(notes model.DepositNote, path string) error {
	if err := c.SelectConn(path).Debug().Create(&notes).Error; err != nil {
		return err
	}
	return nil
}

func (c *DepositNoteConstruct) GetStatusCreated(path string) ([]model.DepositNote, error) {
	notes := []model.DepositNote{}
	query := "select * from deposit_notes dn where image_upload = '';"

	result := c.SelectConn(path).Raw(query).Scan(&notes)
	if result.Error != nil {
		return nil, result.Error
	}
	return notes, nil
}

func (c *DepositNoteConstruct) GetStatusUploaded(path string) ([]model.DepositNote, error) {
	notes := []model.DepositNote{}
	query := "select * from deposit_notes dn where image_upload <> '' and CAST(reply AS varchar(MAX)) = '';"

	result := c.SelectConn(path).Raw(query).Scan(&notes)
	if result.Error != nil {
		return nil, result.Error
	}
	return notes, nil
}

func (c *DepositNoteConstruct) GetStatusDone(path string) ([]model.DepositNote, error) {
	notes := []model.DepositNote{}
	query := "select * from deposit_notes dn where image_upload <> '' and CAST(reply AS varchar(MAX)) <> '';"
	result := c.SelectConn(path).Raw(query).Scan(&notes)
	if result.Error != nil {
		return nil, result.Error
	}
	return notes, nil
}

func (c *DepositNoteConstruct) GetById(id int, path string) (*model.DepositNote, error) {
	notes := new(model.DepositNote)
	if err := c.SelectConn(path).Debug().Where("id = ?", id).Find(notes).Error; err != nil {
		return nil, err
	}
	if notes.ID == 0 {
		return nil, nil
	}
	return notes, nil
}

func (c *DepositNoteConstruct) Update(notes model.DepositNote, path string) error {
	result := c.SelectConn(path).Debug().Updates(&notes)
	if result.RowsAffected == 0 {
		return errors.New("no record found")
	}
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (c *DepositNoteConstruct) SelectConn(path string) *gorm.DB {
	if path == utils.Amazon {
		return c.prod
	}
	return c.dev
}
