// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

type SourceArrayEnum string

const (
	SourceArrayEnumInternal SourceArrayEnum = "INTERNAL"
	SourceArrayEnumExternal SourceArrayEnum = "EXTERNAL"
)

func (e SourceArrayEnum) ToPointer() *SourceArrayEnum {
	return &e
}

func (e *SourceArrayEnum) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "INTERNAL":
		fallthrough
	case "EXTERNAL":
		*e = SourceArrayEnum(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SourceArrayEnum: %v", v)
	}
}