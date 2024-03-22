package gojson

import (
	"encoding/json"
	"fmt"
	"testing"
)

type Customer struct {
	Firstname string `json:"firstname"`
	Middlename string `json:"middlename"`
	Lastname string `json:"lastname"`
	Hobbies []string `json:"hobbies"`
	Address []Address `json:"address"`
}

type Address struct {
	Street string `json:"street"`
	Country string `json:"country"`
	Postalcode string `json:"postal_code"`
}

func TestEncodeJson(t *testing.T)  {
	customer := Customer{
		Firstname: "Hanafi",
		Middlename: "Adhi",
		Lastname: "Prasetyo",
		Hobbies: []string{"Gaming","Traveling","Coding"},
		Address: []Address{
			{
				Street: "jalan Haji intan",
				Country: "Indonesia",
				Postalcode: "332603",
			},
			{
				Street: "jalan Haji Hanafi",
				Country: "Indonesia",
				Postalcode: "332603",
			},
		},
	}

	byte, _ := json.Marshal(customer)
	fmt.Println(string(byte))
}

func TestDecodeJson(t *testing.T){
	jsonRequest := `{"Firstname":"Hanafi","Middlename":"Adhi","Lastname":"Prasetyo","Hobbies":["Gaming","Traveling","Coding"],"Address":[{"Street":"jalan Haji intan","Country":"Indonesia","Postal_code":"332603"},{"Street":"jalan Haji Hanafi","Country":"Indonesia","Postal_code":"332603"}]}`

	jsonBytes := []byte(jsonRequest)

	customer := &Customer{}
	json.Unmarshal(jsonBytes,customer)
	fmt.Println(customer)
}

func TestDecodeJsonArray(t *testing.T)  {
	jsonArray := `[{"Street":"jalan Haji intan","Country":"Indonesia","Postal_code":"332603"},{"Street":"jalan Haji Hanafi","Country":"Indonesia","Postal_code":"332603"}]`
	jsonByte := []byte(jsonArray)

	address := &[]Address{}
	json.Unmarshal(jsonByte,address)
	fmt.Println(address)
}

func TestJsonDynamic(t *testing.T) {
	jsonString := `{"street":"jalan Haji intan","country":"Indonesia","postal_code":"332603"}`
	jsonByte := []byte(jsonString)
	var result map[string]interface{}
	json.Unmarshal(jsonByte,&result)

	fmt.Println(result)
	fmt.Println(result["street"])
	fmt.Println(result["country"])
	fmt.Println(result["postal_code"])

}
