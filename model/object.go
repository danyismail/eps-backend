package model

import (
	"database/sql/driver"
	"encoding/hex"
	"fmt"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type CustomUUID uuid.UUID

func (cu *CustomUUID) Scan(value interface{}) error {
	b, ok := value.([]byte)
	if !ok {
		return fmt.Errorf("could not scan type %T into CustomUUID", value)
	}
	u, err := uuid.FromBytes(swapByteOrder(b))
	if err != nil {
		return err
	}
	*cu = CustomUUID(u)
	return nil
}

// Implement the Valuer interface for CustomUUID
func (cu CustomUUID) Value() (driver.Value, error) {
	u := uuid.UUID(cu)
	return swapByteOrder(u[:]), nil
}

func (cu CustomUUID) String() string {
	var buf [36]byte
	encodeHex(buf[:], cu)
	return string(buf[:])
}

func encodeHex(dst []byte, cu CustomUUID) {
	hex.Encode(dst, cu[:4])
	dst[8] = '-'
	hex.Encode(dst[9:13], cu[4:6])
	dst[13] = '-'
	hex.Encode(dst[14:18], cu[6:8])
	dst[18] = '-'
	hex.Encode(dst[19:23], cu[8:10])
	dst[23] = '-'
	hex.Encode(dst[24:], cu[10:])
}

// Swap the byte order for MSSQL
func swapByteOrder(b []byte) []byte {
	return []byte{
		b[3], b[2], b[1], b[0],
		b[5], b[4],
		b[7], b[6],
		b[8], b[9], b[10], b[11], b[12], b[13], b[14], b[15],
	}
}

type User struct {
	ID        CustomUUID `gorm:"type:uniqueidentifier;primaryKey"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`
	Username  string     `json:"username"`
	Email     string     `json:"email"`
	Role      string     `json:"role"`
	Password  string     `json:"password"`
}

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
	Pph       float64 `json:"pph"`
	LabaNet   float64 `json:"laba_net"`
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
	PaymentPurpose     string     `json:"payment_purpose"`
}

type ValidateSN struct {
	KodeReseller string `json:"kode_reseller"`
	KodeProduk   string `json:"kode_produk"`
	Tujuan       string `json:"tujuan"`
	TglEntri     string `json:"tgl_entri"`
	TglStatus    string `json:"tgl_status"`
	Label        string `json:"supplier"`
	SN           string `json:"sn"`
	SelisihWaktu string `json:"selisih_waktu"`
}

type DuplicateSN struct {
	SN     string `json:"sn"`
	Tujuan string `json:"tujuan"`
	Total  string `json:"total"`
}

type Reseller struct {
	Kode string `json:"kode"`
	Nama string `json:"nama"`
}

type LabaReseller struct {
	KodeProduk   string  `json:"kode_produk"`
	Trx          int     `json:"trx"`
	KodeReseller string  `json:"kode_reseller"`
	Laba         float64 `json:"laba"`
}

type ResellerParam struct {
	Search string
}

type SumLabaReseller struct {
	Nama string  `json:"nama"`
	Trx  int     `json:"trx"`
	Jual int     `json:"jual"`
	Beli int     `json:"beli"`
	Laba float64 `json:"laba"`
}

type CekLabaHourly struct {
	Tanggal string `json:"tanggal"`
	Tgl     int    `json:"tgl"`
	Jam     int    `json:"jam"`
	Trx     int    `json:"trx"`
	Laba    int    `json:"laba"`
}

type ResponseLabaPerJam struct {
	Aggregate map[int][]CekLabaHourly `json:"aggregate"`
	Data      [][]CekLabaHourly       `json:"data"`
}

type CekLabaRugi struct {
	Nama string  `json:"nama"`
	Trx  float64 `json:"trx"`
	Laba float64 `json:"laba"`
	Rugi float64 `json:"rugi"`
}

type MarginReseller struct {
	KodeReseller string  `json:"kode_reseller"`
	Nama         string  `json:"nama"`
	TotalTrx     int     `json:"trx"`
	Pembelian    float64 `json:"pembelian"`
	Penjualan    float64 `json:"penjualan"`
	Laba         float64 `json:"laba"`
}

type MarginSupplier struct {
	Label     string  `json:"label"`
	TotalTrx  int     `json:"trx"`
	Pembelian float64 `json:"pembelian"`
	Penjualan float64 `json:"penjualan"`
	Laba      float64 `json:"laba"`
}

type MarginProvider struct {
	Provider  string  `json:"provider"`
	TotalTrx  int     `json:"trx"`
	Pembelian float64 `json:"pembelian"`
	Penjualan float64 `json:"penjualan"`
	Laba      float64 `json:"laba"`
}

type Dbs struct {
	Name string
	Cnx  *gorm.DB
}

type PPH struct {
	TotalPph float64 `json:"total_pph"`
}
