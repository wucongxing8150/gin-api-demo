package Auth

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

type UserInfo struct {
	Level      int         `json:"level"`
	NickName   string      `json:"nickName"`
	Sex        interface{} `json:"sex"`
	Mobile     interface{} `json:"mobile"`
	Pic        interface{} `json:"pic"`
	BirthDate  interface{} `json:"birthDate"`
	UserID     string      `json:"userId"`
	Score      int         `json:"score"`
	Balance    float64     `json:"balance"`
	UserMobile string      `json:"userMobile"`
	LevelType  int         `json:"levelType"`
	Growth     int         `json:"growth"`
	Status     int         `json:"status"`
	Username   string      `json:"username"`
}

// 免验证接口
var allow []string

// VerifyAuth
// @Description: 鉴权中间件
// @return gin.HandlerFunc
func VerifyAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		RequestURI := c.Request.RequestURI

		fmt.Println(RequestURI)

		for _, v := range allow {
			if RequestURI == v {
				c.Next()
				return
			}
		}

		c.Next()
	}
}

func init() {
	// 无需验证token的接口
	allow = []string{
		// "/demo/demo/list",
	}
}
