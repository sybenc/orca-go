package models

import "time"

type UserPasswordHistory struct {
	ID            uint64    `gorm:"type:bigint" json:"id"`
	UserID        uint64    `gorm:"type:bigint" json:"userPasswordHistoryId"`
	Password      string    `gorm:"type:varchar(255)" json:"password"`
	LastChangedAt time.Time `gorm:"type:datetime" json:"lastChangedAt"`
	LastUsedAt    time.Time `gorm:"type:datetime" json:"lastLoginAt"`
	LastUsedIP    string    `gorm:"type:varchar(128);" json:"lastLoginIp"`
}

func (uph *UserPasswordHistory) TableName() string {
	return "user_password_history"
}
