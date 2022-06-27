package Router

import (
	"fmt"
	"gin-api-demo/App/Api/Controller/Demo"
	"gin-api-demo/App/Api/Middlewares/RateLimit"
	"github.com/spf13/viper"
	"net/http"
	"time"

	"gin-api-demo/App/Api/Middlewares/Auth"
	"gin-api-demo/App/Api/Middlewares/Corss"
	"gin-api-demo/App/Api/Middlewares/Recover"
	"github.com/gin-gonic/gin"
)

type Option func(*gin.Engine)

// Init 初始化
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

	// 限流 令牌桶
	r.Use(RateLimit.RateLimit(viper.GetDuration("rate-limit.fill-interval")*time.Second, viper.GetInt64("rate-limit.capacity")))

	demo := r.Group("/demo")
	{
		// 邮寄地址
		demoGroup := demo.Group("/demo")
		{
			// 邮寄地址列表
			demoGroup.POST("list", Demo.List)
		}
	}

	return r
}
