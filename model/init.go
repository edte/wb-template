// @program:     web-template
// @file:        init.go
// @author:      edte
// @create:      2020-12-19 20:20
// @description: 
package model

// Init 初始化 mysql
func Init() {
	connectDb()
	initTable()
}
