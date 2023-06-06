package middlewares

import (
	"github.com/gin-gonic/gin"
	"hospital-admin-api/api/responses"
	"hospital-admin-api/consts"
	"hospital-admin-api/global"
	"hospital-admin-api/utils"
	"net/http"
	"strings"
)

// Jwt jwt认证
func Jwt() gin.HandlerFunc {
	return func(c *gin.Context) {
		authorization := c.Request.Header.Get("Authorization")
		if authorization == "" || !strings.HasPrefix(authorization, "Bearer ") {
			c.JSON(http.StatusUnauthorized, gin.H{
				"message": "请求未授权",
				"code":    consts.TOKEN_ERROR,
				"data":    "",
			})
			c.Abort()
			return
		}

		// 去除Bearer
		authorization = authorization[7:] // 截取字符

		// 检测是否存在黑名单
		res, _ := global.Redis.Get(c, "jwt_black_"+authorization).Result()
		if res != "" {
			responses.Fail(c)

			c.Abort()
			return
		}

		// 解析jwt
		t, err := utils.ParseJwt(authorization)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"message": "token错误/过期",
				"code":    consts.TOKEN_ERROR,
				"data":    "",
			})

			c.Abort()
			return
		}

		c.Set("UserId", t.UserId) // 用户id
		c.Set("RoleId", t.RoleId) // 角色id
		c.Set("DeptId", t.DeptId) // 部门id
		c.Set("PostId", t.PostId) // 岗位id
		c.Next()
	}
}
