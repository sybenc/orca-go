package models

import (
	"database/sql/driver"
	"orca/pkg/errors"
	"time"
)

type UserGender string

const (
	EnumUserGenderFemale UserGender = "Female"
	EnumUserGenderMale   UserGender = "Male"
	EnumUserGenderOther  UserGender = "Other"
)

type UserProfile struct {
	UserProfileID uint64     `json:"userProfileId"`
	Email         string     `gorm:"type:varchar(64);not null" json:"email"`
	Phone         string     `gorm:"type:varchar(20)" json:"phone"`
	FirstName     string     `gorm:"type:varchar(20)" json:"firstName"`
	LastName      string     `gorm:"type:varchar(20)" json:"lastName"`
	NickName      string     `gorm:"type:varchar(20)" json:"nickName"`
	Gender        UserGender `json:"gender"`
	Country       string     `gorm:"type:varchar(100)" json:"country"`
	Province      string     `gorm:"type:varchar(100)" json:"province"`
	City          string     `gorm:"type:varchar(100)" json:"city"`
	Address       string     `gorm:"type:varchar(255)" json:"address"`
	ZipCode       string     `gorm:"type:varchar(10)" json:"zipCode"`
	Bio           string     `gorm:"type:varchar(255)" json:"bio"`
	Website       string     `gorm:"type:varchar(255)" json:"website"`
	Avatar        string     `gorm:"type:text" json:"avatar"`
	DateOfBirth   time.Time  `json:"dateOfBirth"`
	LastLoginAt   time.Time  `gorm:"type:datetime" json:"lastLoginAt"`
	LastLoginIP   string     `gorm:"type:varchar(128);not null" json:"lastLoginIp"`
}

func (u *UserProfile) TableName() string {
	return "user_profile"
}

func (ug *UserGender) Value() (driver.Value, error) {
	return string(*ug), nil
}

func (ug *UserGender) Scan(value any) error {
	if value == nil {
		return nil
	}
	bytes, ok := value.([]byte)
	if !ok {
		return errors.New("failed to scan UserStatus")
	}
	switch string(bytes) {
	case "Female":
		*ug = EnumUserGenderFemale
	case "Male":
		*ug = EnumUserGenderMale
	case "Other":
		*ug = EnumUserGenderOther
	default:
		return errors.New("unknown UserStatus value")
	}
	return nil
}
