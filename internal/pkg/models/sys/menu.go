package sys

import (
	"time"

	"goapp/internal/pkg/models/basemodel"
	"goapp/internal/pkg/models/db"

	"github.com/jinzhu/gorm"
)

// 菜单
type Menu struct {
	basemodel.Model
	Status      uint8  `gorm:"column:status;type:tinyint(1);not null;" json:"status" form:"status"`             // 状态(1:启用 2:不启用)
	Memo        string `gorm:"column:memo;size:64;" json:"memo" form:"memo"`                                    // 备注
	ParentID    uint64 `gorm:"column:parent_id;not null;" json:"parent_id" form:"parent_id"`                    // 父级ID
	URL         string `gorm:"column:url;size:72;" json:"url" form:"url"`                                       // 菜单URL
	Name        string `gorm:"column:name;size:32;not null;" json:"name" form:"name"`                           // 菜单名称
	Sequence    int    `gorm:"column:sequence;not null;" json:"sequence" form:"sequence"`                       // 排序值
	MenuType    uint8  `gorm:"column:menu_type;type:tinyint(1);not null;" json:"menu_type" form:"menu_type"`    // 菜单类型 1模块2菜单3操作
	Code        string `gorm:"column:code;size:32;not null;unique_index:uk_menu_code;" json:"code" form:"code"` // 菜单代码
	Icon        string `gorm:"column:icon;size:32;" json:"icon" form:"icon"`                                    // icon
	OperateType string `gorm:"column:operate_type;size:32;not null;" json:"operate_type" form:"operate_type"`   // 操作类型 none/add/del/view/update
}

// 表名
func (Menu) TableName() string {
	return TableName("menu")
}

// 添加前
func (m *Menu) BeforeCreate(scope *gorm.Scope) error {
	m.CreatedAt = time.Now()
	m.UpdatedAt = time.Now()
	return nil
}

// 更新前
func (m *Menu) BeforeUpdate(scope *gorm.Scope) error {
	m.UpdatedAt = time.Now()
	return nil
}

// 获取菜单有权限的操作列表
func (Menu) GetMenuButton(adminsid uint64, menuCode string, btns *[]string) (err error) {
	sql := `select operate_type from tb_sys_menu
	      where id in (
					select menu_id from tb_sys_role_menu where 
					menu_id in (select id from tb_sys_menu where parent_id in (select id from tb_sys_menu where code=?))
					and role_id in (select role_id from tb_sys_admins_role where admins_id=?)
				)`
	err = db.DB.Raw(sql, menuCode, adminsid).Pluck("operate_type", btns).Error
	return
}

// 获取管理员权限下所有菜单
func (Menu) GetMenuByAdminsid(adminsid uint64, menus *[]Menu) (err error) {
	sql := `select * from tb_sys_menu
	      where id in (
					select menu_id from tb_sys_role_menu where 
				  role_id in (select role_id from tb_sys_admins_role where admins_id=?)
				)`
	err = db.DB.Raw(sql, adminsid).Find(menus).Error
	return
}

// 删除菜单及关联数据
func (Menu) Delete(menuids []uint64) error {
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
	for _, menuid := range menuids {
		if err := deleteMenuRecurve(tx, menuid); err != nil {
			tx.Rollback()
			return err
		}
	}
	if err := tx.Where("menu_id in (?)", menuids).Delete(&RoleMenu{}).Error; err != nil {
		tx.Rollback()
		return err
	}
	if err := tx.Where("id in (?)", menuids).Delete(&Menu{}).Error; err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit().Error
}

func deleteMenuRecurve(db *gorm.DB, parentID uint64) error {
	where := &Menu{}
	where.ParentID = parentID
	var menus []Menu
	dbslect := db.Where(&where)
	if err := dbslect.Find(&menus).Error; err != nil {
		return err
	}
	for _, menu := range menus {
		if err := db.Where("menu_id = ?", menu.ID).Delete(&RoleMenu{}).Error; err != nil {
			return err
		}
		if err := deleteMenuRecurve(db, menu.ID); err != nil {
			return err
		}
	}
	if err := dbslect.Delete(&Menu{}).Error; err != nil {
		return err
	}
	return nil
}
