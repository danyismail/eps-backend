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
		log.Fatalln("error loading .env file")
	}
	listConn, err := db.New()
	if err != nil {
		log.Fatalln(err)
	}

	r := router.New()
	v1 := r.Group("/api")

	h := handler.NewHandler(listConn)
	h.Register(v1)
	h.HttpErrorHandler(r)

	appHost := os.Getenv("APPLICATION_PORT")
	if appHost == "" {
		log.Fatalln("app port is not define.")
	}
	r.Logger.Fatal(r.Start(":" + appHost))
}
