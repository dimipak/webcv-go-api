package main

import (
	c "app/config"
	router "app/routes"
	"log"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func init() {
	loadEnvFile()
	c.InitGlobals()
}

func main() {

	muxRouter := createServer()

	log.Println("Server Listening on Port: ", c.G_APP.Port)

	log.Fatal(http.ListenAndServe(":"+c.G_APP.Port, muxRouter))
}

func loadEnvFile() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func createServer() *mux.Router {
	cors := handlers.CORS(
		handlers.AllowedHeaders([]string{"*"}),
		handlers.AllowedOrigins([]string{"*"}),
		handlers.AllowCredentials(),
		handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "OPTIONS"}),
	)

	muxRouter := mux.NewRouter()
	muxRouter.Use(cors)
	router.Routes(muxRouter)

	return muxRouter
}
