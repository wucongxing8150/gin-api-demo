package model

import (
	"time"
)

// AdminPost 后台-岗位表
type AdminPost struct {
	PostId     int64     `gorm:"column:post_id;type:bigint(19);primary_key;comment:主键id" json:"post_id"`
	PostName   string    `gorm:"column:post_name;type:varchar(255);comment:部门名称" json:"post_name"`
	PostStatus int       `gorm:"column:post_status;type:tinyint(1);default:1;comment:状态（1:正常 2:删除）" json:"post_status"`
	CreatedAt  time.Time `gorm:"column:created_at;type:datetime;comment:创建时间" json:"created_at"`
	UpdatedAt  time.Time `gorm:"column:updated_at;type:datetime;comment:修改时间" json:"updated_at"`
}

func (m *AdminPost) TableName() string {
	return "gdxz_admin_post"
}
