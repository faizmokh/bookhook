package main

import (
	"log"
	"net/http"

	hooks "github.com/faizmokhtar/bookhook"
)

func main() {
	http.HandleFunc("/", hooks.TwitterWebhook)
	log.Println("Listening on localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
