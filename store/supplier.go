package store

import (
	"eps-backend/db"
	"eps-backend/model"
	"eps-backend/utils"
	"errors"
	"fmt"
)

type SupplierConstruct struct {
	db db.DBConnection
}

func NewSupplierStore(db db.DBConnection) *SupplierConstruct {
	return &SupplierConstruct{db}
}

func (c *SupplierConstruct) GetSuppliers(path string) ([]model.Supplier, error) {
	suppliers := []model.Supplier{}
	sql := "SELECT * FROM suppliers"
	if err := utils.SelectConn(path, c.db).Raw(sql).Scan(&suppliers).Error; err != nil {
		return nil, err
	}
	return suppliers, nil
}

func (c *SupplierConstruct) GetActiveSuppliers(path string) ([]model.Supplier, error) {
	suppliers := []model.Supplier{}
	sql := "SELECT * FROM suppliers WHERE status = 'active'"
	if err := utils.SelectConn(path, c.db).Raw(sql).Scan(&suppliers).Error; err != nil {
		return nil, err
	}
	return suppliers, nil
}

func (c *SupplierConstruct) CreateSupplier(path string, spl model.Supplier) error {
	if err := utils.SelectConn(path, c.db).Create(&spl).Error; err != nil {
		return err
	}
	return nil
}

func (c *SupplierConstruct) GetSupplierById(path string, id int) (model.Supplier, error) {
	supplier := model.Supplier{}
	if err := utils.SelectConn(path, c.db).Table("suppliers").Where("id = ?", id).Scan(&supplier).Error; err != nil {
		return supplier, err
	}
	if supplier.ID == 0 {
		return supplier, errors.New("data not found")
	}
	return supplier, nil
}

func (c *SupplierConstruct) UpdateSupplier(path string, s model.Supplier) error {
	if err := utils.SelectConn(path, c.db).Updates(&s).Error; err != nil {
		return err
	}
	return nil
}

func (c *SupplierConstruct) DeleteSupplier(path string, id int) error {
	sql := fmt.Sprintf("DELETE FROM suppliers WHERE id = %d", id)
	if err := utils.SelectConn(path, c.db).Exec(sql).Error; err != nil {
		return err
	}
	return nil
}
