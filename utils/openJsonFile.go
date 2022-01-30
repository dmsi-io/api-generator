package utils

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

// OpenJSONFile opens and reads content from JSON file into map
func OpenJSONFile(path string) (map[string]interface{}, error) {
	jsonFile, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	defer jsonFile.Close()

	// Creating the maps for JSON
	m := map[string]interface{}{}

	byteValue, _ := ioutil.ReadAll(jsonFile)
	json.Unmarshal(byteValue, &m)

	return m, nil
}
