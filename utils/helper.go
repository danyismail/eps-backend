package utils

import (
	"eps-backend/db"
	"fmt"
	"strings"

	"gorm.io/gorm"
)

func FindPath(path string) string {
	// Split the string by "/"
	parts := strings.Split(path, "/")
	fmt.Println(parts)

	return parts[2]
}

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
	default:
		return conn.DigiAmazone
	}
}
