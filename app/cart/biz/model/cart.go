package model

import "gorm.io/gorm"

type Cart struct {
	gorm.Model
	UserID    uint32 `gorm:"type:int(11);not null;index:idx_user_id"`
	ProductID uint32 `gorm:"type:int(11);not null;"`
	Quantity  uint32 `gorm:"type:int(11);not null;"`
}

func (Cart) TableName() string {
	return "cart"
}
