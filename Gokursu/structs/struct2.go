package structs

import "fmt"

func (c Customer) Save() {
	fmt.Println("Eklendi",c.firstName, c.lastName, c.age)
}
func (c *Customer) Update() {
	fmt.Println("Update() function called")
	// Your update logic here
	c.updated = true // Set the flag to true after update
}

/// farklari


func Struct2() {
	c := Customer{firstName: "Rahil", lastName: "IBRAIHMI", age: 22, guncelName: "Arya"}
	c.Save()
	c.Update()
}

type Customer struct {
	firstName string
	lastName  string
	age       int
	guncelName	string
	updated  bool
}

func (c Customer) Delete() {
	fmt.Println("Silindi")
}

func (c Customer) Sale() {
	fmt.Println("Satildi")
}
