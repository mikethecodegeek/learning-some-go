package main

import "fmt"

type contactInfo struct {
	email string
}


type person struct {
	firstName string
	lastName string
	contact contactInfo
}
func main() {
	
	mike := person {
		firstName: "Mike", 
		lastName: "Sanford",
		contact: contactInfo {
			email : "mikethecodegeek@gmail.com",
		},
	}

	
	
	mike.updateName("Michael")

	mike.print()
}

func (p *person) updateName(firstName string) {
	(*p).firstName = firstName
}

func (p person) print() {
	fmt.Println(p)
}
