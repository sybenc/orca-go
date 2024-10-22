package models

import (
	"database/sql/driver"
	errors "orca/pkg/errors"
	"orca/pkg/validation"
)

type MenuType string

const (
	EnumMenuTypeMenu      MenuType = "Menu"
	EnumMenuTypeDirectory MenuType = "Directory"
	EnumMenuTypeButton    MenuType = "Button"
)

type Menu struct {
	Model `json:",inline"`

	MenuID      uint64   `gorm:"type:bigint" json:"menuId"`
	Label       string   `gorm:"type:varchar(20)" json:"label"`
	Code        string   `gorm:"type:varchar(255)" json:"code"`
	ParentID    *uint64  `gorm:"type:bigint" json:"parentId"`
	Type        MenuType `json:"type"`
	Route       *string  `gorm:"type:text" json:"route"`
	Component   *string  `gorm:"type:text" json:"component"`
	IconName    string   `gorm:"type:varchar(255)" json:"iconName"`
	Order       uint     `gorm:"type:int" json:"order"`
	KeepAlive   bool     `gorm:"type:boolean" json:"keepAlive"`
	Show        bool     `gorm:"type:boolean" json:"show"`
	Status      bool     `gorm:"type:boolean" json:"status"`
	Description string   `gorm:"type:text" json:"description"`

	Roles []*Role `gorm:"many2many:role_menu" json:"roles"`
}

type MenuList struct {
	Total int64   `json:"total"`
	Items []*Menu `json:"items"`
}

func (m *Menu) TableName() string {
	return "menu"
}

func (m *Menu) Validate() error {
	return validation.ValidateStruct(
		m,
		validation.Field(&m.Code, validation.Required, validation.Length(1, 255)),
		validation.Field(&m.Label, validation.Required, validation.Length(1, 20)),
		validation.Field(&m.Type, validation.Required, validation.In(EnumMenuTypeMenu, EnumMenuTypeDirectory, EnumMenuTypeButton)),
		validation.Field(&m.Status, validation.Required))
}

func (mt *MenuType) Value() (driver.Value, error) {
	return string(*mt), nil
}

func (mt *MenuType) Scan(value any) error {
	if value == nil {
		return nil
	}
	bytes, ok := value.([]byte)
	if !ok {
		return errors.New("failed to scan MenuType")
	}
	switch string(bytes) {
	case "Menu":
		*mt = EnumMenuTypeMenu
	case "Directory":
		*mt = EnumMenuTypeDirectory
	case "Button":
		*mt = EnumMenuTypeButton
	default:
		return errors.New("unknown MenuType value")
	}
	return nil
}
