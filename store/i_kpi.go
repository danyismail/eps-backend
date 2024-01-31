package store

import "eps-backend/model"

type KpiStore interface {
	FindAll(page int, view int, mdn string) (*[]model.VKpis, int64, error)
	Test() (*[]model.VKpis, error)
}
