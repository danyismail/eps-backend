package store

import (
	"eps-backend/db"
	"eps-backend/model"
	"eps-backend/utils"
	"errors"
	"fmt"
)

type DepositNoteConstruct struct {
	db db.DBConnection
}

func NewDepositNoteStore(db db.DBConnection) *DepositNoteConstruct {
	return &DepositNoteConstruct{
		db,
	}
}

func (c *DepositNoteConstruct) GetAllStatus(path, date string) ([]model.DepositNote, error) {
	notes := []model.DepositNote{}

	sql := "SELECT id,FORMAT(created_at, 'dd-MM-yyyy HH:mm:ss') created_at,FORMAT(updated_at , 'dd-MM-yyyy HH:mm:ss') updated_at, deleted_at, name, supplier, amount"
	sql = fmt.Sprintf(`%s origin_account, destination_account, image_upload, reply, status FROM deposit_notes WHERE FORMAT(created_at , 'dd-MM-yyyy HH:mm:ss') = '%s'`, sql, date)

	result := utils.SelectConn(path, c.db).Raw(sql).Scan(&notes)
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
	result := utils.SelectConn(path, c.db).Raw(query).Scan(&notes)
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
	result := utils.SelectConn(path, c.db).Raw(query).Scan(&notes)
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

	result := utils.SelectConn(path, c.db).Raw(sql).Scan(&notes)
	if result.Error != nil {
		return nil, result.Error
	}
	return notes, nil
}

func (c *DepositNoteConstruct) Create(path string, notes model.DepositNote) error {
	if err := utils.SelectConn(path, c.db).Create(&notes).Error; err != nil {
		return err
	}
	return nil
}

func (c *DepositNoteConstruct) GetById(path string, id int) (*model.DepositNote, error) {
	notes := new(model.DepositNote)
	if err := utils.SelectConn(path, c.db).Where("id = ?", id).Find(notes).Error; err != nil {
		return nil, err
	}
	if notes.ID == 0 {
		return nil, nil
	}
	return notes, nil
}

func (c *DepositNoteConstruct) Update(path string, notes model.DepositNote) error {
	result := utils.SelectConn(path, c.db).Updates(&notes)
	if result.RowsAffected == 0 {
		return errors.New("no record found")
	}
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (c *DepositNoteConstruct) Delete(path string, id int) error {
	if err := utils.SelectConn(path, c.db).Exec("DELETE FROM deposit_notes WHERE id = ?", id).Error; err != nil {
		return err
	}
	return nil
}
