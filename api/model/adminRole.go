package model

// AdminRole 后台-角色表
type AdminRole struct {
	Model
	RoleId     int64  `gorm:"column:role_id;type:bigint(19);primary_key;comment:主键id" json:"role_id"`
	RoleName   string `gorm:"column:role_name;type:varchar(100);comment:角色名称" json:"role_name"`
	RoleStatus int    `gorm:"column:role_status;type:tinyint(1);default:1;comment:角色状态（1:正常 2:禁用）" json:"role_status"`
	IsAdmin    int    `gorm:"column:is_admin;type:tinyint(1);default:0;comment:是否管理员（0:否 1:是）" json:"is_admin"`
}

func (m *AdminRole) TableName() string {
	return "gdxz_admin_role"
}
