package model

import (
	"time"
)

// AdminUser 后台-用户表
type AdminUser struct {
	UserId    int64     `gorm:"column:user_id;type:bigint(19);primary_key;comment:主键id" json:"user_id"`
	UserName  string    `gorm:"column:user_name;type:varchar(64);comment:用户名" json:"user_name"`
	Password  string    `gorm:"column:password;type:varchar(128);comment:密码" json:"password"`
	Salt      string    `gorm:"column:salt;type:varchar(255);comment:密码掩码" json:"salt"`
	Status    int       `gorm:"column:status;type:tinyint(1);default:2;comment:状态（1:正常 2:审核中 3:删除）" json:"status"`
	NickName  string    `gorm:"column:nick_name;type:varchar(255);comment:昵称" json:"nick_name"`
	Phone     string    `gorm:"column:phone;type:varchar(11);comment:手机号" json:"phone"`
	Avatar    string    `gorm:"column:avatar;type:varchar(255);comment:头像" json:"avatar"`
	Sex       int       `gorm:"column:sex;type:tinyint(1);comment:性别（1:男 2:女）" json:"sex"`
	Email     string    `gorm:"column:email;type:varchar(100);comment:邮箱" json:"email"`
	RoleId    int64     `gorm:"column:role_id;type:bigint(19);comment:角色表" json:"role_id"`
	DeptId    int64     `gorm:"column:dept_id;type:bigint(19);comment:部门id" json:"dept_id"`
	PostId    int64     `gorm:"column:post_id;type:bigint(19);comment:岗位id" json:"post_id"`
	CreateBy  int64     `gorm:"column:create_by;type:bigint(19);comment:创建者id（用户表id）" json:"create_by"`
	UpdateBy  int64     `gorm:"column:update_by;type:bigint(19);comment:更新者id（用户表id）" json:"update_by"`
	CreatedAt time.Time `gorm:"column:created_at;type:datetime;comment:创建时间" json:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at;type:datetime;comment:修改时间" json:"updated_at"`
}

func (m *AdminUser) TableName() string {
	return "gdxz_admin_user"
}
