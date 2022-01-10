package main

import (
	"app/config"
	router "app/routes"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func init() {
	loadEnvFile()
}

func main() {

	port := config.Server().Port

	muxRouter := mux.NewRouter().StrictSlash(true)
	router.Routes(muxRouter)

	log.Println("Server started at port: ", port)

	log.Fatal(http.ListenAndServe(":"+port, muxRouter))
}

func loadEnvFile() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}
