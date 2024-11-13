package gorm

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"strings"
)

type StringArray []string

func (s StringArray) Value() (value driver.Value, err error) {
	value = "[]"
	if s == nil {
		return
	}

	value, err = json.Marshal(s)
	return
}

func (s *StringArray) Scan(value any) (err error) {
	bytes, ok := value.([]byte)
	if !ok {
		err = fmt.Errorf("failed to scan common gorm string array, value: %v", value)
		return
	}

	if len(bytes) > 0 {
		return json.Unmarshal(bytes, &s)
	}
	return
}

func (s *StringArray) RemoveEmptyItem() {
	var filtered []string
	for _, item := range *s {
		trimmedItem := strings.TrimSpace(item)
		if len(trimmedItem) != 0 {
			filtered = append(filtered, trimmedItem)
		}
	}
	*s = filtered
}
