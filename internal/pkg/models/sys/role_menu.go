package sys

import (
	"time"

	"github.com/it234/goapp/internal/pkg/models/basemodel"
	"github.com/it234/goapp/internal/pkg/models/db"

	"github.com/jinzhu/gorm"
)

// 角色-菜单
type RoleMenu struct {
	basemodel.Model
	RoleID uint64 `gorm:"column:role_id;unique_index:uk_role_menu_role_id;not null;"` // 角色ID
	MenuID uint64 `gorm:"column:menu_id;unique_index:uk_role_menu_role_id;not null;"` // 菜单ID
}

// 表名
func (RoleMenu) TableName() string {
	return TableName("role_menu")
}

// 添加前
func (m *RoleMenu) BeforeCreate(scope *gorm.Scope) error {
	m.CreatedAt = time.Now()
	m.UpdatedAt = time.Now()
	return nil
}

// 更新前
func (m *RoleMenu) BeforeUpdate(scope *gorm.Scope) error {
	m.UpdatedAt = time.Now()
	return nil
}

// 设置角色菜单权限
func (RoleMenu) SetRole(roleid uint64, menuids []uint64) error {
	tx := db.DB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	if err := tx.Error; err != nil {
		return err
	}
	if err := tx.Where(&RoleMenu{RoleID: roleid}).Delete(&RoleMenu{}).Error; err != nil {
		return err
	}
	if len(menuids) > 0 {
		for _, mid := range menuids {
			rm := new(RoleMenu)
			rm.RoleID = roleid
			rm.MenuID = mid
			if err := tx.Create(rm).Error; err != nil {
				return err
			}
		}
	}
	return tx.Commit().Error
}
