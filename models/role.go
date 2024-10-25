package models

import (
	"gorm.io/gorm"
	"orca/pkg/utils/idutils"
)

type Role struct {
	Model       `json:",inline"`
	RoleID      string `gorm:"type:varchar(21)" json:"roleId"`
	Label       string `gorm:"type:varchar(20)" json:"label"`
	Code        string `gorm:"type:varchar(255)" json:"code"`
	Status      bool   `gorm:"type:boolean" json:"status"`
	Description string `gorm:"type:text" json:"description"`

	Menu []*Menu `gorm:"many2many:role_menu" json:"menu,omitempty"`
	Api  []*Api  `gorm:"many2many:role_api.sql" json:"api,omitempty"`
}

type RoleList struct {
	Total int64   `json:"total"`
	Items []*Role `json:"items"`
}

func (r *Role) TableName() string { return "roles" }

func (r *Role) BeforeCreate(tx *gorm.DB) error {
	r.RoleID = idutils.Nanoid.Must()
	return nil
}
