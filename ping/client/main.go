package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
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
	PingClientNum    int      `json:"pingClientNum"`
	PingEnodeNum     int      `json:"pingEnodeNum"`
	PingBnodeNum     int      `json:"pingBnodeNum"`
	PingIsTest       bool     `json:"pingIsTest"`
	PingBatchNum     int      `json:"pingBatchNum"`
	PingClientIpPort []string `json:"pingClientIpPort"`
	PingEnodeIpPort  []string `json:"pingEnodeIpPort"`
	PingBnodeIpPort  []string `json:"pingBnodeIpPort"`
}

type ConfigArg interface{}

func getConfig(arg ConfigArg) (config PingClientConfig) {
	ok := false
	switch arg := arg.(type) {
	case string:
		config, ok = getConfigFromFile(arg)
		if !ok {
			fmt.Println("Failed to get config from file, getting from user")
			// config = getConfigFromUser()
		}
	default:
		fmt.Println("Unsupported argument type")
	}
	return
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

// func main() {
// 	getConfig("config.json")
// 	fmt.Println("Hello, ping-client!")
// }

// 1. 读取配置文件
// 2. 连接e
// 3. 循环
//    1. 消息输入
//    2. 消息组装
//       1. 用消息格式匹配下
//       2. 获取同步信息（所以同步click是个独立线程）
//       3. 组装
//    3. 扔进发送池（所以消息发送也是个独立线程）
//    4. （所以消息返回后处理也是个独立线程）

func main() {
	// 1. 读取配置文件
	config := PingClientConfig{
		PingClientNum:    1,
		PingEnodeNum:     32,
		PingBnodeNum:     32,
		PingIsTest:       false,
		PingBatchNum:     1,
		PingClientIpPort: []string{"10.1.0.28:22001"},
		PingEnodeIpPort: []string{
			"10.1.0.112:22001",
			// 其他的enode地址
		},
		PingBnodeIpPort: []string{
			"10.1.0.36:22002",
			// 其他的bnode地址
		},
	}
	// 2. 连接e
	connection2Enode := make([])
	
}
