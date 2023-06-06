package model

// AdminApi 后台-接口表
type AdminAPI struct {
	Model
	APIID     int64  `gorm:"column:api_id;type:bigint(19);primary_key;AUTO_INCREMENT;comment:主键id" json:"api_id"`
	APIName   string `gorm:"column:api_name;type:varchar(100);comment:api名称;NOT NULL" json:"api_name"`
	APIPath   string `gorm:"column:api_path;type:varchar(255);comment:接口路径（全路径 id为:id）;NOT NULL" json:"api_path"`
	APIMethod string `gorm:"column:api_method;type:varchar(10);comment:请求类型（put:修改 post:新增 get:获取 ）;NOT NULL" json:"api_method"`
}

func (m *AdminAPI) TableName() string {
	return "gdxz_admin_api"
}
