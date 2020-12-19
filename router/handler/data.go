// @program:     web-template
// @file:        data.go
// @author:      edte
// @create:      2020-12-19 20:31
// @description: 
package handler

import (
	"github.com/gin-gonic/gin"

	"web-template/logs"
	"web-template/router/response"
	"web-template/router/service"
)

func AddData(c *gin.Context) {
	var f service.Form

	if err := c.ShouldBindJSON(&f); err != nil {
		logs.Begin().Panicln(err)
	}

	service.AddData()

	response.Ok(c)
}
