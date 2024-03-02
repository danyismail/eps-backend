package store

import (
	"eps-backend/model"
	"errors"
	"fmt"

	"gorm.io/gorm"
)

type SupplierConstruct struct {
	dev  *gorm.DB
	prod *gorm.DB
}

func NewSupplierStore(dev *gorm.DB, prod *gorm.DB) *SupplierConstruct {
	return &SupplierConstruct{
		dev,
		prod,
	}
}

func (c *SupplierConstruct) CreateSupplierEps(s model.Supplier) error {
	if err := c.dev.Debug().Create(&s).Error; err != nil {
		return err
	}
	return nil
}

func (c *SupplierConstruct) CreateSupplierAmz(s model.Supplier) error {
	if err := c.prod.Debug().Create(&s).Error; err != nil {
		return err
	}
	return nil
}

func (c *SupplierConstruct) GetSuppliersEps() ([]model.Supplier, error) {
	suppliers := []model.Supplier{}
	sql := "SELECT * FROM suppliers"
	if err := c.dev.Debug().Raw(sql).Scan(&suppliers).Error; err != nil {
		return nil, err
	}
	return suppliers, nil
}

func (c *SupplierConstruct) GetSuppliersEActive() ([]model.Supplier, error) {
	suppliers := []model.Supplier{}
	sql := "SELECT * FROM suppliers WHERE status = 'active'"
	if err := c.dev.Debug().Raw(sql).Scan(&suppliers).Error; err != nil {
		return nil, err
	}
	return suppliers, nil
}

func (c *SupplierConstruct) GetSuppliersAmz() ([]model.Supplier, error) {
	suppliers := []model.Supplier{}
	sql := "SELECT * FROM suppliers"
	if err := c.prod.Debug().Raw(sql).Scan(&suppliers).Error; err != nil {
		return nil, err
	}
	return suppliers, nil
}

func (c *SupplierConstruct) GetSuppliersAActive() ([]model.Supplier, error) {
	suppliers := []model.Supplier{}
	sql := "SELECT * FROM suppliers WHERE status = 'active'"
	if err := c.prod.Debug().Raw(sql).Scan(&suppliers).Error; err != nil {
		return nil, err
	}
	return suppliers, nil
}

func (c *SupplierConstruct) GetSupplierByIdEps(id int) (model.Supplier, error) {
	supplier := model.Supplier{}
	if err := c.dev.Debug().Table("suppliers").Where("id = ?", id).Scan(&supplier).Error; err != nil {
		return supplier, err
	}
	if supplier.ID == 0 {
		return supplier, errors.New("data not found")
	}
	return supplier, nil
}

func (c *SupplierConstruct) GetSupplierByIdAmz(id int) (model.Supplier, error) {
	supplier := model.Supplier{}
	if err := c.prod.Debug().Table("suppliers").Where("id = ?", id).Scan(&supplier).Error; err != nil {
		return supplier, err
	}
	if supplier.ID == 0 {
		return supplier, errors.New("data not found")
	}
	return supplier, nil
}

func (c *SupplierConstruct) UpdateSuppliersEps(s model.Supplier) error {
	if err := c.dev.Debug().Updates(&s).Error; err != nil {
		return err
	}
	return nil
}

func (c *SupplierConstruct) UpdateSuppliersAmz(s model.Supplier) error {
	if err := c.prod.Debug().Updates(&s).Error; err != nil {
		return err
	}
	return nil
}

func (c *SupplierConstruct) DeleteSupplierEps(id int) error {
	sql := fmt.Sprintf("DELETE FROM suppliers WHERE id = %d", id)
	if err := c.dev.Debug().Exec(sql).Error; err != nil {
		return err
	}
	return nil
}

func (c *SupplierConstruct) DeleteSupplierAmz(id int) error {
	sql := fmt.Sprintf("DELETE FROM suppliers WHERE id = %d", id)
	if err := c.prod.Debug().Exec(sql).Error; err != nil {
		return err
	}
	return nil
}
