package db

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

func New() (*gorm.DB, *gorm.DB, error) {

	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbEps := os.Getenv("DB_EPS")
	dbAmz := os.Getenv("DB_AMZ")

	if dbHost == "" || dbPort == "" {
		log.Fatalln("Database credentials not define.")
	}

	connString := "server=" + dbHost + ";user id=" + dbEps + ";password=htLlph3lYkrqfaRTqKEELxaf;encrypt=disable;database=" + dbEps
	fmt.Println(connString)
	devDb, err := gorm.Open(sqlserver.Open(connString), &gorm.Config{})

	if err != nil {
		return nil, nil, err
	}

	connStringProd := "server=" + dbHost + ";user id=" + dbAmz + ";password=htLlph3lYkrqfaRTqKEELxaf;encrypt=disable;database=" + dbAmz
	fmt.Println(connStringProd)
	prodDb, err := gorm.Open(sqlserver.Open(connStringProd), &gorm.Config{})

	if err != nil {
		return nil, nil, err
	}
	return devDb, prodDb, nil
}
