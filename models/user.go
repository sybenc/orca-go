package models

import (
	"database/sql/driver"
	"errors"
	"time"
)

type UserStatus string
type UserGender string

const (
	EnumUserStatusActive     UserStatus = "Active"
	EnumUserStatusUnverified UserStatus = "Unverified"
	EnumUserStatusDisabled   UserStatus = "Disabled"
	EnumUserStatusDeleted    UserStatus = "Deleted"
	EnumUserStatusLocked     UserStatus = "Locked"
	EnumUserStatusCancelled  UserStatus = "Cancelled"
)
const (
	EnumUserGenderFemale UserGender = "Female"
	EnumUserGenderMale   UserGender = "Male"
	EnumUserGenderOther  UserGender = "Other"
)

type User struct {
	Model `json:",inline"`

	UserID      uint64     `gorm:"type:bigint" json:"userId"`
	LastLoginAt uint64     `gorm:"type:timestamp" json:"lastLoginAt"`
	LastLoginIp string     `gorm:"type:varchar(128);not null" json:"lastLoginIp"`
	Username    string     `gorm:"type:varchar(20);not null" json:"username"`
	Email       string     `gorm:"type:varchar(64);not null" json:"email"`
	Phone       string     `gorm:"type:varchar(20)" json:"phone"`
	Status      UserStatus `json:"status"`
	FirstName   string     `gorm:"type:varchar(20)" json:"firstName"`
	LastName    string     `gorm:"type:varchar(20)" json:"lastName"`
	NickName    string     `gorm:"type:varchar(20)" json:"nickName"`
	Gender      UserGender `json:"gender"`
	Country     string     `gorm:"type:varchar(100)" json:"country"`
	Province    string     `gorm:"type:varchar(100)" json:"province"`
	City        string     `gorm:"type:varchar(100)" json:"city"`
	Address     string     `gorm:"type:varchar(255)" json:"address"`
	ZipCode     string     `gorm:"type:varchar(10)" json:"zipCode"`
	Bio         string     `gorm:"type:varchar(255)" json:"bio"`
	Website     string     `gorm:"type:varchar(255)" json:"website"`
	Avatar      string     `gorm:"type:text" json:"avatar"`
	DateOfBirth time.Time  `json:"dateOfBirth"`
}

type UserList struct {
	Total int64   `json:"total"`
	Items []*User `json:"items"`
}

func (u *User) TableName() string {
	return "users"
}

func (us *UserStatus) Value() (driver.Value, error) {
	return string(*us), nil
}

func (us *UserStatus) Scan(value any) error {
	if value == nil {
		return nil
	}
	val, ok := value.(string)
	if !ok {
		return errors.New("failed to scan UserStatus")
	}
	switch val {
	case "Active":
		*us = EnumUserStatusActive
	case "Unverified":
		*us = EnumUserStatusUnverified
	case "Disabled":
		*us = EnumUserStatusDisabled
	case "Deleted":
		*us = EnumUserStatusDeleted
	case "Locked":
		*us = EnumUserStatusLocked
	case "Cancelled":
		*us = EnumUserStatusCancelled
	default:
		return errors.New("unknown UserStatus value")
	}
	return nil
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
