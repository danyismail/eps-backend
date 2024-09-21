package model

type RevenuePerHour struct {
	TglEntri  string  `json:"tgl_entri"`
	DariJam   string  `json:"dari_jam"`
	SdJam     string  `json:"sd_jam"`
	Seq       int     `json:"seq"`
	JumlahTrx int     `json:"jumlah_trx"`
	Laba      float64 `json:"laba"`
	TotalTrx  float64 `json:"total_trx"`
	TotalLaba float64 `json:"total_laba"`
}
