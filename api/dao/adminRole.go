package dao

import (
	"hospital-admin-api/api/model"
	"hospital-admin-api/global"
)

type AdminRole struct {
}

// GetAdminRoleFirstById 角色id获取用户角色
// roleId 角色id
func (r *AdminRole) GetAdminRoleFirstById(roleId int64) (m model.AdminRole, err error) {
	err = global.Db.First(&m, roleId).Error
	if err != nil {
		return m, err
	}
	return m, nil
}
