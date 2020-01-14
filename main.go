package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	// read message
	msg, err := ReadMessage("mattermost.json")

	if os.IsNotExist(err) {
		log.Fatalf("Missing ./mattermost.json file, a previous action should populate it.")
	}

	if err != nil {
		log.Fatalf("error reading message: %s", err)
	}

	// webhook
	webhook := os.Getenv("MATTERMOST_WEBHOOK_URL")

	if webhook == "" {
		log.Fatalf("Missing MATTERMOST_WEBHOOK_URL environment variable")
	}

	// channel
	if s := os.Getenv("MATTERMOST_CHANNEL"); s != "" {
		msg.ChannelName = s
	}

	// username
	if s := os.Getenv("MATTERMOST_USERNAME"); s != "" {
		msg.Username = s
	}

	// icon
	if s := os.Getenv("MATTERMOST_ICON"); s != "" {
		msg.IconURL = s
	}

	if isEmpty(msg) {
		log.Println("mattermost.json is empty exiting without failing.")
		os.Exit(0)
	}

	err = Send(webhook, msg)
	if err != nil {
		log.Fatalf("error sending message: %s", err)
	}

	fmt.Printf("Mattermost message sent!\n")
}
