package hooks

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
)

func TwitterWebhook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	switch r.Method {
	case "GET":
		log.Println("info: challenge-response check")
		token := r.URL.Query()["crc_token"]
		if len(token) < 1 {
			log.Println("error: no crc_token given")
			return
		}

		h := hmac.New(sha256.New, []byte(os.Getenv("ACCESS_TOKEN_SECRET")))
		h.Write([]byte(token[0]))
		encoded := base64.StdEncoding.EncodeToString(h.Sum(nil))

		response := make(map[string]string)
		response["response_token"] = "sha256=" + encoded

		responseJSON, _ := json.Marshal(response)
		fmt.Fprintf(w, string(responseJSON))
	case "POST":
	default:
		fmt.Fprintln(w, "go away!")
	}

	_ = os.Getenv("ACCESS_TOKEN")
	_ = os.Getenv("ACCESS_TOKEN_SECRET")
}
