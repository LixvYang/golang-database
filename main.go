package main

import (
	"fmt"
	"os"
	"encoding/json"
)

type User struct {
	Name string 
	Age json.Number
	Contact string 
	Company string
	Address Address
}

type Address struct {
	City string
	State string
	Country string
	Pincode json.Number
}

func main() {
	dir := "./"
	
	dir, err := New(dir, nil)
	if err != nil {
		fmt.Println("Error:", err)
	}

	empolyees := []User{
		
	}

	for _, value := range empolyees {
		db.Write("users", values.Name, User{
			Name:value.Name,
			Age: value.Age,
			Contact: value.Contact,
			Company: value.Company,
			Address: Address,
		})
	}

	records, err := db.ReadAll("users")
	if err != nil {
		fmt.Println("Error", err)
	}
	fmt.Println(records)
}

