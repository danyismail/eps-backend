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
