package dao

import (
	"hospital-admin-api/api/model"
	"hospital-admin-api/global"
)

type AdminRoleMenu struct {
}

// GetAdminMenuListByRoleId GetAdminRoleById 角色id获取用户角色
// roleId 角色id
func (r *AdminRoleMenu) GetAdminMenuListByRoleId(roleId int64) (m []*model.AdminRoleMenu, err error) {
	err = global.Db.Where("role_id = ?", roleId).Preload("Menu").Find(&m).Error
	if err != nil {
		return nil, err
	}
	return m, nil
}
