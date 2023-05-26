// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"

	"gorm.io/plugin/soft_delete"
)

const TableNameMember = "greet_member"

// Member mapped from table <greet_member>
type Member struct {
	ID        int64                 `gorm:"column:id;type:bigint unsigned;primaryKey;autoIncrement:true" json:"id"` // Id
	Mobile    string                `gorm:"column:mobile;type:char(11);not null" json:"mobile"`                     // 登陆手机号
	Nickname  string                `gorm:"column:nickname;type:varchar(50);not null" json:"nickname"`              // 用户昵称
	Password  string                `gorm:"column:password;type:char(32);not null" json:"password"`                 // 密码
	DeletedAt soft_delete.DeletedAt `gorm:"column:deleted_at;type:timestamp;default:null" json:"deleted_at"`
	CreatedAt *time.Time            `gorm:"column:created_at;type:timestamp;autoCreateTime" json:"created_at"`
	UpdatedAt *time.Time            `gorm:"column:updated_at;type:timestamp;autoUpdateTime" json:"updated_at"`
}

// TableName Member's table name
func (*Member) TableName() string {
	return TableNameMember
}
