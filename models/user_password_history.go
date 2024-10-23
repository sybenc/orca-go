package models

import "time"

type UserPasswordHistory struct {
	UserPasswordHistoryID uint64    `gorm:"type:bigint" json:"userPasswordHistoryId"`
	Password              string    `gorm:"type:varchar(255)" json:"password"`
	ChangedAt             uint64    `gorm:"type:datetime" json:"changedAt"`
	LastLoginAt           time.Time `gorm:"type:datetime" json:"lastLoginAt"`
	LastLoginIP           string    `gorm:"type:varchar(128);not null" json:"lastLoginIp"`
}

func (uph *UserPasswordHistory) TableName() string {
	return "user_password_history"
}
