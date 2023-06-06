package model

import (
	"database/sql/driver"
	"errors"
	"fmt"
	"gorm.io/gorm"
	"reflect"
	"strings"
	"time"
)

type Model struct {
	CreatedAt LocalTime `gorm:"column:created_at;type:datetime;comment:创建时间" json:"created_at"`
	UpdatedAt LocalTime `gorm:"column:updated_at;type:datetime;comment:修改时间" json:"updated_at"`
}

// LocalTime 自定义数据类型
type LocalTime time.Time

func (t *LocalTime) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		return nil
	}
	var err error
	// 前端接收的时间字符串
	str := string(data)
	// 去除接收的str收尾多余的"
	timeStr := strings.Trim(str, "\"")
	t1, err := time.Parse("2006-01-02 15:04:05", timeStr)
	*t = LocalTime(t1)
	return err
}

func (t LocalTime) MarshalJSON() ([]byte, error) {
	formatted := fmt.Sprintf("\"%v\"", time.Time(t).Format("2006-01-02 15:04:05"))
	return []byte(formatted), nil
}

func (t LocalTime) Value() (driver.Value, error) {
	// MyTime 转换成 time.Time 类型
	tTime := time.Time(t)
	return tTime.Format("2006-01-02 15:04:05"), nil
}

func (t *LocalTime) Scan(v interface{}) error {
	switch vt := v.(type) {
	case time.Time:
		// 字符串转成 time.Time 类型
		*t = LocalTime(vt)
	default:
		return errors.New("类型处理错误")
	}
	return nil
}

func (t *LocalTime) String() string {
	return fmt.Sprintf("hhh:%s", time.Time(*t).String())
}

// BeforeCreate 注册 BeforeCreate 回调函数
func (m *Model) BeforeCreate(tx *gorm.DB) (err error) {
	// 动态访问 YourModel 结构体本身
	model := tx.Statement.Dest

	// 设置创建时间
	layout := "2006-01-02 15:04:05"
	strTime := "2019-08-09 11:35:52"
	parsedTime, err := time.Parse(layout, strTime)
	if err != nil {
		return err
	}

	// 使用反射设置创建时间字段
	createdAtField := reflect.ValueOf(model).Elem().FieldByName("CreatedAt")
	if createdAtField.CanSet() {
		createdAtField.Set(reflect.ValueOf(parsedTime))
	}

	return nil
}
