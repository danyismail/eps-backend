package db

import (
	"fmt"
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
	Backup      *gorm.DB
}

func New() (DBConnection, error) {

	//DATABASE
	DB_HOST_DIGI := os.Getenv("DB_HOST_DIGI")
	DB_PORT_DIGI := os.Getenv("DB_PORT_DIGI")

	DB_DIGIAMAZONE := os.Getenv("DB_DIGIAMAZONE")
	USERNAME_DIGIAMAZONE := os.Getenv("USERNAME_DIGIAMAZONE")
	PASSWORD_DIGIAMAZONE := os.Getenv("PASSWORD_DIGIAMAZONE")

	DB_DIGIEPS := os.Getenv("DB_DIGIEPS")
	USERNAME_DIGIEPS := os.Getenv("USERNAME_DIGIEPS")
	PASSWORD_DIGIEPS := os.Getenv("PASSWORD_DIGIEPS")

	DB_HOST_REPLICA := os.Getenv("DB_HOST_REPLICA")
	DB_PORT_REPLICA := os.Getenv("DB_PORT_REPLICA")

	DB_REPLICA_AMAZONE := os.Getenv("DB_REPLICA_AMAZONE")
	USERNAME_REPLICA_AMAZONE := os.Getenv("DB_USER_REPLICA_AMAZONE")
	PASSWORD_REPLICA_AMAZONE := os.Getenv("DB_PASSWORD_REPLICA_AMAZONE")

	DB_REPLICA_EPS := os.Getenv("DB_REPLICA_EPS")
	USERNAME_REPLICA_EPS := os.Getenv("USERNAME_REPLICA_EPS")
	PASSWORD_REPLICA_EPS := os.Getenv("PASSWORD_REPLICA_EPS")

	DB_HOST_BACKUP := os.Getenv("DB_HOST_BACKUP")
	DB_PORT_BACKUP := os.Getenv("DB_PORT_BACKUP")
	DB_BACKUP := os.Getenv("DB_BACKUP")
	USERNAME_BACKUP := os.Getenv("USERNAME_BACKUP")
	PASSWORD_BACKUP := os.Getenv("PASSWORD_BACKUP")

	if DB_HOST_DIGI == "" || DB_HOST_REPLICA == "" || DB_HOST_BACKUP == "" {
		log.Fatalln("database credentials not define.")
	}

	listConnDB := []string{
		"server=" + DB_HOST_DIGI + "," + DB_PORT_DIGI + ";user id=" + USERNAME_DIGIAMAZONE + ";password=" + PASSWORD_DIGIAMAZONE + ";encrypt=disable;database=" + DB_DIGIAMAZONE,
		"server=" + DB_HOST_DIGI + "," + DB_PORT_DIGI + ";user id=" + USERNAME_DIGIEPS + ";password=" + PASSWORD_DIGIEPS + ";encrypt=disable;database=" + DB_DIGIEPS,
		"server=" + DB_HOST_REPLICA + "," + DB_PORT_REPLICA + ";user id=" + USERNAME_REPLICA_AMAZONE + ";password=" + PASSWORD_REPLICA_AMAZONE + ";encrypt=disable;database=" + DB_REPLICA_AMAZONE,
		"server=" + DB_HOST_REPLICA + "," + DB_PORT_REPLICA + ";user id=" + USERNAME_REPLICA_EPS + ";password=" + PASSWORD_REPLICA_EPS + ";encrypt=disable;database=" + DB_REPLICA_EPS,
		"server=" + DB_HOST_BACKUP + "," + DB_PORT_BACKUP + ";user id=" + USERNAME_BACKUP + ";password=" + PASSWORD_BACKUP + ";encrypt=disable;database=" + DB_BACKUP,
	}

	for _, v := range listConnDB {
		fmt.Println(v)
	}

	instanceDB := DBConnection{}
	for i, v := range listConnDB {
		switch i {
		case 0:
			db, err := setConnectionDB("digi_amazone", v)
			if err != nil {
				return instanceDB, err
			}
			instanceDB.DigiAmazone = db
		case 1:
			db, err := setConnectionDB("digi_eps", v)
			if err != nil {
				return instanceDB, err
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
		case 4:
			db, err := setConnectionDB("backup_amazone", v)
			if err != nil {
				return instanceDB, err
			}
			instanceDB.Backup = db
		}
	}
	log.Println("successfully create all conn..")
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
