package model

type BrandRevenue struct {
	Server    string  `json:"server"`
	Member    int     `json:"member"`
	Trx       int     `json:"trx"`
	Penjualan float64 `json:"penjualan"`
	Pembelian float64 `json:"pembelian"`
	Tekor     float64 `json:"tekor"`
	Bakar     float64 `json:"bakar"`
	Laba      float64 `json:"laba"`
	Komisi    float64 `json:"komisi"`
	Ppn11     float64 `json:"ppn11"`
	Pph22     float64 `json:"pph22"`
}
