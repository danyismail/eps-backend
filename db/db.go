package db

import (
	"log"
	"os"

	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

func New() (*gorm.DB, *gorm.DB, error) {

	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_URL")

	if "" == dbHost || "" == dbPort {
		log.Fatalln("Database credentials not define.")
	}

	connString := "server=" + dbHost + ";user id=digi_eps;password=htLlph3lYkrqfaRTqKEELxaf;encrypt=disable;database=digi_eps"
	devDb, err := gorm.Open(sqlserver.Open(connString), &gorm.Config{})

	if err != nil {
		return nil, nil, err
	}

	connStringProd := "server=" + dbHost + ";user id=digi_amz;password=htLlph3lYkrqfaRTqKEELxaf;encrypt=disable;database=digi_amz"
	prodDb, err := gorm.Open(sqlserver.Open(connStringProd), &gorm.Config{})

	if err != nil {
		return nil, nil, err
	}

	return devDb, prodDb, nil

}
