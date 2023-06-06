package router

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"hospital-admin-api/api/controller"
	"hospital-admin-api/api/exception"
	"hospital-admin-api/api/middlewares"
	"hospital-admin-api/consts"
	"net/http"
)

// Init 初始化路由
func Init() *gin.Engine {
	r := gin.New()

	r.Use(middlewares.Logrus())
	r.Use(gin.Recovery())

	// 404处理
	r.NoRoute(func(c *gin.Context) {
		path := c.Request.URL.Path
		method := c.Request.Method
		c.JSON(http.StatusNotFound, gin.H{
			"msg":  fmt.Sprintf("%s %s not found", method, path),
			"code": consts.CLIENT_HTTP_NOT_FOUND,
			"data": "",
		})
	})

	// 异常处理
	r.Use(exception.Recover())

	// 跨域处理
	r.Use(middlewares.Cors())

	// 加载基础路由
	api := controller.Api{}

	// 公开路由-不验证权限
	publicRouter(r, api)

	// 验证jwt
	r.Use(middlewares.Jwt())

	// 验证权限
	r.Use(middlewares.Auth())

	// 注册私有路由
	privateRouter(r, api)

	return r
}

// publicRouter 公开路由-不验证权限
func publicRouter(r *gin.Engine, api controller.Api) {
	adminGroup := r.Group("/admin")
	baseGroup := adminGroup.Group("/basic")
	{
		// 验证码
		baseGroup.GET("captcha", api.Basic.GetCaptcha)
	}
}

// privateRouter 私有路由
func privateRouter(r *gin.Engine, api controller.Api) {
	adminGroup := r.Group("/admin")
	base1Group := adminGroup.Group("/basic")
	{
		// 验证码
		base1Group.GET("captcha-test/:id", api.Basic.GetCaptchaTest)
	}
}
