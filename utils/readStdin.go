package utils

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

func ReadJSONFromStdin() (map[string]interface{}, error) {
	info, err := os.Stdin.Stat()
	if err != nil {
		return nil, err
	}

	if info.Mode()&os.ModeCharDevice != 0 {
		fmt.Println("The command is intended to work with pipes.")
		fmt.Println("Usage: pbpaste | api-generator")
		return nil, err
	}

	reader := bufio.NewReader(os.Stdin)

	// Creating the maps for JSON
	m := map[string]interface{}{}

	byteValue, _ := ioutil.ReadAll(reader)
	json.Unmarshal(byteValue, &m)

	return m, nil
}
