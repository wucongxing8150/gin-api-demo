package dao

import (
	"hospital-admin-api/api/model"
	"hospital-admin-api/global"
)

type AdminMenuApi struct {
}

// GetAdminMenuApiListByMenuID 菜单id获取菜单api
// menuId 菜单id
func (r *AdminMenuApi) GetAdminMenuApiListByMenuID(menuId int64) (m []*model.AdminMenuApi, err error) {
	err = global.Db.Where("menu_id = ?", menuId).Preload("API").Find(&m).Error
	if err != nil {
		return nil, err
	}
	return m, nil
}
