package JwtAuth

import (
	"gin-api-demo/Utils/JwtToken"
	"gin-api-demo/Utils/Response"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"strings"
)

// 免验证接口
var allow []string

func JwtAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		RequestURI := c.Request.RequestURI

		for _, v := range allow {
			if RequestURI == v {
				// 无需验证token接口
				c.Next()
				return
			}
		}

		authorization := c.Request.Header.Get("Authorization")

		if authorization == "" {
			Response.FailWithMessage("缺少token", c)
			c.Abort()
			return
		}

		bearer := viper.GetBool("jwt.bearer")
		if bearer == true {
			// 携带bearer头
			parts := strings.SplitN(authorization, " ", 2)
			if !(len(parts) == 2 && parts[0] == "Bearer") {
				Response.FailWithMessage("请求头中token格式有误", c)
				c.Abort()
				return
			}

			authorization = parts[1]
		}

		// 验证token
		token, err := JwtToken.ParseToken(authorization)

		if token == nil || err != nil {
			Response.FailWithMessage("token解析错误", c)
			c.Abort()
			return
		}

		// 写入gin.Context引用对象中
		c.Set("token", token)
		c.Next()
	}
}

func init() {
	// 无需验证token的接口
	allow = []string{
		// "/demo/demo/list",
	}
}
