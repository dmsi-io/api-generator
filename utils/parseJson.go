package utils

import (
	"fmt"
	"strings"
)

// ParseJSON parses a json object into usable form
func ParseJSON() {
	m, err := OpenJSONFile("branches.json")

	if err != nil {
		panic(err)
	}
	parseMap(m)
}

func parseMap(aMap map[string]interface{}) {
	for key, val := range aMap {
		switch concreteVal := val.(type) {
		case map[string]interface{}:
			fmt.Println(key)
			parseMap(val.(map[string]interface{}))
		case []interface{}:
			fmt.Println(key)
			parseArray(val.([]interface{}))
		default:
			printDataWithType(key, concreteVal)
		}
	}
}

func parseArray(anArray []interface{}) {
	for i, val := range anArray {
		switch concreteVal := val.(type) {
		case map[string]interface{}:
			fmt.Println("Index:", i)
			parseMap(val.(map[string]interface{}))
		case []interface{}:
			fmt.Println("Index:", i)
			parseArray(val.([]interface{}))
		default:
			printDataWithType(fmt.Sprintf("Index %d", i), concreteVal)
		}
	}
}

func printDataWithType(index string, value interface{}) {
	dataType := inferDataType(value)

	fmt.Printf("%s : %v (%s)\n", index, value, dataType)
}

func inferDataType(value interface{}) string {
	dataType := fmt.Sprintf("%T", value)
	if dataType == "float64" && !strings.Contains(fmt.Sprintf("%v", value), ".") {
		dataType = "int"
	}

	return dataType
}
