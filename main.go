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

	err := godotenv.Load()
	if err != nil {
		log.Fatalln("Error loading .env file")
	}
	devDb, prodDb, err := db.New()
	if err != nil {
		log.Fatalln(err)
	}

	r := router.New()
	v1 := r.Group("/api")

	h := handler.NewHandler(devDb, prodDb)
	h.Register(v1)
	h.HttpErrorHandler(r)

	appHost := os.Getenv("APPLICATION_PORT")

	if "" == appHost {
		log.Fatalln("key of APPLICATION HOST are not define.")
	}

	r.Logger.Fatal(r.Start(":" + appHost))
}
