package belajargolangjson

import (
	"encoding/json"
	"fmt"
	"testing"
)

type Customer struct {
	FirstName string
	LastName  string
	Age       uint
	Married   bool
}

func TestJSONObject(t *testing.T) {
	customer := Customer{
		FirstName: "Bobi",
		LastName:  "Bubi",
		Age:       10,
		Married:   false,
	}

	bytes, err := json.Marshal(customer)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(bytes))
}
