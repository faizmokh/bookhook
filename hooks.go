package hooks

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func TwitterWebhook(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Yo!")

	err := godotenv.Load()
	if err != nil {
		log.Fatal("error: unable to load .env file")
	}

	_ = os.Getenv("ACCESS_TOKEN")
	_ = os.Getenv("ACCESS_TOKEN_SECRET")
}
