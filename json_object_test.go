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
	Age        uint
	Married    bool
	Hobbies    []string
	Address    []Address
}

type Address struct {
	Street     string
	Country    string
	PostalCode string
}

func TestJSONObject(t *testing.T) {
	customer := Customer{
		FirstName:  "Bobi",
		MiddleName: "Bebi",
		LastName:   "Bubi",
		Age:        10,
		Married:    false,
	}

	bytes, err := json.Marshal(customer)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(bytes))
}
