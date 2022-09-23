package util

import (
	"encoding/json"
)

func ConvertToType[varType any](obj any) (varType, error) {
	c := new(varType)
	bytes, err := json.Marshal(obj)
	if err != nil {
		return *c, err
	}
	err = json.Unmarshal(bytes, c)
	if err != nil {
		return *c, err
	}
	return *c, nil
}
