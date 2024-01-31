package db

import (
	"log"
	"os"

	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

func New() *gorm.DB {

	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_URL")

	if "" == dbHost || "" == dbPort {
		log.Fatalln("Database credentials not define.")
	}

	// db, err := gorm.Open(dbDriver, dbURL)

	// if err != nil {
	// 	log.Fatalln(err)
	// }

	// db.DB().SetMaxIdleConns(3)

	// if f, _ := strconv.ParseBool(os.Getenv("DB_LOG")); f {
	// 	db.LogMode(true)
	// 	db.SetLogger(l)
	// }

	// return db
	// Connection string
	// connString := "server=36.64.19.114;user id=digi_eps;password=htLlph3lYkrqfaRTqKEELxaf;encrypt=disable;database=digi_eps"
	// //connString := "sqlserver://sa:digi_eps@36.64.19.114:1433?database=digi_eps&encrypt=disable&connection+timeout=30"

	// // Connect to the database
	// db, err := sql.Open("sqlserver", connString)
	// if err != nil {
	// 	fmt.Println("Error connecting to database:", err.Error())
	// }
	// defer db.Close()

	// // Perform a query
	// rows, err := db.Query("SELECT TOP 10 * FROM v_kpi")
	// if err != nil {
	// 	fmt.Println("Error executing query:", err.Error())
	// }
	// defer rows.Close()
	// return db

	// github.com/denisenkom/go-mssqldb
	// dsn := "sqlserver://gorm:digi_eps@36.64.19.114:1433?database=digi_eps&encrypt=disable"
	connString := "server=" + dbHost + ";user id=digi_eps;password=htLlph3lYkrqfaRTqKEELxaf;encrypt=disable;database=digi_eps"
	db, err := gorm.Open(sqlserver.Open(connString), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}

	return db
}
