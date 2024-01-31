package store

import "eps-backend/model"

type KpiStore interface {
	FindAll(page int, view int) (*[]model.VKpi, error)
	Test() (*[]model.VKpi, error)
}
