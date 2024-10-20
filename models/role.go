package models

type Role struct {
	RoleID      uint64 `gorm:"type:bigint" json:"roleId"`
	Label       string `gorm:"type:varchar(20)" json:"label"`
	Code        string `gorm:"type:varchar(255)" json:"code"`
	Status      bool   `gorm:"type:boolean" json:"status"`
	Description string `gorm:"type:text" json:"description"`

	Menu []*Menu `gorm:"many2many:role_menu" json:"menu"`
}

type RoleList struct {
	Total int64   `json:"total"`
	Items []*Role `json:"items"`
}

func (Role) TableName() string { return "roles" }
