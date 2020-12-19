// @program:     web-template
// @file:        init.go
// @author:      edte
// @create:      2020-12-19 20:29
// @description: 
package router

import (
	"web-template/logs"
	"web-template/router/handler"

	"github.com/gin-gonic/gin"
)

var (
	routes = []struct {
		name    string
		method  string
		path    string
		handler gin.HandlerFunc
	}{
		{
			name:    "add data",
			method:  "/data",
			path:    "POST",
			handler: handler.AddData,
		},
	}
)

// Init 初始化路由
func Init() {
	r := gin.Default()

	for _, route := range routes {
		r.Handle(route.method, route.path, route.handler)
	}

	logs.Begin().Infoln(r.Run())
}
