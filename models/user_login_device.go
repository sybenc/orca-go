package models

import "time"

type UserLoginDevice struct {
	ID          uint64    `gorm:"type:bigint" json:"id"`
	UserID      uint64    `gorm:"type:bigint" json:"userLoginDeviceId"`
	OS          string    `gorm:"type:varchar(64)" json:"os"`
	DeviceName  string    `gorm:"type:varchar(100)" json:"deviceName"`
	DeviceType  string    `gorm:"type:varchar(50)" json:"deviceType"`
	Browser     string    `gorm:"type:varchar(50)" json:"browser"`
	LastLoginAt time.Time `gorm:"type:datetime" json:"lastLoginAt"`
	LastLoginIP string    `gorm:"type:varchar(128);not null" json:"lastLoginIp"`
}
