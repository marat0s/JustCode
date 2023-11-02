package main

import (
	"encoding/json"
	"fmt"
)

func generateStruct(jsonStr string) string {
	//конвертируем json в map
	var jsonData map[string]interface{} //
	err := json.Unmarshal([]byte(jsonStr), &jsonData)
	if err != nil {
		return ""
	}

	return generateGoStruct(jsonData, "Person")
}

func generateGoStruct(jsonData map[string]interface{}, structName string) string {
	code := fmt.Sprintf(structName) // "type Person struct {\n"

	for key, value := range jsonData {
		fieldType, _ := value.(string)

		code += fmt.Sprintf("\t%v %v `json:\"%v\"`\n", key, fieldType, key) //"\tname string `json:\"name\"`\n"
	}
	code += "}\n\n" //закрываем структуру

	return code
}

func main() {
	jsonStr := `{
		"name": "string",
		"age": "int"
	}`

	fmt.Println(generateStruct(jsonStr))
}
