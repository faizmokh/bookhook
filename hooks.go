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

type CRC struct {
	ResponseToken string `json:"response_token"`
}

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

		log.Println("info: crc token ", token)

		h := hmac.New(sha256.New, []byte(os.Getenv("CONSUMER_SECRET")))
		h.Write([]byte(token[0]))
		encoded := base64.StdEncoding.EncodeToString(h.Sum(nil))

		response := CRC{
			ResponseToken: "sha256=" + encoded,
		}

		responseJSON, _ := json.Marshal(response)
		log.Println("response: ", response)
		fmt.Fprintf(w, string(responseJSON))
	case "POST":
		log.Println("listening to twitter account activity")
		var t Tweet
		err := json.NewDecoder(r.Body).Decode(&t)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		e := t.TweetCreateEvents[0]
		_, found := FindHashtag(e.Entities.Hashtags, "toread")
		if !found {
			log.Println("info: no #toread hash tag found")
			return
		}

		key, found := FindUrl(e.Entities.Urls)
		if found {
			log.Println("info: entities expanded url is ", e.Entities.Urls[key].ExpandedURL)
		}

		key, found := FindUrl(e.QuotedStatus.Entities.Urls)
		if found {
			log.Println("info: quoted entities expanded url is", e.QuotedStatus.Entities.Urls[key].ExpandedURL)
		}
	default:
		fmt.Fprintln(w, "go away!")
	}
}
