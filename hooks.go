package hooks

import (
	"fmt"
	"net/http"
	"os"
)

func TwitterWebhook(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Yo!")

	_ = os.Getenv("ACCESS_TOKEN")
	_ = os.Getenv("ACCESS_TOKEN_SECRET")
}
