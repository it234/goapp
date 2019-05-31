package sys

import (
	"time"

	"goapp/internal/pkg/models/basemodel"
	"goapp/internal/pkg/models/db"

	"github.com/jinzhu/gorm"
)

// 后台用户
type Admins struct {
	basemodel.Model
	Memo     string `gorm:"column:memo;size:64;" json:"memo" form:"memo"`                                                          //备注
	UserName string `gorm:"column:user_name;size:32;unique_index:uk_admins_user_name;not null;" json:"user_name" form:"user_name"` // 用户名
	RealName string `gorm:"column:real_name;size:32;" json:"real_name" form:"real_name"`                                           // 真实姓名
	Password string `gorm:"column:password;type:char(32);not null;" json:"password" form:"password"`                               // 密码(sha1(md5(明文))加密)
	Email    string `gorm:"column:email;size:64;" json:"email" form:"email"`                                                       // 邮箱
	Phone    string `gorm:"column:phone;type:char(20);" json:"phone" form:"phone"`                                                 // 手机号
	Status   uint8  `gorm:"column:status;type:tinyint(1);not null;" json:"status" form:"status"`                          // 状态(1:正常 2:未激活 3:暂停使用)
}

// 表名
func (Admins) TableName() string {
	return TableName("admins")
}

// 添加前
func (m *Admins) BeforeCreate(scope *gorm.Scope) error {
	m.CreatedAt = time.Now()
	m.UpdatedAt = time.Now()
	return nil
}

// 更新前
func (m *Admins) BeforeUpdate(scope *gorm.Scope) error {
	m.UpdatedAt = time.Now()
	return nil
}

// 删除用户及关联数据
func (Admins) Delete(adminsids []uint64) error {
	tx := db.DB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	if err := tx.Error; err != nil {
		return err
	}
	if err := tx.Where("id in (?)", adminsids).Delete(&Admins{}).Error; err != nil {
		return err
	}
	if err := tx.Where("admins_id in (?)", adminsids).Delete(&AdminsRole{}).Error; err != nil {
		return err
	}
	return tx.Commit().Error
}
