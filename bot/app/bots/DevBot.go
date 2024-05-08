package bots

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"time"

	"github.com/ksolive/bots/config"
	"github.com/ksolive/bots/mq"
)

// 最简单的模拟bot，不断向后输出随机结果
// bot 结构：消息获取，消息处理，消息后发
type Bot struct {
	name string
}

func NewBot(name string) *Bot {
	return &Bot{name: name}
}

func (b *Bot) getMessage() string {
	// 随机生成一个json string
	rand.Seed(time.Now().UnixNano())
	data := struct {
		Name string `json:"name"`
	}{Name: b.name}
	jsonData, _ := json.Marshal(data)
	return string(jsonData)
}

func (b *Bot) processMessage(msg string) string {
	// 无处理，直接返回
	return msg
}

func (b *Bot) sendMessage2Out(msg string) {
	// 输出
	fmt.Println(msg)
}

func (b *Bot) sendMessage2Mq(msg string) {
	// 输出到mq
	mq.Bots2Nlp(msg)
}

func (b *Bot) Run() {
	for {
		msg := b.getMessage()
		msg = b.processMessage(msg)
		b.sendMessage2Out(msg)
		b.sendMessage2Mq(msg)
		time.Sleep(1 * time.Second)
	}
}

func init() {
	params := config.LoadConfig()
	for i := 0; i < params.BotNum; i++ {
		bot := NewBot(fmt.Sprintf("bot-%d", i))
		go bot.Run()
	}
}
