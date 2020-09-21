package main

import (
	"log"
	"net/http"

	hooks "github.com/faizmokhtar/bookhook"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("error: unable to load .env file")
	}

	http.HandleFunc("/", hooks.TwitterWebhook)
	log.Println("Listening on localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
