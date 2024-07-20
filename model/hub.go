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

type BrandCategoryRevenue struct {
	Provider    string  `json:"provider"`
	JenisProduk string  `json:"jenis_produk"`
	Trx         int     `json:"trx"`
	Laba        float64 `json:"laba"`
}

type BrandCategoryData struct {
	JenisProduk string  `json:"jenis_produk"`
	Trx         int     `json:"trx"`
	Laba        float64 `json:"laba"`
}

type Response struct {
	Category  []BrandCategoryData `json:"category"`
	TotalTrx  int                 `json:"total_trx"`
	TotalLaba float64             `json:"total_laba"`
}

type DetailResponse struct {
	Response map[string]Response `json:"response"`
	SubTotal int                 `json:"sub_total"`
	SubLaba  float64             `json:"total"`
}
