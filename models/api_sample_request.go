package models

import (
	"database/sql/driver"
	"encoding/json"
	"net/http"
	"orca/pkg/errors"
)

type HttpHeader http.Header

type ApiSampleRequest struct {
	ID     uint64            `gorm:"type:bigint" json:"id"`
	ApiID  string            `gorm:"type:varchar(21)" json:"apiId"`
	Header HttpHeader        `json:"header"`
	Params map[string]string `gorm:"type:json" json:"params"`
}

func (h *HttpHeader) Value() (driver.Value, error) {
	return json.Marshal(h)
}

func (h *HttpHeader) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return errors.New("type assertion to []byte failed")
	}

	return json.Unmarshal(bytes, h)
}
