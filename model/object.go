package model

import "time"

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

type CurrentDeposit struct {
	KodeModul      string  `json:"kode_modul"`
	Label          string  `json:"label"`
	TotalTransaksi int64   `json:"total_transaksi"`
	PemakaianSaldo float64 `json:"pemakaian_saldo"`
	SaldoSekarang  float64 `json:"saldo_sekarang"`
}

type SalesReport struct {
	Ma        int64   `json:"ma"`
	Trx       int64   `json:"trx"`
	Pembelian float64 `json:"pembelian"`
	Penjualan float64 `json:"penjualan"`
}
