package store

import (
	"eps-backend/model"
)

type KpiStore interface {
	//data, totalSelectedData, view per page, err
	GetAll(path, startDt, endDt string, page int, view int, mdn string, status int, shift string) (data *[]model.VKpis, attribute model.AttributeKPI, err error)
}

type DepositStore interface {
	GetBalance(conn string) ([]model.CurrentDeposit, error)
}

type SalesStore interface {
	GetSales(path string) ([]model.SalesReport, error)
	GetSalesPeriode(path, from, to string) ([]model.SalesReport, error)
}

type SupplierStore interface {
	GetSuppliers(path string) ([]model.Supplier, error)
	GetActiveSuppliers(path string) ([]model.Supplier, error)
	CreateSupplier(path string, spl model.Supplier) error
	GetSupplierById(path string, id int) (model.Supplier, error)
	UpdateSupplier(path string, spl model.Supplier) error
	DeleteSupplier(path string, id int) error
}

type DepositNote interface {
	GetAllStatus(path, date string) ([]model.DepositNote, error)
	GetStatusCreated(path string) ([]model.DepositNote, error)
	GetStatusUploaded(path string) ([]model.DepositNote, error)
	GetStatusDone(path, startDt, endDt string) ([]model.DepositNote, error)
	Create(path string, notes model.DepositNote) error
	GetById(path string, id int) (*model.DepositNote, error)
	Update(path string, notes model.DepositNote) error
	Delete(path string, id int) error
}
