package main

import "fmt"

type Person struct {
	firstname string
	lastname  string
}

// Contact {firstname, lastname, phone, email, age}
type Contact struct {
	p     Person
	phone int64
	email string
	age   uint8
}

func main() {
	var c1 *Contact = new(Contact)

	// +Read Contact
	read(c1)

	fmt.Printf("%+v\n", c1)

	// + Display Contact
	display(*c1)
	// fmt.Printf("%+v\n", c1)
}

func read(c *Contact) {
	fmt.Printf("Enter firstname: ")
	fmt.Scanf("%s", &c.p.firstname)
	fmt.Printf("Enter lastname: ")
	fmt.Scanf("%s", &c.p.lastname)
	fmt.Printf("Enter phone: ")
	fmt.Scanf("%d", &c.phone)
	fmt.Printf("Enter email: ")
	fmt.Scanf("%s", &c.email)
	fmt.Printf("Enter age: ")
	fmt.Scanf("%d", &c.age)
}

func display(c Contact) {
	fmt.Printf("Firstname: %s\n", c.p.firstname)
	fmt.Printf("Lastname: %s\n", c.p.lastname)
	fmt.Printf("Phone: %d\n", c.phone)
	fmt.Printf("Email: %s\n", c.email)
	fmt.Printf("Age: %d\n", c.age)
}
