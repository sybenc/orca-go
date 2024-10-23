package models

import (
	"database/sql/driver"
	"errors"
	"orca/pkg/utils/authutils"
)

type UserStatus string

const (
	EnumUserStatusActive     UserStatus = "Active"
	EnumUserStatusUnverified UserStatus = "Unverified"
	EnumUserStatusDisabled   UserStatus = "Disabled"
	EnumUserStatusDeleted    UserStatus = "Deleted"
	EnumUserStatusLocked     UserStatus = "Locked"
	EnumUserStatusCancelled  UserStatus = "Cancelled"
)

type UserAuth struct {
	UserAuthID uint64     `gorm:"type:bigint" json:"userAuthId"`
	Password   string     `gorm:"type:var(255)" json:"password"`
	Status     UserStatus `json:"status"`
}

func (ua *UserAuth) TableName() string {
	return "user_auths"
}

func (ua *UserAuth) Hash(str string) (string, error) {
	return authutils.Argon2id.Hash(str)
}

func (ua *UserAuth) Verify(str, storedHash string) bool {
	return authutils.Argon2id.Verify(str, storedHash)
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
