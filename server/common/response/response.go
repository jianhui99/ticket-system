package response

import (
	"github.com/gin-gonic/gin"
)

// 输出数据结构体
type P struct {
	Code         int
	Message      string
	ErrorMessage map[string]string
	Params       gin.H
}

// Response 结构体
type R struct {
	// 输出类型，例如：JSON、XML、YAML
	T        string
	P        P
	HttpCode int
}

func ApiResponse(context *gin.Context, data interface{}, httpCode int, message string) {
	if message == "" {
		if httpCode >= 200 && httpCode < 300 {
			message = "Success"
		} else {
			message = "Fail"
		}
	}
	if httpCode == 0 {
		if data != nil {
			httpCode = 200
		} else {
			httpCode = 500
		}
	}

	context.JSON(httpCode, gin.H{
		"code":    httpCode,
		"message": message,
		"data":    data,
	})

	context.Abort()
}

func Success(context *gin.Context, data interface{}, message ...string) {
	msg := ""
	if len(message) > 0 {
		msg = message[0]
	}
	ApiResponse(context, data, 200, msg)
}

func Fail(context *gin.Context, httpCode int, message ...string) {
	msg := ""
	if len(message) > 0 {
		msg = message[0]
	}
	ApiResponse(context, nil, httpCode, msg)
}
