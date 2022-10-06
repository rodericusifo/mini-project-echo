package util

import (
	"encoding/json"
)

func ConvertToStruct[T any](obj any) (T, error) {
	c := new(T)
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
