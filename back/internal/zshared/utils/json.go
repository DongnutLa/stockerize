package utils

import (
	"encoding/json"
	"fmt"
)

func StructToJsonString(obj any) string {
	jsonObj := obj
	b, err := json.Marshal(jsonObj)

	if err != nil {
		return ""
	}
	jsonData := string(b)
	return jsonData
}

func JsonToStruct[T any](jsonStr string) T {
	var data T
	var dataAny any

	err := json.Unmarshal([]byte(jsonStr), &data)
	err = json.Unmarshal([]byte(jsonStr), &dataAny)

	if err != nil {
		return data
	}

	return data
}

func EventDataToStruct[T any](data interface{}) *T {
	var result T
	mapData, ok := data.(map[string]interface{})
	if ok {
		jsonData, err := json.Marshal(mapData)
		if err != nil {
			fmt.Println("error marshalling map to json:", err)
			return nil
		}

		err = json.Unmarshal(jsonData, &result)
		if err != nil {
			fmt.Println("error unmarshalling json to struct:", err)
			return nil
		}

		return &result
	} else {
		switch v := data.(type) {
		case T:
			return &v
		case *T:
			return v
		default:
			fmt.Printf("<ASSERTION ERROR> Invalid type %T\n", data)
			return nil
		}
	}
}
