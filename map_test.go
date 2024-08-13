package belajargolangjson

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestMapDecode(t *testing.T) {
	jsonString := `{"id": "P01","name": "Lenovo v15", "price": 20000}`
	jsonBytes := []byte(jsonString)

	var result map[string]interface{}
	json.Unmarshal(jsonBytes, &result)

	fmt.Println(result)
	fmt.Println(result["id"])
	fmt.Println(result["name"])
	fmt.Println(result["price"])
}

func TestMapEncode(t *testing.T) {
	product := map[string]interface{}{
		"id":    "P01",
		"name":  "Lenovo1",
		"price": 9000000,
	}
	bytes, _ := json.Marshal(product)
	fmt.Println(string(bytes))
}
