package model

import "time"

// AdminDept 后台-部门表
type AdminDept struct {
	DeptId     int64     `gorm:"column:dept_id;type:bigint(19);primary_key;comment:主键id" json:"dept_id"`
	ParentId   int64     `gorm:"column:parent_id;type:bigint(19);comment:本表父级id" json:"parent_id"`
	DeptName   string    `gorm:"column:dept_name;type:varchar(255);comment:部门名称" json:"dept_name"`
	DeptStatus int       `gorm:"column:dept_status;type:tinyint(1);default:1;comment:部门状态（1:正常 2:删除）" json:"dept_status"`
	CreatedAt  time.Time `gorm:"column:created_at;type:datetime;comment:创建时间" json:"created_at"`
	UpdatedAt  time.Time `gorm:"column:updated_at;type:datetime;comment:修改时间" json:"updated_at"`
}

func (m *AdminDept) TableName() string {
	return "gdxz_admin_dept"
}
