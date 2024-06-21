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
	case BACKUP_AMAZONE:
		return conn.Backup
	case OTODEV:
		return conn.Otodev
	default:
		return conn.DigiAmazone
	}
}

func EmptyString(s string) bool {
	return s == ""
}

func StringPointer(s string) *string {
	return &s
}
