package Recover

import (
	"log"
	"net/http"
	"runtime/debug"

	"github.com/gin-gonic/gin"
)

// Recover
// @Description: 处理全局异常
// @return gin.HandlerFunc
func Recover() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if r := recover(); r != nil {
				// 打印错误堆栈信息
				log.Printf("panic: %v\n", r)
				debug.PrintStack()
				c.JSON(http.StatusOK, gin.H{
					"code":    "1",
					"message": errorToString(r),
					"data":    "",
				})
				// 终止后续接口调用，不加的话recover到异常后，还会继续执行接口里后续代码
				c.Abort()
			}
		}()
		// 加载完 defer recover，继续后续接口调用
		c.Next()
	}
}

// recover错误，转string
func errorToString(r interface{}) string {
	switch v := r.(type) {
	case error:
		return v.Error()
	default:
		return r.(string)
	}
}
