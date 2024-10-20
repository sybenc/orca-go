package response

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"orca/pkg/code"
	"orca/pkg/erorrs"
)

// Success 表示本次请求成功，并返回请求信息和HTTP状态
func Success(c *gin.Context, data any, message string) {
	c.JSON(http.StatusOK, gin.H{
		"code":    code.Success,
		"data":    data,
		"status":  http.StatusOK,
		"message": message,
	})
}

func Successf(c *gin.Context, data any, message string, args ...any) {
	c.JSON(http.StatusOK, gin.H{
		"code":    code.Success,
		"data":    data,
		"status":  http.StatusOK,
		"message": fmt.Sprintf(message, args...),
	})
}

// Fail 表示本次请求失败，并返回请求失败的错误码和错误信息
func Fail(c *gin.Context, err error) {
	fmt.Printf("%+v\n", err)
	coder := errors.ParseCoder(err)
	c.JSON(coder.HttpStatus(), gin.H{
		"code":      coder.Code(),
		"data":      nil,
		"status":    coder.HttpStatus(),
		"message":   coder.Message(),
		"reference": coder.Reference(),
	})
}
