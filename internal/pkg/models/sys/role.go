package sys

import (
	"time"

	"goapp/internal/pkg/models/basemodel"
	"goapp/internal/pkg/models/db"

	"github.com/jinzhu/gorm"
)

// 角色
type Role struct {
	basemodel.Model
	Memo     string `gorm:"column:memo;size:64;" json:"memo" form:"memo"`                 // 备注
	Name     string `gorm:"column:name;size:32;not null;" json:"name" form:"name"`        // 名称
	Sequence int    `gorm:"column:sequence;not null;" json:"sequence" form:"sequence"`    // 排序值
	ParentID uint64 `gorm:"column:parent_id;not null;" json:"parent_id" form:"parent_id"` // 父级ID
}

// 表名
func (Role) TableName() string {
	return TableName("role")
}

// 添加前
func (m *Role) BeforeCreate(scope *gorm.Scope) error {
	m.CreatedAt = time.Now()
	m.UpdatedAt = time.Now()
	return nil
}

// 更新前
func (m *Role) BeforeUpdate(scope *gorm.Scope) error {
	m.UpdatedAt = time.Now()
	return nil
}

// 删除角色及关联数据
func (Role) Delete(roleids []uint64) error {
	tx := db.DB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	if err := tx.Error; err != nil {
		return err
	}
	if err := tx.Where("id in (?)", roleids).Delete(&Role{}).Error; err != nil {
		return err
	}
	if err := tx.Where("role_id in (?)", roleids).Delete(&RoleMenu{}).Error; err != nil {
		return err
	}
	return tx.Commit().Error
}
