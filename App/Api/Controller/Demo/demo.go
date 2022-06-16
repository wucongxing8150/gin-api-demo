package Demo

import (
	"fmt"
	"gin-api-demo/Cores/redis"
	"github.com/gin-gonic/gin"
	"time"
)

func List(c *gin.Context) {
	result, _ := redis.Rdb.Set(c, "test", "hello word!", 1*time.Hour).Result()

	fmt.Println(result)

	result_1, err_1 := redis.Rdb.Get(c, "test").Result()
	if err_1 != nil {
		fmt.Println(err_1.Error())
		return
	}
	fmt.Println(result_1)
}
