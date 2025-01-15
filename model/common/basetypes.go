package common

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
)

type JSONMap map[string]interface{}

func (m JSONMap) Value() (driver.Value, error) {
	if m == nil {
		return nil, nil
	}
	return json.Marshal(m)
}

func (m JSONMap) Scan(value interface{}) error {
	if value == nil {
		m = make(map[string]interface{})
		return nil
	}
	var err error
	switch v := value.(type) {
	case []byte:
		err = json.Unmarshal(v, &m)
	case string:
		err = json.Unmarshal([]byte(v), &m)
	default:
		err = errors.New("basetypes.JSONMap.Scan: invalid value type")
	}
	return err
}

type TreeNode[T any] interface {
	GetChildren() []T
	SetChildren(children T)
	GetID() int
	GetParentID() int
}
