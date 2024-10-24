package models

import (
	"orca/pkg/code"
)

type ApiSampleResponse struct {
	ID         uint64     `gorm:"type:bigint" json:"id"`
	ApiID      string     `gorm:"type:varchar(21)" json:"apiId"`
	Header     HttpHeader `json:"header"`
	Code       code.Code  `json:"code"`
	HttpStatus int        `json:"httpStatus"`
	Success    bool       `gorm:"type:boolean" json:"type"`
	Data       *any       `gorm:"json" json:"data"`
	Message    string     `gorm:"type:varchar(255)" json:"message"`
	References string     `gorm:"type:text" json:"references"`
}
