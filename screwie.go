package main

import (
	"fmt"
	"os"
	"log"
	"strings"
	"screwie/slack"
	"screwie/mongo"
	"screwie/stock"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "usage: screwie slack-bot-token\n")
		os.Exit(1)
	}

	ws, id := slack.SlackConnect(os.Args[1])
	fmt.Println("Screwie is ready :), ^C exits")
	session := mongo.Start()
	defer session.Close()
	for {
	    msg, err := slack.GetMessage(ws)
	    if err != nil {
	        log.Fatal(err)
	    }
	    if msg.Type == "message" && strings.HasPrefix(msg.Text, "<@"+id+">") {
            parts := strings.Fields(msg.Text)
            if(parts[1] == "stock"){
                go func(m slack.Message){
                    m.Text = stock.GetStock(parts[2])
                    slack.PostMessage(ws, m)
                }(msg)
            } else if parts[1] == "monitor" {
                if parts[2] == "stock" {
                } else if parts[2] == "movie" {
                }
            } else {
                go func(m slack.Message) {
                    m.Text = fmt.Sprintf("Sorry, dude, unknown command :(\n")
                    slack.PostMessage(ws, m)
                }(msg)
            }
	    }
	}
}
