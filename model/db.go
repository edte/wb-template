// @program:     web-template
// @file:        db.go
// @author:      edte
// @create:      2020-12-19 20:21
// @description: 
package model

import (
	"fmt"

	"web-template/config"
	"web-template/logs"

	"github.com/jinzhu/gorm"
)

var (
	db *gorm.DB
)

// connectDb 连接 mysql
func connectDb() {
	conf := config.MysqlConf()

	d, err := gorm.Open(conf.Type, fmt.Sprintf("%s:%s@(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", conf.User, conf.Password, conf.Host, conf.Port, conf.Name))
	if err != nil {
		logs.Begin().Fatalf("failed to connect database:%v", err)
	}

	db = d
}
