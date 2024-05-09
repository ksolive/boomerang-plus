package config

import (
	"fmt"
)

type DevConfig struct {
	BotNum int
}

type BotConfig struct {
	BotNum       int
	AppId        string
	ClientSecret string
}

func LoadDevConfig() *DevConfig {
	fmt.Println("loadConfig ")
	return &DevConfig{BotNum: 10}
}

func LoadBotConfig() *BotConfig {
	fmt.Println("loadConfig ")
	return &BotConfig{BotNum: 10, AppId: "102108663", ClientSecret: "fSF2pdRF3rfTI7wlaPE3tjZPF5vmdULC"}
}
