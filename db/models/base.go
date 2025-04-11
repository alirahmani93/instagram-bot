package models

import (
    "time"
    "gorm.io/gorm"
)

type BaseModel struct {
    ID        uint           `gorm:"primaryKey" json:"id"`
    IsActive  bool           `gorm:"default:true" json:"is_active"`
    CreatedAt time.Time      `gorm:"autoCreateTime" json:"created_at"`
    UpdatedAt time.Time      `gorm:"autoUpdateTime" json:"updated_at"`
}

// BeforeCreate برای تنظیم مقدار پیش‌فرض IsActive
func (b *BaseModel) BeforeCreate(tx *gorm.DB) (err error) {
    if !b.IsActive {
        b.IsActive = true
    }
    return
}
