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

// VerifyAuth
// @Description: 鉴权中间件
// @return gin.HandlerFunc
func VerifyAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		RequestURI := c.Request.RequestURI

		fmt.Println(RequestURI)

		// 免验证接口
		allowInterface := []string{
			// "/",
		}

		for _, v := range allowInterface {
			if RequestURI == v {
				return
			}
		}

		c.Next()
	}
}
