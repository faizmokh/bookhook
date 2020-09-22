package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"

	"github.com/dghubble/oauth1"
	hooks "github.com/faizmokhtar/bookhook"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("error: unable to load .env file")
	}

	if args := os.Args; len(args) > 1 && args[1] == "-register" {
		go registerWebhook()
	}

	if args := os.Args; len(args) > 1 && args[1] == "-subscribe" {
		go subscribeWebhook()
	}

	http.HandleFunc("/", hooks.TwitterWebhook)
	log.Println("Listening on localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func registerWebhook() {
	fmt.Println("Registering webhook...")
	httpClient := createClient()

	//Set parameters
	path := "https://api.twitter.com/1.1/account_activity/all/" + os.Getenv("WEBHOOK_ENV") + "/webhooks.json"
	values := url.Values{}
	values.Set("url", os.Getenv("APP_URL"))

	//Make Oauth Post with parameters
	resp, _ := httpClient.PostForm(path, values)
	defer resp.Body.Close()
	//Parse response and check response
	body, _ := ioutil.ReadAll(resp.Body)
	var data map[string]interface{}
	if err := json.Unmarshal([]byte(body), &data); err != nil {
		panic(err)
	}
	fmt.Println(data)
	fmt.Println("Webhook id of " + data["id"].(string) + " has been registered")
}

func subscribeWebhook() {
	fmt.Println("Subscribing webhook...")
	client := createClient()
	path := "https://api.twitter.com/1.1/account_activity/all/" + os.Getenv("WEBHOOK_ENV") + "/subscriptions.json"
	resp, _ := client.PostForm(path, nil)
	body, _ := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	if resp.StatusCode == 204 {
		fmt.Println("Subscribed successfully")
	} else if resp.StatusCode != 204 {
		fmt.Println("Could not subscribe the webhook. Response below:")
		fmt.Println(string(body))
	}
}

func createClient() *http.Client {
	//Create oauth client with consumer keys and access token
	config := oauth1.NewConfig(os.Getenv("CONSUMER_KEY"), os.Getenv("CONSUMER_SECRET"))
	token := oauth1.NewToken(os.Getenv("ACCESS_TOKEN_KEY"), os.Getenv("ACCESS_TOKEN_SECRET"))
	return config.Client(oauth1.NoContext, token)
}
