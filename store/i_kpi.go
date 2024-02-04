package store

import "eps-backend/model"

type KpiStore interface {
	FindAll(startDt, endDt string, page int, view int, mdn string, status int, shift string) (*[]model.VKpis, int64, error)
	Test() (*[]model.VKpis, error)
}
