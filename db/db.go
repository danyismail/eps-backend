package db

import (
	"log"
	"os"

	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type DBConnection struct {
	DigiAmazone *gorm.DB
	DigiEps     *gorm.DB
	Amazone     *gorm.DB
	Eps         *gorm.DB
}

func New() (DBConnection, error) {

	//DATABASE
	DB_HOST := os.Getenv("DB_HOST")
	DB_HOST_REPLICA := os.Getenv("DB_HOST_REPLICA")
	DB_PORT := os.Getenv("DB_PORT")
	//EPS
	DATABASE_EPS := os.Getenv("DB_EPS")
	PASSWORD_EPS := os.Getenv("DB_PASSWORD_EPS")
	//AMAZONE
	DATABASE_AMAZONE := os.Getenv("DB_AMZ")
	PASSWORD_AMAZONE := os.Getenv("DB_PASSWORD_AMZ")
	//REPLICA
	DATABASE_AMAZONE_REPLICA := os.Getenv("DB_REPLICA_AMAZONE")
	DATABASE_EPS_REPLICA := os.Getenv("DB_REPLICA_EPS")
	USER_AMAZONE_REPLICA := os.Getenv("DB_USER_REPLICA")
	USER_EPS_REPLICA := os.Getenv("DB_USER_EPS")
	PASSWORD_AMAZONE_REPLICA := os.Getenv("DB_PASSWORD_REPLICA_AMAZONE")
	PASSWORD_EPS_REPLICA := os.Getenv("DB_PASSWORD_REPLICA_EPS")

	if DB_HOST == "" || DB_HOST_REPLICA == "" || DB_PORT == "" {
		log.Fatalln("database credentials not define.")
	}

	listConnDB := []string{
		"server=" + DB_HOST + ";user id=" + DATABASE_AMAZONE + ";password=" + PASSWORD_AMAZONE + ";encrypt=disable;database=" + DATABASE_AMAZONE,
		"server=" + DB_HOST + ";user id=" + DATABASE_EPS + ";password=" + PASSWORD_EPS + ";encrypt=disable;database=" + DATABASE_EPS,
		"server=" + DB_HOST_REPLICA + ";user id=" + USER_AMAZONE_REPLICA + ";password=" + PASSWORD_AMAZONE_REPLICA + ";encrypt=disable;database=" + DATABASE_AMAZONE_REPLICA,
		"server=" + DB_HOST_REPLICA + ";user id=" + USER_EPS_REPLICA + ";password=" + PASSWORD_EPS_REPLICA + ";encrypt=disable;database=" + DATABASE_EPS_REPLICA,
	}

	instanceDB := DBConnection{}
	for i, v := range listConnDB {
		switch i {
		case 0:
			db, err := setConnectionDB("digi_amazone", v)
			if err != nil {
				// return instanceDB, err
			}
			instanceDB.DigiAmazone = db
		case 1:
			db, err := setConnectionDB("digi_eps", v)
			if err != nil {
				// return instanceDB, err
			}
			instanceDB.DigiEps = db
		case 2:
			db, err := setConnectionDB("replica_amazone", v)
			if err != nil {
				return instanceDB, err
			}
			instanceDB.Amazone = db
		case 3:
			db, err := setConnectionDB("replica_eps", v)
			if err != nil {
				return instanceDB, err
			}
			instanceDB.Eps = db
		}
	}
	return instanceDB, nil
}

func setConnectionDB(from, strConn string) (db *gorm.DB, err error) {
	db, err = gorm.Open(sqlserver.Open(strConn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Error),
	})
	if err != nil {
		log.Printf("error connection from %s with message %v", from, err)
		return nil, err
	}
	log.Printf("success connection from %s", from)
	return db, nil
}
