package RateLimit

import (
	"github.com/gin-gonic/gin"
	"github.com/juju/ratelimit"
	"net/http"
	"time"
)

// RateLimit
// @Description: 令牌桶限流
func RateLimit(fillInterval time.Duration, cap int64) func(c *gin.Context) {
	// 每个 fillInterval 一个令牌的速率填充，直到给定的最大容量
	bucket := ratelimit.NewBucket(fillInterval, cap)
	return func(c *gin.Context) {
		// 如果取不到令牌 直接返回
		if bucket.TakeAvailable(1) == 0 {
			c.String(http.StatusOK, "rate limit")
			c.Abort()
			return
		}
		c.Next()
	}
}
