package main

import (
	"fmt"
	"os"
	"encoding/json"
	"sync"
	"github.com/jcelliott/lumber"
)

const Version = "1.0.0"

type (
	Logger interface {
		Fatal(string, ...interface{})
		Error(string, ...interface{})
		Warn(string, ...interface{})
		Info(string, ...interface{})
		Debug(string, ...interface{})
		Trace(string, ...interface{})
	}

	Driver struct {
		mutex sync.Mutex
		mutexes map[string]*sync.Mutex
		dir string
		log Logger
	}

	Options struct {
		Logger
	}
)

func New(dir string, options *Options) (*Driver, error) {

}

func (d *Driver) Write() error {

}

func (d *Driver) Read() error {

}

func (d *Driver) ReadAll() () {

}

func (d *Driver) Delete() error {

}

func (d *Driver) getOrCreateMutex() *sync.Mutex {

}


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
		{"John", "23", "2312312321", "Google", Address{"san francisco", "guiug", "USA", "123123"}},
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

	allusers := []User{}
	for _, f := range records {
		empolyeeFound := User{}
		if err := json.Unmarshal([]byte(f), &empolyeeFound); err != nil {
			fmt.Println("Error", err)
		}
		allusers = append(allusers, empolyeeFound)
	}
	fmt.Println((allusers))

	if err := db.Delete("user", "john"); err != nil {
		fmt.Println("Error", err)
	}
	if err := db.Delete("user", "john"); err != nil {

	}
}

