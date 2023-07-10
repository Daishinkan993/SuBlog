package main

import (
	"SuBlog/DB"
	"SuBlog/routing"
	"fmt"
	"github.com/joho/godotenv"
	"github.com/rs/cors"
	"log"
	"net/http"
	"os"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	port := os.Getenv("PORT")
	dbUrl := os.Getenv("DB_URL")
	dbName := os.Getenv("DB_NAME")

	db, err := DB.Init(dbUrl, dbName)
	if err != nil {
		log.Fatal(err)
	}

	handler := cors.Default().Handler(routing.Init(db))

	fmt.Println("http://127.0.0.1" + port + "/")

	log.Fatal(http.ListenAndServe(port, handler))
}
