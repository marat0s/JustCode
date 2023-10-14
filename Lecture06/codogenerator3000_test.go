package main

import (
	"testing"
)

func TestGenerateStruct(t *testing.T) {
	jsonStr := `{
		"name": "string",
		"age": "int"
	}`

	expected := "type Root struct {\n\tname string `json:\"name\"`\n\tage int `json:\"age\"`\n}\n\n"
	result := generateStruct(jsonStr)

	if expected != result {
		t.Errorf("Expected:%v\nGot:%v", expected, result)
	}
}
