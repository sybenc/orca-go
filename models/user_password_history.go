package models

import "time"

type UserPasswordHistory struct {
	ID                    uint64    `gorm:"type:bigint" json:"id"`
	UserPasswordHistoryID uint64    `gorm:"type:bigint" json:"userPasswordHistoryId"`
	Password              string    `gorm:"type:varchar(255)" json:"password"`
	ChangedAt             time.Time `gorm:"type:datetime" json:"changedAt"`
	LastUsedAt            time.Time `gorm:"type:datetime" json:"lastLoginAt"`
	LastUsedIP            string    `gorm:"type:varchar(128);not null" json:"lastLoginIp"`
}

func (uph *UserPasswordHistory) TableName() string {
	return "user_password_history"
}
