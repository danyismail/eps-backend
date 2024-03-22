package main

import (
	"eps-backend/db"
	"eps-backend/handler"
	"eps-backend/router"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {

	//set output log file
	file, err := os.OpenFile("app.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal("Failed to open log file:", err)
	}

	err = godotenv.Load()
	if err != nil {
		log.Fatalln("error loading .env file")
	}

	//connection pooling
	listConn, err := db.New()
	if err != nil {
		log.Fatalln(err)
	}

	defer file.Close()

	r := router.New(file)
	v1 := r.Group("/api")

	h := handler.NewHandler(listConn, r)
	h.Register(v1)
	h.HttpErrorHandler(r)

	appHost := os.Getenv("APPLICATION_PORT")
	if appHost == "" {
		log.Fatalln("app port is not define.")
	}
	r.Logger.Fatal(r.Start(":" + appHost))
}
