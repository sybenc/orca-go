package models

import (
	"gorm.io/gorm"
	"orca/pkg/utils/idutils"
)

type User struct {
	Model `json:",inline"`

	UserID   uint64 `gorm:"type:bigint" json:"userId"`
	Username string `gorm:"type:varchar(20);not null" json:"username"`

	UserProfile         *UserProfile           `gorm:"foreignKey:UserProfileID;references:UserID" json:"userProfile,omitempty"`
	UserAuth            *UserAuth              `gorm:"foreignKey:UserAuthID;references:UserID" json:"userAuth,omitempty"`
	UserPasswordHistory []*UserPasswordHistory `gorm:"foreignKey:UserPasswordHistoryID;references:UserID" json:"userPasswordHistory,omitempty"`
	UserLoginDevices    []*UserLoginDevice     `gorm:"foreignKey:UserLoginDeviceID;references:UserID" json:"userLoginDevices,omitempty"`
}

type UserList struct {
	Total int64   `json:"total"`
	Items []*User `json:"items"`
}

func (u *User) TableName() string {
	return "users"
}

func (u *User) BeforeCreate(tx *gorm.DB) error {
	var err error
	u.UserID, err = idutils.Sonyflake.NextID()
	if err != nil {
		return err
	}
	return nil
}
