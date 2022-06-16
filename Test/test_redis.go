package Test

import (
	"fmt"
	"gin-api-demo/Cores/redis"
	"github.com/gin-gonic/gin"
)

func Test(ctx *gin.Context) {
	result, err := redis.Rdb.Get(ctx, "111").Result()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(result)
}
