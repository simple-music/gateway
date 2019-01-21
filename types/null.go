package types

import (
	"encoding/json"
)

var (
	nullBytes = []byte("null")
)

type NullString struct {
	Valid  bool
	String string
}

func (v *NullString) MarshalJSON() ([]byte, error) {
	if v.Valid {
		return json.Marshal(v.String)
	}
	return nullBytes, nil
}

func (v *NullString) UnmarshalJSON(data []byte) error {
	err := json.Unmarshal(data, &v.String)
	v.Valid = err == nil
	return nil
}
