package main

import (
	"fmt"

	"github.com/ksolive/bots/bots"
)

func main() {
	fmt.Println("Hello, bots!")
	bot := bots.NewBot("devBot")
	bot.Run()
}
