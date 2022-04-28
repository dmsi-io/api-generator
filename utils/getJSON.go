package utils

func GetJSON(filename string) (map[string]interface{}, error) {

	jsonMap, err := OpenJSONFile(filename)
	if err == nil {
		return jsonMap, nil
	}

	return ReadJSONFromStdin()
}
