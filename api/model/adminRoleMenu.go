package model

// AdminRoleMenu 后台-角色-菜单表
type AdminRoleMenu struct {
	RoleID int64     `gorm:"column:role_id;type:bigint(19);primary_key;comment:权限id" json:"role_id"`
	MenuID int64     `gorm:"column:menu_id;type:bigint(19);primary_key;comment:菜单id" json:"menu_id"`
	Menu   AdminMenu `gorm:"foreignKey:MenuID"`
	Role   AdminRole `gorm:"foreignKey:RoleID"`
}

func (m *AdminRoleMenu) TableName() string {
	return "gdxz_admin_role_menu"
}
