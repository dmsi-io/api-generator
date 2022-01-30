package jenshared

import (
	"fmt"
	"strings"

	"github.com/dave/jennifer/jen"
)

// AddStructsFromJSON appends new Structs to the provided file from provided JSON
func AddStructsFromJSON(f *jen.File, m map[string]interface{}) {
	f.Add(CreateStructsFromJSON(m))
}

// CreateStructsFromJSON generates structs from provided JSON
func CreateStructsFromJSON(m map[string]interface{}) *jen.Statement {
	parseMap(m, make(StructItemMap))
	return nil
}

func parseMap(aMap map[string]interface{}, items StructItemMap) (string, StructItemMap) {
	for key, val := range aMap {
		switch concreteVal := val.(type) {
		case map[string]interface{}:
			parseMap(val.(map[string]interface{}), items)
		case []interface{}:
			fmt.Println(key)
			parseArray(val.([]interface{}), items)
		default:
			fmt.Println(key, ":", inferDataType(concreteVal))
		}
	}
	return "", items
}

func parseArray(anArray []interface{}, items StructItemMap) (string, StructItemMap) {
	for i, val := range anArray {
		switch concreteVal := val.(type) {
		case map[string]interface{}:
			fmt.Println("Index:", i)
			parseMap(val.(map[string]interface{}), items)
		case []interface{}:
			fmt.Println("Index:", i)
			parseArray(val.([]interface{}), items)
		default:
			fmt.Println("Index", i, ":", inferDataType(concreteVal))
		}
	}
	return "", items
}

func inferDataType(value interface{}) string {
	dataType := fmt.Sprintf("%T", value)
	if dataType == "float64" && !strings.Contains(fmt.Sprintf("%v", value), ".") {
		dataType = "int"
	}

	return dataType
}
