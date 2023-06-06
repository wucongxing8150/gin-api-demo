package model

// AdminMenuApi 后台-菜单-接口表
type AdminMenuApi struct {
	MenuId int64     `gorm:"column:menu_id;type:bigint(19);primary_key;comment:菜单id" json:"menu_id"`
	ApiId  int64     `gorm:"column:api_id;type:bigint(19);primary_key;comment:接口id" json:"api_id"`
	Menu   AdminMenu `gorm:"foreignkey:MenuId;association_foreignkey:MenuID"`
	API    AdminAPI  `gorm:"foreignkey:ApiId;association_foreignkey:APIID"`
}

func (m *AdminMenuApi) TableName() string {
	return "gdxz_admin_menu_api"
}
