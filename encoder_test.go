package belajargolangjson

import (
	"encoding/json"
	"fmt"
	"os"
	"testing"
)

func TestEncoder(t *testing.T) {
	writter, _ := os.Create("CustomerOut.json")
	encoder := json.NewEncoder(writter)

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

	_ = encoder.Encode(customer)

	fmt.Println(customer)

}
