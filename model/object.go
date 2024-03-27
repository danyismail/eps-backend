package model

import (
	"time"

	"gorm.io/gorm"
)

type VKpis struct {
	TglEntri    time.Time `json:"tanggal_entri"`
	TglStatus   time.Time `json:"tanggal_status"`
	TglTempo    time.Time `json:"tanggal_tempo"`
	KodeProduk  string    `json:"kode_produk"`
	Tujuan      string    `json:"tujuan"`
	Status      string    `json:"status"`
	WaktuRespon string    `json:"waktu_respon"`
	Kpi         int       `json:"kpi"`
	Shift       string    `json:"shift"`
}

type AttributeKPI struct {
	Total   int64
	View    int64
	Success int64
	Failed  int64
}

type CurrentDeposit struct {
	KodeModul      string  `json:"kode_modul"`
	Label          string  `json:"label"`
	TotalTransaksi int64   `json:"total_transaksi"`
	PemakaianSaldo float64 `json:"pemakaian_saldo"`
	SaldoSekarang  float64 `json:"saldo_sekarang"`
}

type SalesReport struct {
	Trx       int64   `json:"trx"`
	Pembelian float64 `json:"pembelian"`
	Penjualan float64 `json:"penjualan"`
	Laba      float64 `json:"laba"`
}

type Supplier struct {
	gorm.Model
	Name   string `json:"name"`
	Status string `json:"status"`
}

type DepositNote struct {
	ID                 int        `json:"id"`
	CreatedAt          string     `json:"created_at"`
	UpdatedAt          string     `json:"updated_at"`
	DeletedAt          *time.Time `json:"deleted_at"`
	Name               string     `json:"name"`
	Supplier           string     `json:"supplier"`
	Amount             float64    `json:"amount"`
	OriginAccount      string     `json:"origin_account"`
	DestinationAccount string     `json:"destination_account"`
	ImageUpload        string     `json:"image_upload"`
	Reply              string     `json:"reply"`
	Status             string     `json:"status"`
}
