package main

import (
	"fmt"
	"os"
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
	fmt.Println("sup nigga")
	fmt.Println("screwie ready, ^C exits")
	session := mongo.Start()
	defer session.Close()
	for {
	    stock.PrintStock("test")
	}
}
