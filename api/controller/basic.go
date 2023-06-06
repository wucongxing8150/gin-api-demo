package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"hospital-admin-api/api/responses"
	"hospital-admin-api/utils"
	"regexp"
	"strings"
)

type Basic struct{}

// GetCaptcha 获取验证码
func (b *Basic) GetCaptcha(c *gin.Context) {
	t := utils.Token{}
	t.UserId = 1
	t.RoleId = 1
	t.DeptId = 1
	t.PostId = 1

	au, err := t.NewJWT()
	if err != nil {
		responses.FailWithMessage(err.Error(), c)
		return
	}

	responses.OkWithData(au, c)
	// global.Logger.WithFields(logrus.Fields{
	// 	"name":   "key",
	// 	"values": "value",
	// }).Info("测试")
	//
	// result, err := global.Redis.Get(c, "111").Result()
	// if err != nil {
	// 	fmt.Println(err.Error())
	// 	responses.Fail(c)
	// 	return
	// }
	// fmt.Println(result)

	responses.Ok(c)
}

// GetCaptchaTest 获取验证码
func (b *Basic) GetCaptchaTest(c *gin.Context) {
	// path := "/admin/basic/captcha-test/:id"
	// url := c.Request.RequestURI
	//
	// method := "get"
	//
	// if KeyMatch2(url, path) && "get" == method {
	// 	responses.Ok(c)
	// 	return
	// }
	// responses.Fail(c)
	// return

	// au, err := utils.NewJWT(123456)
	// if err != nil {
	// 	responses.FailWithMessage(err.Error(), c)
	// }
	//
	// responses.OkWithData(au, c)
	// global.Logger.WithFields(logrus.Fields{
	// 	"name":   "key",
	// 	"values": "value",
	// }).Info("测试")
	//
	// result, err := global.Redis.Get(c, "111").Result()
	// if err != nil {
	// 	fmt.Println(err.Error())
	// 	responses.Fail(c)
	// 	return
	// }
	// fmt.Println(result)

	responses.Ok(c)
}

func KeyMatch2(key1 string, key2 string) bool {
	key2 = strings.Replace(key2, "/*", "/.*", -1)
	fmt.Println(key2)
	re := regexp.MustCompile(`:[^/]+`)
	key2 = re.ReplaceAllString(key2, "$1[^/]+$2")

	return RegexMatch(key1, "^"+key2+"$")
}

func RegexMatch(key1 string, key2 string) bool {
	res, err := regexp.MatchString(key2, key1)
	if err != nil {
		panic(err)
	}
	return res
}
