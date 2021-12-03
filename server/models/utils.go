package models

import (
	"encoding/json"
	"fmt"
)

func ParseArray(x ...interface{}) ([]string, error) {
	// strArr := make([]string, len(x))
	strArr := make([]string, 0)
	fmt.Println(x...)

	for _, elem := range x {
		json, err := json.Marshal(elem)
		fmt.Println(string(json))
		if err != nil {
			return nil, err
		}
		strArr = append(strArr, string(json))
	}

	return strArr, nil
}
