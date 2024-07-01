package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {
	var PORT = getEnvVariable("PORT")
	r := mux.NewRouter()
	InitRoutes(r)

	fmt.Println("Starting server on port:", PORT)
	if err := http.ListenAndServe(":" + PORT, r); err != nil {
		log.Fatal("Server down: ", err)
	}
}

func getEnvVariable(key string) string {
	err := godotenv.Load("../../config/.env")
	if err != nil {
		fmt.Println("Error loading .env")
	}
	return os.Getenv(key)
}