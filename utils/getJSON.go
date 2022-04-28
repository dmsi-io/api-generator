package utils

import "fmt"

func GetJSON(filename string) (map[string]interface{}, error) {

	jsonMap, err := OpenJSONFile(fmt.Sprintf("%s.json", filename))
	if err == nil {
		return jsonMap, nil
	}

	return ReadJSONFromStdin()
}
