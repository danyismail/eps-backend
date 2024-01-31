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
	d := db.New()

	r := router.New()
	v1 := r.Group("/api")

	h := handler.NewHandler(d)
	h.Register(v1)
	h.HttpErrorHandler(r)

	appHost := os.Getenv("APPLICATION_PORT")

	if "" == appHost {
		log.Fatalln("key of APPLICATION HOST are not define.")
	}

	r.Logger.Fatal(r.Start(":" + appHost))
}
