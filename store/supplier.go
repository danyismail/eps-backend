package store

import (
	"eps-backend/db"
	"eps-backend/model"
	"errors"
	"fmt"
)

type SupplierConstruct struct {
	db db.DBConnection
}

func NewSupplierStore(db db.DBConnection) *SupplierConstruct {
	return &SupplierConstruct{db}
}

func (c *SupplierConstruct) CreateSupplierEps(s model.Supplier) error {
	if err := c.db.DigiEps.Create(&s).Error; err != nil {
		return err
	}
	return nil
}

func (c *SupplierConstruct) CreateSupplierAmz(s model.Supplier) error {
	if err := c.db.DigiAmazone.Create(&s).Error; err != nil {
		return err
	}
	return nil
}

func (c *SupplierConstruct) GetSuppliersEps() ([]model.Supplier, error) {
	suppliers := []model.Supplier{}
	sql := "SELECT * FROM suppliers"
	if err := c.db.DigiEps.Raw(sql).Scan(&suppliers).Error; err != nil {
		return nil, err
	}
	return suppliers, nil
}

func (c *SupplierConstruct) GetSuppliersEActive() ([]model.Supplier, error) {
	suppliers := []model.Supplier{}
	sql := "SELECT * FROM suppliers WHERE status = 'active'"
	if err := c.db.DigiAmazone.Raw(sql).Scan(&suppliers).Error; err != nil {
		return nil, err
	}
	return suppliers, nil
}

func (c *SupplierConstruct) GetSuppliersAmz() ([]model.Supplier, error) {
	suppliers := []model.Supplier{}
	sql := "SELECT * FROM suppliers"
	if err := c.db.DigiAmazone.Raw(sql).Scan(&suppliers).Error; err != nil {
		return nil, err
	}
	return suppliers, nil
}

func (c *SupplierConstruct) GetSuppliersAActive() ([]model.Supplier, error) {
	suppliers := []model.Supplier{}
	sql := "SELECT * FROM suppliers WHERE status = 'active'"
	if err := c.db.DigiAmazone.Raw(sql).Scan(&suppliers).Error; err != nil {
		return nil, err
	}
	return suppliers, nil
}

func (c *SupplierConstruct) GetSupplierByIdEps(id int) (model.Supplier, error) {
	supplier := model.Supplier{}
	if err := c.db.DigiEps.Table("suppliers").Where("id = ?", id).Scan(&supplier).Error; err != nil {
		return supplier, err
	}
	if supplier.ID == 0 {
		return supplier, errors.New("data not found")
	}
	return supplier, nil
}

func (c *SupplierConstruct) GetSupplierByIdAmz(id int) (model.Supplier, error) {
	supplier := model.Supplier{}
	if err := c.db.DigiAmazone.Table("suppliers").Where("id = ?", id).Scan(&supplier).Error; err != nil {
		return supplier, err
	}
	if supplier.ID == 0 {
		return supplier, errors.New("data not found")
	}
	return supplier, nil
}

func (c *SupplierConstruct) UpdateSuppliersEps(s model.Supplier) error {
	if err := c.db.DigiEps.Updates(&s).Error; err != nil {
		return err
	}
	return nil
}

func (c *SupplierConstruct) UpdateSuppliersAmz(s model.Supplier) error {
	if err := c.db.DigiAmazone.Updates(&s).Error; err != nil {
		return err
	}
	return nil
}

func (c *SupplierConstruct) DeleteSupplierEps(id int) error {
	sql := fmt.Sprintf("DELETE FROM suppliers WHERE id = %d", id)
	if err := c.db.DigiEps.Exec(sql).Error; err != nil {
		return err
	}
	return nil
}

func (c *SupplierConstruct) DeleteSupplierAmz(id int) error {
	sql := fmt.Sprintf("DELETE FROM suppliers WHERE id = %d", id)
	if err := c.db.DigiAmazone.Exec(sql).Error; err != nil {
		return err
	}
	return nil
}
