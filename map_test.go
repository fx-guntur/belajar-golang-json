package belajargolangjson

import (
	"encoding/json"
	"fmt"
	"testing"
)

// gunakan map ketika tipe data yang diterima / dikirim tidak terprediksi sehingga bisa pakai any / interface

func TestMapDecode(t *testing.T) {
	jsonString := `{"id":"P001","name":"Acer Nitro 5","image_url":"https://google.com"}`
	jsonBytes := []byte(jsonString)

	var result map[string]interface{}
	_ = json.Unmarshal(jsonBytes, &result)

	fmt.Println(result)
	fmt.Println(result["name"])
	fmt.Println(result["image_url"])
}

func TestMapEncode(t *testing.T) {
	product := map[string]interface{}{
		"id":        "P001",
		"name":      "Acer Nitro 5",
		"image_url": "https://google.com",
	}

	bytes, _ := json.Marshal(product)
	fmt.Println(string(bytes))
}
