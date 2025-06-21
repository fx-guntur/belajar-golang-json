package belajargolangjson

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestJSONArray(t *testing.T) {
	customer := Customer{
		FirstName:  "Bobi",
		MiddleName: "Bebi",
		LastName:   "Bubi",
		Age:        10,
		Married:    false,
		Hobbies: []string{
			"ngegame",
			"ngoding",
		},
	}

	bytes, _ := json.Marshal(customer)
	fmt.Println(string(bytes))
}

func TestJSONDecode(t *testing.T) {
	jsonString := `{"FirstName":"Bobi","MiddleName":"Bebi","LastName":"Bubi","Age":10,"Married":false,"Hobbies":["ngegame","ngoding"]}`
	jsonBytes := []byte(jsonString)

	customer := &Customer{}

	err := json.Unmarshal(jsonBytes, customer)
	if err != nil {
		panic(err)
	}

	fmt.Println(customer)
	fmt.Println(customer.FirstName)
	fmt.Println(customer.Hobbies)
}

func TestJSONArrayComplexEncode(t *testing.T) {
	customer := Customer{
		FirstName:  "Bobi",
		MiddleName: "Bebi",
		LastName:   "Bubi",
		Age:        10,
		Married:    false,
		Hobbies: []string{
			"ngegame",
			"ngoding",
		},
		Address: []Address{
			{
				Street:     "Jalan belum ada",
				Country:    "Indonesia",
				PostalCode: "14045",
			},
			{
				Street:     "Jalan lagi diaspal",
				Country:    "Indonesia Utara",
				PostalCode: "14046",
			},
		},
	}

	bytes, err := json.Marshal(customer)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(bytes))

}

func TestJSONArrayyComplexDecode(t *testing.T) {
	jsonString := `{"FirstName":"Bobi","MiddleName":"Bebi","LastName":"Bubi","Age":10,"Married":false,"Hobbies":["ngegame","ngoding"],"Address":[{"Street":"Jalan belum ada","Country":"Indonesia","PostalCode":"14045"},{"Street":"Jalan lagi diaspal","Country":"Indonesia Utara","PostalCode":"14046"}]}`
	jsonBytes := []byte(jsonString)

	customer := &Customer{}
	err := json.Unmarshal(jsonBytes, customer)
	if err != nil {
		panic(err)
	}

	fmt.Println(customer)
	fmt.Println(customer.Address)
	fmt.Println(customer.Hobbies)
}

func TestOnlyJSONArrayyComplexDecode(t *testing.T) {
	jsonString := `[{"Street":"Jalan belum ada","Country":"Indonesia","PostalCode":"14045"},{"Street":"Jalan lagi diaspal","Country":"Indonesia Utara","PostalCode":"14046"}]`
	jsonBytes := []byte(jsonString)

	addresses := &[]Address{}
	err := json.Unmarshal(jsonBytes, addresses)
	if err != nil {
		panic(err)
	}

	fmt.Println(addresses)
}

func TestOnlyJSONArrayComplexEncode(t *testing.T) {

	addresses := []Address{
		{
			Street:     "Jalan belum ada",
			Country:    "Indonesia",
			PostalCode: "14045",
		},
		{
			Street:     "Jalan lagi diaspal",
			Country:    "Indonesia Utara",
			PostalCode: "14046",
		},
	}

	bytes, err := json.Marshal(addresses)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(bytes))

}
