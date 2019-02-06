package utils

import (
	"encoding/json"
)

// InterfaceToJSONString - converts to a JSON string
func InterfaceToJSONString(data interface{}) (*string, error) {
	dataBytes, err := json.MarshalIndent(data, "", " ")
	if err != nil {
		return nil, err
	}
	dataJSON := string(dataBytes)
	return &dataJSON, nil
}
