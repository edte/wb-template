// @program:     web-template
// @file:        init.go
// @author:      edte
// @create:      2020-12-19 20:18
// @description: 
package config

import (
	"io/ioutil"
	"log"
	"os"

	"gopkg.in/yaml.v2"
)

// Init 初始化默认配置
func Init() {
	// 读取配置文件
	envFile := configFileName
	// 读取配置文件, 解决跑测试的时候找不到配置文件的问题，最多往上找5层目录
	for i := 0; i < 5; i++ {
		if _, err := os.Stat(envFile); err == nil {
			break
		} else {
			envFile = "../" + envFile
		}
	}

	yamlFile, err := ioutil.ReadFile(envFile)
	if err != nil {
		log.Printf("failed to read config file:%s", err)
		return
	}

	err = yaml.Unmarshal(yamlFile, &conf)
	if err != nil {
		log.Printf("failed to unmarshal config file:%s", err)
		return
	}

	// 初始化 log file
	logFile := LogConf().Path
	for i := 0; i < 5; i++ {
		if _, err := os.Stat(logFile); err == nil {
			break
		} else {
			envFile = "../" + logFile
		}
	}

	_, err = os.Create(logFile)
	if err != nil {
		log.Printf("failed to : %v\n", err)
		return
	}
}
