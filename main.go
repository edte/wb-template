// @program:     web-template
// @file:        main.go
// @author:      edte
// @create:      2020-12-19 20:11
// @description: 
package main

import (
	"web-template/config"
	"web-template/model"
	"web-template/router"
)

func main() {
	config.Init()
	model.Init()
	router.Init()
}
