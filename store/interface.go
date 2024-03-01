package store

import "eps-backend/model"

type KpiStore interface {
	//data, totalSelectedData, view per page, err
	FindAll(startDt, endDt string, page int, view int, mdn string, status int, shift string) (data *[]model.VKpis, attribute model.AttributeKPI, err error)
	FindAllProd(startDt, endDt string, page int, view int, mdn string, status int, shift string) (data *[]model.VKpis, attribute model.AttributeKPI, err error)
	Test() (*[]model.VKpis, error)
}

type DepositStore interface {
	GetBalanceToday() ([]model.CurrentDeposit, error)
	GetBalanceTodayProd() ([]model.CurrentDeposit, error)
}

type SalesStore interface {
	GetSalesToday() ([]model.SalesReport, error)
	GetSalesTodayProd() ([]model.SalesReport, error)
	Sales(from, to string) ([]model.SalesReport, error)
	SalesProd(from, to string) ([]model.SalesReport, error)
}

type SupplierStore interface {
	CreateSupplierEps(s model.Supplier) error
	CreateSupplierAmz(s model.Supplier) error
	GetSuppliersEps() ([]model.Supplier, error)
	GetSuppliersAmz() ([]model.Supplier, error)
	GetSupplierByIdEps(id int) (model.Supplier, error)
	GetSupplierByIdAmz(id int) (model.Supplier, error)
	UpdateSuppliersEps(model.Supplier) error
	UpdateSuppliersAmz(model.Supplier) error
	DeleteSupplierEps(id int) error
	DeleteSupplierAmz(id int) error
}

type DepositNote interface {
	Create(notes model.DepositNote) error
	Update(notes model.DepositNote) error
	GetById(id int) (*model.DepositNote, error)
}
