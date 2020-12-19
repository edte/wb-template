// @program:     web-template
// @file:        response.go
// @author:      edte
// @create:      2020-12-19 20:32
// @description: 
package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// 业务码
const (
	CodeOk               = iota + 1000 // 正常响应
	CodeFormError                      // 请求表单错误
	CodeFormatError                    // 参数校验有问题
	CodeTemplateAddError               // 增加模板失败

	DataOk        = "ok"
	DataFormError = "request form error!"
)

// Ok
func Ok(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"code": CodeOk})
}

// FormError
func FormError(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"code": CodeFormError, "message": DataFormError})
}

// OkWithData
func OkWithData(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, gin.H{"code": CodeOk, "message": DataOk, "data": data})
}

// Error
func Error(c *gin.Context, code int, msg string) {
	c.JSON(http.StatusOK, gin.H{"code": code, "message": msg})
}
