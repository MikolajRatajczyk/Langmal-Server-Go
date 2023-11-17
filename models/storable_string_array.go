package models

import (
	"database/sql"
	"database/sql/driver"
	"encoding/json"
	"errors"
)

// enables writing and reading array of strings from SQLite DB
type StorableStringArray []string

var _ sql.Scanner = (*StorableStringArray)(nil)

func (ssa *StorableStringArray) Scan(src any) error {
	bytes, ok := src.([]byte)
	if !ok {
		return errors.New("type assertion failed: it's not []byte")
	}
	return json.Unmarshal(bytes, ssa)
}

var _ driver.Value = StorableStringArray{}

func (ssa StorableStringArray) Value() (driver.Value, error) {
	return json.Marshal(ssa)
}
