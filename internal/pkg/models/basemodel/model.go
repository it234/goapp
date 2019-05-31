package basemodel

import (
	"time"
)

type Model struct {
	ID        uint64    `gorm:"column:id;primary_key;auto_increment;" json:"id" form:"id"`                     // 主键
	CreatedAt time.Time `gorm:"column:created_at;type:datetime;not null;" json:"created_at" form:"created_at"` // 创建时间
	UpdatedAt time.Time `gorm:"column:updated_at;type:datetime;not null;" json:"updated_at" form:"updated_at"` // 更新时间
	CreatedBy uint64    `gorm:"column:created_by;default:0;not null;" json:"created_by" form:"created_by"`     // 创建人
	UpdatedBy uint64    `gorm:"column:updated_by;default:0;not null;" json:"updated_by" form:"updated_by"`     // 更新人
}

func GetTablePrefix() string {
	return "tb_"
}
