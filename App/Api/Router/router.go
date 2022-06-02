package Router

import (
	"fmt"
	"net/http"

	"gin-api-demo/App/Api/Middlewares/Auth"
	"gin-api-demo/App/Api/Middlewares/Corss"
	"gin-api-demo/App/Api/Middlewares/Recover"
	"github.com/gin-gonic/gin"
)

type Option func(*gin.Engine)

// 初始化
func Init() *gin.Engine {

	r := gin.Default()

	// 404处理
	r.NoRoute(func(c *gin.Context) {
		path := c.Request.URL.Path
		method := c.Request.Method
		c.JSON(http.StatusNotFound, gin.H{
			"msg":  fmt.Sprintf("%s %s not found", method, path),
			"code": 404,
			"data": "",
		})
	})

	// 处理异常
	r.Use(Recover.Recover())

	// 跨域中间件
	r.Use(Corss.Cors())

	// 鉴权中间件
	r.Use(Auth.VerifyAuth())

	return r
}
