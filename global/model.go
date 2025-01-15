package global

import (
	"time"

	"gorm.io/gorm"
)

type YAGAMI_MODEL struct {
	ID        uint `gorm:"primarykey" json:"ID"`
	CreateAt  time.Time
	UpdatedAt time.Time
	DeleteAt  gorm.DeletedAt `gorm:"index" json:"-"`
}
