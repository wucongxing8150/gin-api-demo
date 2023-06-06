package model

// AdminMenu 后台-菜单表
type AdminMenu struct {
	Model
	MenuId     int64  `gorm:"column:menu_id;type:bigint(19);primary_key;comment:主键id" json:"menu_id"`
	MenuName   string `gorm:"column:menu_name;type:varchar(30);comment:菜单名称" json:"menu_name"`
	ParentId   int    `gorm:"column:parent_id;type:int(10);default:0;comment:父菜单ID（0表示一级）" json:"parent_id"`
	MenuStatus int    `gorm:"column:menu_status;type:tinyint(1);default:1;comment:菜单状态（0:隐藏 1:正常）此优先级最高" json:"menu_status"`
	MenuType   int    `gorm:"column:menu_type;type:tinyint(1);comment:菜单类型（1:模块 2:菜单 2:按钮）" json:"menu_type"`
	Permission string `gorm:"column:permission;type:varchar(255);comment:标识" json:"permission"`
	OrderNum   int    `gorm:"column:order_num;type:int(4);default:0;comment:显示顺序" json:"order_num"`
	Icon       string `gorm:"column:icon;type:varchar(255);comment:图标地址" json:"icon"`
	Path       string `gorm:"column:path;type:varchar(255);comment:页面地址（#表示当前页）" json:"path"`
}

func (m *AdminMenu) TableName() string {
	return "gdxz_admin_menu"
}
