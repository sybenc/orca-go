package models

import (
	"database/sql/driver"
	"errors"
)

type MenuType string

const (
	Menu_     MenuType = "Menu"
	Directory MenuType = "Directory"
	Button    MenuType = "Button"
)

type Menu struct {
	Model `json:",inline"`

	MenuID      uint64   `gorm:"type:bigint" json:"menuId"`
	Label       string   `gorm:"type:varchar(20)" json:"label"`
	Code        string   `gorm:"type:varchar(255)" json:"code"`
	ParentID    uint64   `gorm:"type:bigint" json:"parentId"`
	Type        MenuType `json:"type"`
	Route       string   `gorm:"type:text" json:"route"`
	Component   string   `gorm:"type:text" json:"component"`
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
	Items []*User `json:"items"`
}

func (mt *MenuType) Value() (driver.Value, error) {
	return string(*mt), nil
}

func (mt *MenuType) Scan(value any) error {
	if value == nil {
		return nil
	}
	val, ok := value.(string)
	if !ok {
		return errors.New("failed to scan UserStatus")
	}
	switch val {
	case "Menu":
		*mt = Menu_
	case "Directory":
		*mt = Directory
	case "Button":
		*mt = Button
	default:
		return errors.New("unknown UserStatus value")
	}
	return nil
}
