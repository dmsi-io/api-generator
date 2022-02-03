package jenshared

import (
	"fmt"
	"strings"

	"github.com/dave/jennifer/jen"
)

// AddStructsFromJSON appends new Structs to the provided file from provided JSON
func AddStructsFromJSON(f *jen.File, m map[string]interface{}) {
	structItemMap := CreateStructItemMapFromJSON(m)
	AddStructs(f, structItemMap)
}

// CreateStructItemMapFromJSON generates StructItemMap from provided JSON
func CreateStructItemMapFromJSON(m map[string]interface{}) StructItemMap {
	methodName := m["method"].(string)
	response := m["response"].(map[string]interface{})
	bodyResponseName := createBodyResponse(methodName)
	rootResponseName := createRootResponse(methodName)

	structItemMap := parseMap(response, bodyResponseName, make(StructItemMap))

	structItemMap[rootResponseName] = StructItems{{Name: "Response", Type: bodyResponseName, JSONName: "response"}}

	return structItemMap
}

func parseMap(aMap map[string]interface{}, parent string, items StructItemMap) StructItemMap {
	for key, val := range aMap {
		title := strings.Title(key)
		switch concreteVal := val.(type) {
		case map[string]interface{}:
			items[title] = make(StructItems, 0)
			items[parent] = append(items[parent], StructItem{JSONName: key, Name: title, Type: title})
			parseMap(val.(map[string]interface{}), title, items)
		case []interface{}:
			items[parent] = append(items[parent], StructItem{JSONName: key, Name: title, Type: fmt.Sprintf("[]%s", title)})
			parseFirstIndexArray(val.([]interface{}), title, items)
		default:
			items[parent] = append(items[parent], StructItem{JSONName: key, Name: title, Type: inferDataType(concreteVal)})
		}
	}
	return items
}

func parseFirstIndexArray(anArray []interface{}, parent string, items StructItemMap) StructItemMap {
	if len(anArray) > 0 {
		i := 0
		val := anArray[i]
		switch concreteVal := val.(type) {
		case map[string]interface{}:
			parseMap(val.(map[string]interface{}), parent, items)
		case []interface{}:
			innerParent := fmt.Sprintf("Inner%s", parent)
			items[parent] = append(items[parent], StructItem{Name: innerParent, Type: fmt.Sprintf("[]%s", strings.Title(innerParent))})
			parseFirstIndexArray(val.([]interface{}), innerParent, items)
		default:
			delete(items, parent)
			for key, itemArray := range items {
				for ii, item := range itemArray {
					if strings.Title(item.Name) == parent && item.Type == fmt.Sprintf("[]%s", parent) {
						items[key][ii].Type = fmt.Sprintf("[]%s", inferDataType(concreteVal))
					}
				}
			}
		}
	}
	return items
}

func inferDataType(value interface{}) string {
	dataType := fmt.Sprintf("%T", value)
	if dataType == "float64" && !strings.Contains(fmt.Sprintf("%v", value), ".") {
		dataType = "int"
	}

	return dataType
}

func createRootResponse(method string) string {
	return fmt.Sprintf("%sRootResponse", method)
}

func createBodyResponse(method string) string {
	return fmt.Sprintf("%sBodyResponse", method)
}

func createEndpoint(method string) string {
	return fmt.Sprintf("%sEndpoint", method)
}
