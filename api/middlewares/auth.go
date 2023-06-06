package middlewares

import (
	"github.com/gin-gonic/gin"
	"hospital-admin-api/api/dao"
	"hospital-admin-api/api/responses"
	"hospital-admin-api/consts"
	"net/http"
	"strings"
)

// Auth Auth认证
func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 获取角色id
		roleId := c.GetInt64("RoleId")
		if roleId == 0 {
			responses.Fail(c)
			c.Abort()
			return
		}

		// 获取用户id
		userId := c.GetInt64("UserId")
		if userId == 0 {
			responses.Fail(c)
			c.Abort()
			return
		}

		// 获取用户数据
		AdminUserDao := dao.AdminUser{}
		adminUser, err := AdminUserDao.GetAdminUserFirstById(userId)
		if err != nil || adminUser.UserId == 0 {
			responses.FailWithMessage("用户数据错误", c)
			c.Abort()
			return
		}

		if adminUser.Status == 2 {
			responses.FailWithMessage("用户审核中", c)
			c.Abort()
			return
		}

		if adminUser.Status == 3 {
			responses.FailWithMessage("用户已删除或禁用", c)
			c.Abort()
			return
		}

		// 获取角色数据
		AdminRoleDao := dao.AdminRole{}
		adminRole, err := AdminRoleDao.GetAdminRoleFirstById(roleId)
		if err != nil || adminRole.RoleId == 0 {
			responses.FailWithMessage("用户数据错误", c)
			c.Abort()
			return
		}

		// 超级管理员不验证权限
		if adminRole.IsAdmin == 1 {
			c.Next()
			return
		}

		// 获取角色菜单
		AdminRoleMenuDao := dao.AdminRoleMenu{}
		adminRoleMenu, _ := AdminRoleMenuDao.GetAdminMenuListByRoleId(roleId)
		if adminRoleMenu == nil {
			c.JSON(http.StatusForbidden, gin.H{
				"message": "暂无权限",
				"code":    consts.CLIENT_HTTP_UNAUTHORIZED,
				"data":    "",
			})
			c.Abort()
			return
		}

		var apiPermissions = make(map[string]bool)

		// 获取菜单对应api
		AdminMenuApiDao := dao.AdminMenuApi{}
		for _, v := range adminRoleMenu {
			AdminMenuApi, _ := AdminMenuApiDao.GetAdminMenuApiListByMenuID(v.MenuID)
			if AdminMenuApi == nil {
				// 菜单无需权限
				c.Next()
				return
			}

			// 将API权限存储在apiPermissions中
			for _, api := range AdminMenuApi {
				apiPermissions[api.API.APIPath+api.API.APIMethod] = true
			}
		}

		path := ""
		// 找到最后一个数字的索引
		lastSlashIndex := strings.LastIndex(c.Request.RequestURI, "/")
		if lastSlashIndex != -1 {
			// 替换最后一个数字部分为 :id
			path = c.Request.RequestURI[:lastSlashIndex] + "/:id" + c.Request.Method

		} else {
			c.JSON(http.StatusOK, gin.H{
				"message": "请求路径错误",
				"code":    consts.SERVER_ERROR,
				"data":    "",
			})

			c.Abort()
			return
		}

		// 在apiPermissions中查找对应的API权限
		hasPermission := apiPermissions[path]
		if !hasPermission {
			c.JSON(http.StatusForbidden, gin.H{
				"message": "暂无权限",
				"code":    consts.CLIENT_HTTP_UNAUTHORIZED,
				"data":    "",
			})

			c.Abort()
			return
		}

		c.Next()
	}
}

// Auth 权限
// func Auth() gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		fmt.Println(123)

//
// 		// result, err := dao.AdminRole.GetAdminRoleById(roleId)
// 		// fmt.Println(result)
// 		// if err != nil {
// 		// 	responses.FailWithMessage("用户数据错误", c)
// 		// 	c.Abort()
// 		// 	return
// 		// }
// 		// responses.OkWithData(&result, c)
// 		// c.Abort()
//
// 		// 获取请求路径
// 		// url := c.Request.RequestURI
// 		c.Next()
// 	}
// }
