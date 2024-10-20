package models

import "time"

type Model struct {
	CreatedAt time.Time `gorm:"type:datetime" json:"createdAt"`
	UpdatedAt time.Time `gorm:"type:datetime" json:"updatedAt"`
	DeletedAt uint64    `gorm:"type:bigint" json:"deletedAt"`
}
