package belajargolangjson

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestJSONArray(t *testing.T) {

	customer := Customer{
		FirstName:  "Nur",
		MiddleName: "Dwi",
		LastName:   "Saputra",
		Age:        18,
		Married:    false,
		Hobbies:    []string{"Gaming", "Reading", "Coding"},
	}

	bytes, _ := json.Marshal(customer)
	fmt.Println(string(bytes))

}

func TestJSONArrayDecode(t *testing.T) {
	jsonString := `{"Firstname":"Arya","MiddleName":"Rahmat","LastName":"Faizin","Age":18,"Married":true,"Hobbies":["Gaming","Reading","Coding"]}`
	jsonBytes := []byte(jsonString)

	customer := Customer{}
	err := json.Unmarshal(jsonBytes, &customer)
	if err != nil {
		panic(err)
	}
	fmt.Println(customer)
	fmt.Println(customer.FirstName)
	fmt.Println(customer.MiddleName)
	fmt.Println(customer.LastName)
	fmt.Println(customer.Age)
	fmt.Println(customer.Married)
	fmt.Println(customer.Hobbies)
}

func TestJSONArrayComplex(t *testing.T) {
	customer := Customer{
		FirstName: "Dwi",
		Address: []Address{
			{
				Street:     "Jalan",
				Country:    "Singapore",
				PostalCode: "999",
			},
			{
				Street:     "jilin",
				Country:    "Hongkong",
				PostalCode: "833729",
			},
		},
	}
	bytes, _ := json.Marshal(customer)
	fmt.Println(string(bytes))
}

func TestJSONArrayComplexDecode(t *testing.T) {

	jsonString := `{"Firstname":"Arya","MiddleName":"Rahmat","LastName":"Faizin","Age":18,"Married":true,"Hobbies":["Gaming","Reading","Coding"]}`
	jsonBytes := []byte(jsonString)

	customer := Customer{}
	err := json.Unmarshal(jsonBytes, &customer)
	if err != nil {
		panic(err)
	}
	fmt.Println(customer)
	fmt.Println(customer.FirstName)
	fmt.Println(customer.Address)
}
