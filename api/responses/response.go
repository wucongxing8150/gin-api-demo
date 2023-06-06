package responses

import (
	"github.com/gin-gonic/gin"
	"hospital-admin-api/consts"
	"net/http"
)

type Response struct {
	Code    int         `json:"code"`
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
}

func Result(code int, data interface{}, msg string, c *gin.Context) {
	if data == nil {
		data = gin.H{}
	}
	c.JSON(http.StatusOK, Response{
		code,
		data,
		msg,
	})
}

func Ok(c *gin.Context) {
	Result(consts.HTTP_SUCCESS, map[string]interface{}{}, "成功", c)
}

func OkWithMessage(message string, c *gin.Context) {
	Result(consts.HTTP_SUCCESS, map[string]interface{}{}, message, c)
}

func OkWithData(data interface{}, c *gin.Context) {
	Result(consts.HTTP_SUCCESS, data, "成功", c)
}

func OkWithDetailed(data interface{}, message string, c *gin.Context) {
	Result(consts.HTTP_SUCCESS, data, message, c)
}

func Fail(c *gin.Context) {
	Result(consts.HTTP_ERROR, map[string]interface{}{}, "失败", c)
}

func FailWithMessage(message string, c *gin.Context) {
	Result(consts.HTTP_ERROR, map[string]interface{}{}, message, c)
}

func FailWithDetailed(data interface{}, message string, c *gin.Context) {
	Result(consts.HTTP_ERROR, data, message, c)
}
