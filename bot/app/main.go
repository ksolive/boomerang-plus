package main

import (
	"fmt"
	"time"

	"github.com/ksolive/bots/bots"
	"github.com/ksolive/bots/config"
)

func main() {
	fmt.Println("Hello, bots!")
	botsConfig := config.LoadBotConfig()
	bot := bots.NewBot("test", *botsConfig) // Dereference botsConfig
	go bot.Run()
	time.Sleep(10 * time.Second)
}
