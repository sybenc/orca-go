package models

import (
	"database/sql/driver"
	"errors"
)

type HttpParamType string

const (
	EnumApiRequestParamTypeQuery  HttpParamType = "Query"
	EnumApiRequestParamTypePath   HttpParamType = "Path"
	EnumApiRequestParamTypeHeader HttpParamType = "Header"
	EnumApiRequestParamTypeBody   HttpParamType = "Body"
	EnumApiRequestParamTypeCookie HttpParamType = "Cookie"
)

type ApiRequestParam struct {
	ID          uint64         `gorm:"type:bigint" json:"id"`
	ApiID       string         `gorm:"type:varchar(21)" json:"apiId"`
	Key         string         `gorm:"type:varchar(255)" json:"key"`
	Type        *HttpParamType `json:"type"`
	Description string         `gorm:"type:text" json:"description"`
	Required    bool           `gorm:"type:boolean" json:"required"`
	Example     string         `gorm:"type:text" json:"example"`
	Default     string         `gorm:"type:text" json:"default"`
}

func (hpt *HttpParamType) Value() (driver.Value, error) {
	return string(*hpt), nil
}

func (hpt *HttpParamType) Scan(value any) error {
	if value == nil {
		return nil
	}
	bytes, ok := value.([]byte)
	if !ok {
		return errors.New("failed to scan HttpParamType")
	}
	switch string(bytes) {
	case "Query":
		*hpt = EnumApiRequestParamTypeQuery
	case "Path":
		*hpt = EnumApiRequestParamTypePath
	case "Header":
		*hpt = EnumApiRequestParamTypeHeader
	case "Body":
		*hpt = EnumApiRequestParamTypeBody
	case "Cookie":
		*hpt = EnumApiRequestParamTypeCookie
	default:
		return errors.New("unknown HttpParamType value")
	}
	return nil
}
