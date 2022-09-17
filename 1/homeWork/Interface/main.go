// 4 ta turdagi mahsulot methodlarini yozish kere
// Interface ishlatish kere

// struct
// method

package main

import (
	"fmt"
)

func main() {
	tel := Telephone{
		2022,
		"IPhone 14",
		16,
	}

	wallaper(tel)

	cars := Car{"Camaro", "Chevrolet", 2022, "Gasoline"}

	wallaper(cars)

	home := Home{"Nest One", "Murad buildings", 51, 2022}

	wallaper(home)

	comp := Computer{"Aser Aspire-5", "ACER", "noutbook", 17.4, 2022}

	wallaper(comp)
}

// ========================Telephone==========================

// struct
type Telephone struct {
	year    int
	name    string
	monitor float64
}

// method

func (a Telephone) Display() string {
	return fmt.Sprintf("Published year: %d, Product's Name: %s,  Display size: %f", a.year, a.name, a.monitor)
}

// ========================Car==========================

// struct
type Car struct {
	name     string
	brend    string
	age      int
	fuelType string //
}

// method
func (c Car) Display() string {
	return fmt.Sprintf("Car name %v, Car Brend %v, Car Product in %v, Car fuel type %v", c.name, c.brend, c.age, c.fuelType)
}

// ========================Home==========================

// struct
type Home struct {
	name     string
	homeType string
	floor    int
	age      int
}

// method

func (h Home) Display() string {
	return fmt.Sprintf("Home name: %s, Home brend: %s, The total floor of the house: %d, The year the house was built: %d", h.name, h.homeType, h.floor, h.age)

}

// ========================Computer==========================

// struct
type Computer struct {
	name     string
	brend    string
	compType string
	display  float64
	age      int
}

// method

func (cs Computer) Display() string {
	return fmt.Sprintf("computer name: %s, Computer brend: %s, Computer type: %s, Display size: %0.1f, Year: %d", cs.name, cs.brend, cs.compType, cs.display, cs.age)
}

//============================================================

// interface

type makeInterface interface {
	Display() string
}

func wallaper(obj makeInterface) {
	switch obj.(type) {
	case Telephone:
		fmt.Println("Product Telephone-->")
	case Car:
		fmt.Println("Product Car-->")
	case Home:
		fmt.Println("Product Home-->")
	case Computer:
		fmt.Println("Product Computer-->")

	}

	fmt.Println(obj.Display())
	fmt.Println()
}
