package belajargolangjson

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestJSONTag(t *testing.T) {
	product := Product{
		Id:       "P1",
		Name:     "Lenovo",
		ImageURL: "http://example.com/image.png",
	}
	bytes, _ := json.Marshal(product)
	fmt.Println(string(bytes))
}

func TestJSONTagEncode(t *testing.T) {
	jsonString := `{"id":"P0001","name":"Apple Mac Book Pro","imageurl":"http://example.com/image.png"}`
	jsonBytes := []byte(jsonString)

	product := &Product{}

	json.Unmarshal(jsonBytes, product)
	fmt.Println(product)
}
