package config

import (
	"fmt"
)

type DevConfig struct {
	BotNum int
}

func LoadConfig() *DevConfig {
	fmt.Println("loadConfig ")
	return &DevConfig{BotNum: 10}
}
