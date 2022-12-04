package utils

import "fmt"

func GetString(data map[string]interface{}, key string) (string, error) {
	dataRaw, ok := data[key]
	if !ok {
		return "", fmt.Errorf("%v field is required", key)
	}

	dataString, ok := dataRaw.(string)
	if !ok {
		return "", fmt.Errorf("%v field type must be a string", key)
	}

	return dataString, nil
}

func GetFloat(data map[string]interface{}, key string) (float64, error) {
	dataRaw, ok := data[key]
	if !ok {
		return 0.0, fmt.Errorf("%v field is required", key)
	}

	dataFloat, ok := dataRaw.(float64)
	if !ok {
		return 0.0, fmt.Errorf("%v field type must be a float", key)
	}

	return dataFloat, nil
}

/* func GetFloatInterface(data interface{}, key string) (float64, error) {
	data[]
	
	dataRaw, ok := data[key]
	if !ok {
		return 0.0, fmt.Errorf("%v field is required", key)
	}

	dataFloat, ok := dataRaw.(float64)
	if !ok {
		return 0.0, fmt.Errorf("%v field type must be a float", key)
	}

	return dataFloat, nil
} */

