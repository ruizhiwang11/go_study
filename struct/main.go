package main

import "fmt"

type contactinfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	contact   contactinfo
}

func main() {

	// alex := person{"Alex", "Anderson"}
	josh := person{
		firstName: "Joth",
		lastName:  "Joth",
		contact: contactinfo{
			email:   "josh@gmail.com",
			zipCode: 123456,
		},
	}
	//fmt.Println("First Name: ", alex.firstName, "Last Name: ", alex.lastName)
	fmt.Println("First Name: ", josh.firstName, "Last Name: ", josh.lastName)
	//fmt.Println(alex)
	josh.print()
	josh.updateName("coco")
	josh.print()
}
func (p *person) updateName(newFirstName string) {
	p.firstName = newFirstName
}

func (p person) print() {
	fmt.Printf("%+v\n", p)
}
