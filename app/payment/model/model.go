package model

import "gorm.io/plugin/soft_delete"

type Model struct {
	ID        int64 `gorm:"column:id;primaryKey;autoIncrement:true;comment:主键"`
	CreatedAt int64 `gorm:"column:created_at;autoCreateTime:milli"`
	UpdatedAt int64 `gorm:"column:updated_at;autoUpdateTime:milli"`

	DeletedAt soft_delete.DeletedAt `gorm:"index;comment:逻辑删除标记"`
}
