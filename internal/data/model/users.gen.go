// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"

	"gorm.io/gorm"
)

const TableNameUser = "users"

// User 用户注册表
type User struct {
	ID          int32          `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	Username    string         `gorm:"column:username;not null;comment:用户名" json:"username"`              // 用户名
	Password    string         `gorm:"column:password;not null;comment:密码" json:"password"`               // 密码
	AvatarURL   string         `gorm:"column:avatar_url;not null;comment:头像url" json:"avatar_url"`        // 头像url
	Nickname    string         `gorm:"column:nickname;not null;comment:昵称" json:"nickname"`               // 昵称
	Gender      int32          `gorm:"column:gender;not null;comment:性别：0未知，1男，2女" json:"gender"`         // 性别：0未知，1男，2女
	Signature   string         `gorm:"column:signature;not null;comment:个性签名" json:"signature"`           // 个性签名
	Status      int32          `gorm:"column:status;not null;comment:状态: 1正常，2封禁" json:"status"`          // 状态: 1正常，2封禁
	LastLoginAt *time.Time     `gorm:"column:last_login_at;comment:最后上线时间" json:"last_login_at"`          // 最后上线时间
	LastLoginIP string         `gorm:"column:last_login_ip;not null;comment:最后上线ip" json:"last_login_ip"` // 最后上线ip
	CreatedAt   *time.Time     `gorm:"column:created_at;comment:创建时间" json:"created_at"`                  // 创建时间
	UpdatedAt   *time.Time     `gorm:"column:updated_at;comment:更新时间" json:"updated_at"`                  // 更新时间
	DeletedAt   gorm.DeletedAt `gorm:"column:deleted_at;comment:删除时间" json:"deleted_at"`                  // 删除时间
}

// TableName User's table name
func (*User) TableName() string {
	return TableNameUser
}
