package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"os"
	"os/user"
)

func username() string {
	username := "<unknown>"
	usr, err := user.Current()
	if err == nil {
		username = usr.Username
	}

	hostname := "<unknown>"
	host, err := os.Hostname()
	if err == nil {
		hostname = host
	}
	return fmt.Sprintf("%s@%s", username, hostname)
}

var cfg *Config

func RandomHandler(w http.ResponseWriter, r *http.Request) {
	username := r.FormValue("user_name")
	msg := SlackMsg{
		Channel:   cfg.Channel,
		Username:  *cfg.Username,
		Parse:     "full",
		Text:      fmt.Sprintf("%s's number : %d", username, rand.Intn(100)),
		IconEmoji: "",
	}

	err := msg.Post(cfg.WebhookUrl)
	if err != nil {
		log.Fatalf("Post failed: %v", err)
		fmt.Fprintf(w, "Post failed: %v", err)
	}
}

func main() {
	var err error
	cfg, err = ReadConfig()
	if err != nil {
		log.Fatalf("Could not read config: %v", err)
	}

	http.HandleFunc("/random", RandomHandler)
	err = http.ListenAndServe(":"+os.Getenv("PORT"), nil)
	if err != nil {
		panic(err)
	}
}
