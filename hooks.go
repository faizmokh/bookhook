package hooks

import (
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func TwitterWebHook(w http.ResponseWriter, r *http.Request) {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("error: unable to load .env file")
	}

	accessToken := os.Getenv("ACCESS_TOKEN")
	accessTokenSecret := os.Getenv("ACCESS_TOKEN_SECRET")

	log.Println("ACCESS TOKEN: ", accessToken)
	log.Println("ACCESS TOKEN SECRET: ", accessTokenSecret)
}
