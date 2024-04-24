package config

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

/*
config:

int pingClientNum
int pingEnodeNum
int pingBnodeNum
bool pingIsTest
int pingBatchNum

[]string pingClientIpPort
[]string pingEnodeIpPort
[]string pingBnodeIpPort

*/

type PingClientConfig struct {
	PingClientNum    int
	PingEnodeNum     int
	PingBnodeNum     int
	PingIsTest       bool
	PingBatchNum     int
	PingClientIpPort []string
	PingEnodeIpPort  []string
	PingBnodeIpPort  []string
}

type ConfigArg interface{}

func getConfig(arg ConfigArg) (config PingClientConfig) {
	ok := false
	switch arg := arg.(type) {
	case string:
		config, ok = getConfigFromFile(arg)
		if !ok {
			fmt.Println("Failed to get config from file, getting from user")
			config = getConfigFromUser()
		}
	case []string:
		config, ok = getConfigFromArgs(arg)
		if !ok {
			fmt.Println("Failed to get config from args, getting from user")
			config = getConfigFromUser()
		}
	default:
		fmt.Println("Unsupported argument type")
	}
}

func getConfigFromFile(filename string) (config PingClientConfig, ok bool) {
	fmt.Println("Getting config from file:", filename)
	// 判断是不是json，是则打开
	if filepath.Ext(filename) == ".json" {
		file, err := os.Open(filename)
		if err != nil {
			fmt.Println("Error opening file:", err)
			return
		}
		defer file.Close()
		// json读到PingClientConfig，并返回
		err = json.NewDecoder(file).Decode(&config)
		if err != nil {
			fmt.Println("Error decoding JSON:", err)
			return config, false
		}
		return config, true
	} else {
		fmt.Println("File is not in JSON format")
		return config, false
	}
}

func getConfigFromArgs(args []string) (config PingClientConfig, ok bool) {
	fmt.Println("Getting config from args:", args)
	return config, true
}

func getConfigFromUser() (config PingClientConfig) {
	fmt.Println("Getting config from user")
	return config
}

