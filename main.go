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

}

