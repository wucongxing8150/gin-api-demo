package dao

import (
	"hospital-admin-api/api/model"
	"hospital-admin-api/global"
)

type AdminUser struct {
}

// GetAdminUserFirstById 用户id获取用户数据
// roleId 用户id
func (r *AdminUser) GetAdminUserFirstById(userId int64) (m model.AdminUser, err error) {
	err = global.Db.First(&m, userId).Error
	if err != nil {
		return m, err
	}
	return m, nil
}
