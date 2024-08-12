package belajargolangjson

import (
	"encoding/json"
	"fmt"
	"testing"
)

type Customer struct {
	FirstName  string
	MiddleName string
	LastName   string
	Age        int
	Married    bool
	Hobbies    []string
}

func TestJSONObject(t *testing.T) {
	customer := Customer{
		FirstName:  "Arya",
		MiddleName: "Rahmat",
		LastName:   "Faizin",
		Age:        18,
		Married:    false,
		Hobbies:    []string{"Gaming", "Reading", "Coding"},
	}
	bytes, _ := json.Marshal(customer)
	fmt.Println(string(bytes))
}
