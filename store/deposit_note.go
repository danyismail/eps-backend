package store

import (
	"eps-backend/db"
	"eps-backend/model"
	"eps-backend/utils"
	"errors"
	"fmt"

	"gorm.io/gorm"
)

type DepositNoteConstruct struct {
	db db.DBConnection
}

func NewDepositNoteStore(db db.DBConnection) *DepositNoteConstruct {
	return &DepositNoteConstruct{
		db,
	}
}

func (c *DepositNoteConstruct) Create(notes model.DepositNote, path string) error {
	if err := c.SelectConn(path).Debug().Create(&notes).Error; err != nil {
		return err
	}
	return nil
}

func (c *DepositNoteConstruct) GetAllStatus(path, date string) ([]model.DepositNote, error) {
	notes := []model.DepositNote{}

	sql := "SELECT id,FORMAT(created_at, 'dd-MM-yyyy HH:mm:ss') created_at,FORMAT(updated_at , 'dd-MM-yyyy HH:mm:ss') updated_at, deleted_at, name, supplier, amount"
	sql = fmt.Sprintf(`%s origin_account, destination_account, image_upload, reply, status FROM deposit_notes WHERE FORMAT(created_at , 'dd-MM-yyyy HH:mm:ss') = '%s'`, sql, date)

	result := c.SelectConn(path).Debug().Raw(sql).Scan(&notes)
	if result.Error != nil {
		return nil, result.Error
	}
	return notes, nil
}

func (c *DepositNoteConstruct) GetStatusCreated(path string) ([]model.DepositNote, error) {
	notes := []model.DepositNote{}
	// query := "select * from deposit_notes dn where image_upload = '';"
	query := `
	SELECT 
		id,
		FORMAT(created_at , 'dd-MM-yyyy HH:mm:ss') created_at,
		FORMAT(updated_at , 'dd-MM-yyyy HH:mm:ss') updated_at,
		deleted_at,
		name,
		supplier,
		amount,
		origin_account,
		destination_account,
		image_upload,
		reply,
		status
	FROM deposit_notes
	WHERE status = 'pending';
	`
	result := c.SelectConn(path).Raw(query).Scan(&notes)
	if result.Error != nil {
		return nil, result.Error
	}
	return notes, nil
}

func (c *DepositNoteConstruct) GetStatusUploaded(path string) ([]model.DepositNote, error) {
	notes := []model.DepositNote{}
	// query := "select * from deposit_notes dn where image_upload <> '' and CAST(reply AS varchar(MAX)) = '';"
	query := `
	SELECT 
		id,
		FORMAT(created_at , 'dd-MM-yyyy HH:mm:ss') created_at,
		FORMAT(updated_at , 'dd-MM-yyyy HH:mm:ss') updated_at,
		deleted_at,
		name,
		supplier,
		amount,
		origin_account,
		destination_account,
		image_upload,
		reply,
		status
	FROM deposit_notes
	WHERE status = 'process';
	`
	result := c.SelectConn(path).Raw(query).Scan(&notes)
	if result.Error != nil {
		return nil, result.Error
	}
	return notes, nil
}

func (c *DepositNoteConstruct) GetStatusDone(path, startDt, endDt string) ([]model.DepositNote, error) {
	notes := []model.DepositNote{}

	whereCondition := " WHERE status = 'success'"
	sql := "SELECT id,FORMAT(created_at , 'dd-MM-yyyy HH:mm:ss') created_at,FORMAT(updated_at , 'dd-MM-yyyy HH:mm:ss') updated_at, deleted_at, name, supplier, amount,"
	sql = fmt.Sprintf("%s origin_account, destination_account, image_upload, reply, status FROM deposit_notes", sql)

	if startDt != "" && endDt != "" {
		whereCondition = fmt.Sprintf(" WHERE cast(created_at as date) BETWEEN '%s' AND '%s' AND status = 'success'", startDt, endDt)
	}
	sql += " " + whereCondition

	result := c.SelectConn(path).Debug().Raw(sql).Scan(&notes)
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

func (c *DepositNoteConstruct) Delete(id int, path string) error {
	if err := c.SelectConn(path).Debug().Exec("DELETE FROM deposit_notes WHERE id = ?", id).Error; err != nil {
		return err
	}
	return nil
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
	switch path {
	case utils.DIGI_AMAZONE:
		return c.db.DigiAmazone
	case utils.DIGI_EPS:
		return c.db.DigiEps
	default:
		return c.db.DigiEps
	}
}
