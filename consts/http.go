package consts

// 业务状态码
const (
	HTTP_SUCCESS = 200 // 成功

	HTTP_ERROR = -1 // 失败

	USER_STATUS_ERROR = 201 // 用户状态异常

	CLIENT_HTTP_ERROR = 400 // 错误请求

	CLIENT_HTTP_UNAUTHORIZED = 401 // 未授权

	HTTP_PROHIBIT = 403 // 禁止请求

	CLIENT_HTTP_NOT_FOUND = 404 // 资源未找到

	TOKEN_ERROR = 405 // token错误/无效

	TOKEN_EXPTIRED = 406 // token过期

	SERVER_ERROR = 500 // 服务器异常
)
