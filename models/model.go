package models

type Model struct {
	CreatedAt uint64 `gorm:"type:timestamp" json:"createdAt"`
	UpdatedAt uint64 `gorm:"type:timestamp" json:"updatedAt"`
	DeletedAt uint64 `gorm:"type:bigint" json:"deletedAt"`
}
