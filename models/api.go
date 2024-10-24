package models

import (
	"database/sql/driver"
	"gorm.io/gorm"
	"orca/pkg/errors"
	"orca/pkg/utils/idutils"
)

type HttpMethod string

const (
	EnumHttpMethodGet    HttpMethod = "GET"
	EnumHttpMethodPost   HttpMethod = "POST"
	EnumHttpMethodPut    HttpMethod = "PUT"
	EnumHttpMethodDelete HttpMethod = "DELETE"
)

type Api struct {
	Model       `json:",inline"`
	ApiID       string     `gorm:"type:varchar(21)" json:"apiId"`
	Name        string     `gorm:"type:varchar(64)" json:"name"`
	EndPoint    string     `gorm:"type:varchar(255)" json:"endPoint"`
	HttpMethod  HttpMethod `json:"httpMethod"`
	Group       string     `gorm:"type:varchar(64)" json:"group"`
	Version     string     `gorm:"type:varchar(64)" json:"version"`
	Status      bool       `gorm:"type:boolean" json:"status"`
	Description string     `gorm:"type:text" json:"description"`

	RequestParams  []*ApiRequestParam   `json:"requestParams,omitempty"`
	SampleRequest  []*ApiSampleRequest  `json:"sampleRequest,omitempty"`
	SampleResponse []*ApiSampleResponse `json:"sampleResponse,omitempty"`
}

func (a *Api) TableName() string {
	return "apis"
}

func (a *Api) BeforeCreate(tx *gorm.DB) error {
	a.ApiID = idutils.Nanoid.Must()
	return nil
}

func (hm *HttpMethod) Value() (driver.Value, error) {
	return string(*hm), nil
}

func (hm *HttpMethod) Scan(value any) error {
	if value == nil {
		return nil
	}
	bytes, ok := value.([]byte)
	if !ok {
		return errors.New("failed to scan HttpMethod")
	}
	switch string(bytes) {
	case "GET":
		*hm = EnumHttpMethodGet
	case "POST":
		*hm = EnumHttpMethodPost
	case "PUT":
		*hm = EnumHttpMethodPut
	case "DELETE":
		*hm = EnumHttpMethodDelete
	default:
		return errors.New("unknown HttpMethod value")
	}
	return nil
}
