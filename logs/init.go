// @program:     web-template
// @file:        init.go
// @author:      edte
// @create:      2020-12-19 20:18
// @description: 
package logs

import (
	"fmt"
	"io"
	"log"
	"os"

	"web-template/config"

	"github.com/sirupsen/logrus"
)

// Begin 设置了 logs 的默认配置
func Begin() *logrus.Logger {
	conf := config.LogConf()

	// 写入文件
	src, err := os.OpenFile(conf.Path, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		log.Printf("failed to open logs file: %v\n", err)
	}
	fmt.Println(err)

	// 实例化
	logger := logrus.New()

	// 设置输出
	logger.Out = src

	// 设置日志格式
	logger.SetFormatter(&logrus.TextFormatter{})

	// 设置方法名
	logger.SetReportCaller(true)

	// 同时打印日志和存到文件中
	logger.SetOutput(io.MultiWriter(os.Stdout, src))

	// 设置日志格式
	logger.SetFormatter(&logrus.JSONFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
	})

	return logger
}
