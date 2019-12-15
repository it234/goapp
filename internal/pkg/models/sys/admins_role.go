package sys

import (
	"time"

	"goapp/internal/pkg/models/basemodel"
	"goapp/internal/pkg/models/db"

	"github.com/jinzhu/gorm"
)

// 用户-角色
type AdminsRole struct {
	basemodel.Model
	AdminsID uint64 `gorm:"column:admins_id;unique_index:uk_admins_role_admins_id;not null;"` // 管理员ID
	RoleID   uint64 `gorm:"column:role_id;unique_index:uk_admins_role_admins_id;not null;"`   // 角色ID
}

// 表名
func (AdminsRole) TableName() string {
	return TableName("admins_role")
}

// 添加前
func (m *AdminsRole) BeforeCreate(scope *gorm.Scope) error {
	m.CreatedAt = time.Now()
	m.UpdatedAt = time.Now()
	return nil
}

// 更新前
func (m *AdminsRole) BeforeUpdate(scope *gorm.Scope) error {
	m.UpdatedAt = time.Now()
	return nil
}

// 分配用户角色
func (AdminsRole) SetRole(adminsid uint64, roleids []uint64) error {
	tx := db.DB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	if err := tx.Error; err != nil {
		tx.Rollback()
		return err
	}
	if err := tx.Where(&AdminsRole{AdminsID: adminsid}).Delete(&AdminsRole{}).Error; err != nil {
		tx.Rollback()
		return err
	}
	if len(roleids) > 0 {
		for _, rid := range roleids {
			rm := new(AdminsRole)
			rm.RoleID = rid
			rm.AdminsID = adminsid
			if err := tx.Create(rm).Error; err != nil {
				tx.Rollback()
				return err
			}
		}
	}
	return tx.Commit().Error
}
