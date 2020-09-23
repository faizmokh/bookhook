package hooks

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"html"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"
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
			return
		}

		key, found := FindUrl(e.Entities.Urls)
		if found {
			log.Println("info: url is ", e.Entities.Urls[key].ExpandedURL)
			err = sendMessage(e.Entities.Urls[key].ExpandedURL)
			if err != nil {
				log.Fatal("error: unable to send message to telegram channel")
			}
			return
		}

		key, found = FindUrl(e.QuotedStatus.ExtendedTweet.Entities.Urls)
		if found {
			log.Println("info: url is", e.QuotedStatus.ExtendedTweet.Entities.Urls[key].ExpandedURL)
			err = sendMessage(e.QuotedStatus.ExtendedTweet.Entities.Urls[key].ExpandedURL)
			if err != nil {
				log.Fatal("error: unable to send message to telegram channel")
			}
			return
		}
	default:
		fmt.Fprintln(w, "go away!")
	}
}

func sendMessage(text string) error {
	htmlText := fmt.Sprintf("<a href\\=\"%s\">%s</a>", html.EscapeString(text), html.EscapeString(text))
	body := &SendMessage{
		Text:                  htmlText,
		ChatID:                os.Getenv("CHAT_ID"),
		DisableWebPagePreview: false,
		DisableNotification:   true,
		ParseMode:             "HTML",
	}
	buf := new(bytes.Buffer)
	json.NewEncoder(buf).Encode(body)

	timeout := time.Duration(5 * time.Second)
	client := &http.Client{
		Timeout: timeout,
	}

	req, err := http.NewRequest("POST", fmt.Sprintf("https://api.telegram.org/bot%s/sendMessage", os.Getenv("BOT_TOKEN")), buf)
	req.Header.Add("Content-Type", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
		return err
	}

	log.Println("info: telegram response", string(b))

	return err
}
