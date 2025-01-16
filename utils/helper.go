package utils

import (
	"eps-backend/db"

	"gorm.io/gorm"
)

func SelectConn(path string, conn db.DBConnection) *gorm.DB {
	switch path {
	case DIGI_AMAZONE:
		return conn.DigiAmazone
	case DIGI_EPS:
		return conn.DigiEps
	case REPLICA_AMAZONE:
		return conn.Amazone
	case REPLICA_EPS:
		return conn.Eps
	case OTODEV:
		return conn.Otodev
	case SERVERONE:
		return conn.ServerOne
	case VALUEPULSA:
		return conn.ValuePulsa
	default:
		return conn.DigiAmazone
	}
}

func GetKPIConfig(path string) string {
	switch path {
	case DIGI_AMAZONE:
		return "60"
	case DIGI_EPS:
		return "60"
	case REPLICA_AMAZONE:
		return "180"
	case REPLICA_EPS:
		return "180"
	case BACKUP_AMAZONE:
		return "60"
	case OTODEV:
		return "120"
	default:
		return "180"
	}
}

func EmptyString(s string) bool {
	return s == ""
}

func StringPointer(s string) *string {
	return &s
}
