package Demo

import (
	"github.com/gin-gonic/gin"
)

func List(c *gin.Context) {
	// redis操作示例
	// result, _ := redis.Rdb.Set(c, "test", "hello word!", 1*time.Hour).Result()
	//
	// fmt.Println(result)
	//
	// result_1, err_1 := redis.Rdb.Get(c, "test").Result()
	// if err_1 != nil {
	// 	fmt.Println(err_1.Error())
	// 	return
	// }

	// jwt-token示例
	// 创建
	// claims := JwtToken.CustomJwtStruct{
	// 	UserId:   "123456",
	// 	UserName: "test",
	// 	StandardClaims: jwt.StandardClaims{
	// 		ExpiresAt: int64(time.Now().Unix() + 3600),
	// 		Issuer:    "wucongxing",
	// 	},
	// }
	//
	// token, err := JwtToken.CreateToken(claims)
	// fmt.Println(token)
	// fmt.Println(err)

	// 获取token的上下文数据
	// claims, _ := c.Keys["token"].(*JwtToken.CustomJwtStruct)
	// fmt.Println(claims)
	// fmt.Println(claims.UserId)

}
