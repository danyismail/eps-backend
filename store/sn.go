package store

import (
	"eps-backend/db"
	"eps-backend/model"
	"eps-backend/utils"
)

type SNConstruct struct {
	db db.DBConnection
}

func NewSNStore(db db.DBConnection) *SNConstruct {
	return &SNConstruct{
		db: db,
	}
}

func (c *SNConstruct) GetNullableSN(path string) ([]model.ValidateSN, error) {
	var nullSN []model.ValidateSN
	query := `
		select
			kode_reseller, 
			kode_produk ,
			tujuan,
			tgl_entri,
			tgl_status,
			sn,
			CONCAT(
				LEFT(CONVERT(VARCHAR(8), DATEADD(SECOND, DATEDIFF(SECOND, tgl_status, GETDATE()), 0), 108), 2), ' jam ',
				SUBSTRING(CONVERT(VARCHAR(8), DATEADD(SECOND, DATEDIFF(SECOND, tgl_status, GETDATE()), 0), 108), 4, 2), ' menit'
			) AS selisih_waktu
		from
			transaksi t
		where
			(status = 20
				and 
		tgl_entri BETWEEN CAST(CONVERT(date,
				DATEADD(day, -1, GETDATE())) AS datetime)
					AND GETDATE())
			AND sn in(
			NULL,
			'N/A',
			'SALDO',
			'SAL',
			'HARGA',
			'HRG',
			'REF',
			'message',
			'0',
			'A',
			'PROGRESS',
			'UPDATE',
			'-',
			'RECON'
		);
	`
	if err := utils.SelectConn(path, c.db).Raw(query).Scan(&nullSN).Error; err != nil {
		return nil, err
	}
	return nullSN, nil
}

func (c *SNConstruct) GetDuplicateSN(path string) ([]model.DuplicateSN, error) {
	var duplicateSN []model.DuplicateSN
	query := `
		select
			sn,
			tujuan,
			count(sn) as total
		from
			(
			select
				*
			from
				transaksi
			where
				status = 20
				and tgl_entri BETWEEN CAST(CONVERT(date,
				DATEADD(day, -1, GETDATE())) AS datetime)
					AND GETDATE()
		) a
		group by
			a.sn,
			a.tujuan
		having
			a.sn not in (
			'N/A',
			'SALDO',
			'SAL',
			'HARGA',
			'HRG',
			'REF',
			'message',
			'0',
			'A',
			'PROGRESS',
			'UPDATE',
			'-',
			'RECON',
			'CEK BYU',
			''
		)
			and count(a.sn) > 1
		order by
			COUNT(a.sn) desc;
	`
	if err := utils.SelectConn(path, c.db).Raw(query).Scan(&duplicateSN).Error; err != nil {
		return nil, err
	}
	return duplicateSN, nil
}
