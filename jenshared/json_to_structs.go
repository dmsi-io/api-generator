package jenshared

import (
	"fmt"
	"strings"

	"github.com/dave/jennifer/jen"
)

// AddStructsFromJSON appends new Structs to the provided file from provided JSON
func AddStructsFromJSON(f *jen.File, m map[string]interface{}, methodName, topLevelObject string) {
	structItemMap := CreateStructItemMapFromJSON(m, methodName, topLevelObject)
	AddStructs(f, structItemMap)
}

// CreateStructItemMapFromJSON generates StructItemMap from provided JSON
func CreateStructItemMapFromJSON(m map[string]interface{}, methodName, topLevelObject string) StructItemMap {
	response, ok := m[topLevelObject].(map[string]interface{})
	if !ok {
		fmt.Printf("Failed to parse JSON from: %s\n", topLevelObject)
		return StructItemMap{}
	}

	bodyResponseName := createBodyResponse(methodName, topLevelObject)
	rootResponseName := createRootResponse(methodName, topLevelObject)

	structItemMap := parseMap(response, bodyResponseName, make(StructItemMap))

	structItemMap[rootResponseName] = StructItems{{Name: strings.Title(topLevelObject), Type: bodyResponseName, JSONName: topLevelObject}}

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

func createRootResponse(method, topLevelObject string) string {
	return fmt.Sprintf("%sRoot%s", method, strings.Title(topLevelObject))
}

func createBodyResponse(method, topLevelObject string) string {
	return fmt.Sprintf("%sBody%s", method, strings.Title(topLevelObject))
}

func createEndpoint(method string) string {
	return fmt.Sprintf("%sEndpoint", method)
}
