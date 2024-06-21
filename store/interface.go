package store

import (
	"eps-backend/model"
	"eps-backend/structs"
)

type UserStore interface {
	Create(userRequest structs.CreateUser) (user *model.User, err error)
	GetAll(page, view int) (users []model.User, err error)
	Count() int64
	GetUser(param map[string]interface{}) (result *model.User, err error)
	Delete(id string) error
}

type KpiStore interface {
	GetAll(path, startDt, endDt string, page int, view int, mdn string, status int, shift string) (data *[]model.VKpis, attribute model.AttributeKPI, err error)
}

type DepositStore interface {
	GetBalance(conn string) ([]model.CurrentDeposit, error)
}

type SalesStore interface {
	GetSales(path string) ([]model.SalesReport, error)
	GetPPH(path string) ([]model.SalesReport, error)
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

type SNStore interface {
	GetNullableSN(path string) ([]model.ValidateSN, error)
	GetDuplicateSN(path string) ([]model.DuplicateSN, error)
}

type ResellerStore interface {
	GetLaba(path, startDt, endDt, id string) ([]model.LabaReseller, error)
	GetSum(path, startDt, endDt, id string) (*model.SumLabaReseller, error)
	GetList(path, arg string) ([]model.Reseller, error)
	GetLabaHourly(path string) ([][]model.CekLabaHourly, error)
	GetLabaRugi(path, from, to string) ([]model.CekLabaRugi, error)
}

type MarginStore interface {
	ByReseller(path, startDt, endDt string) ([]model.MarginReseller, error)
	BySupplier(path, startDt, endDt string) ([]model.MarginSupplier, error)
	ByProvider(path, startDt, endDt string) ([]model.MarginProvider, error)
}

type HubStore interface {
	GetBrandRevenue(startDt, endDt string) ([]model.BrandRevenue, error)
}
