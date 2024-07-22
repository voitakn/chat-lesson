package main

import (
	"chat-db/internal/config"
	"chat-db/internal/services"
	"chat-db/internal/utils"
	"chat-db/pkg/pgdb"
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
)

func init() {
	if err := config.Load(".env"); err != nil {
		log.Fatal("Didn`t read .env config")
		return
	}
	ctx := context.Background()

	if err := pgdb.New(ctx, os.Getenv("DB_DSN")); err != nil {
		log.Fatal("@[main] can't init service s3client: ", err)
		return
	}
}

func main() {
	port := os.Getenv("API_PORT")

	http.HandleFunc("/", checkService)
	http.HandleFunc("/person/create", services.PersonCreate)

	if err := http.ListenAndServe(fmt.Sprintf(":%s", port), nil); err != nil {
		log.Panic("Error starting server: " + err.Error())
	}
}

func checkService(w http.ResponseWriter, r *http.Request) {
	utils.ResponseJson(w, []byte(`{"success":true, "app": "messages-api"}`))
}
