package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	//contactInfo   contactInfo
	// a simpler way to declare it is
	contactInfo
}

func main() {
	// //alex := person{"Alex", "Anderson"}
	// vagner := person{firstName: "Vagner", lastName: "Oliveira"}
	// //fmt.Println(alex)
	// fmt.Println(vagner)

	// // Variable with zero values
	// // string = ""
	// // int, float = 0
	// // bool = false
	// // Go doesn't assign nill by default as other languages
	// var john person
	// fmt.Println(john)
	// // %+v will print out all the field names and values from john
	// fmt.Printf("%+v", john)

	jim := person{
		firstName: "Jim",
		lastName:  "Party",
		contactInfo: contactInfo{
			email:   "jim@gmail.com",
			zipCode: 94000,
		},
	}

	//jimPointer := &jim
	//jimPointer.updateName("Jimmy")

	// Pointer shortcut
	jim.updateName("Jimmy")
	jim.print()

	alex := person{
		firstName: "Alex",
		lastName:  "Anderson",
		contactInfo: contactInfo{
			email:   "alex@gmail.com",
			zipCode: 94001,
		},
	}

	//alex.firstName = "Alexandre"
	alex.print()

}

func (pointerToPerson *person) updateName(newFirstName string) {
	(*pointerToPerson).firstName = newFirstName
}

func (p person) print() {
	fmt.Printf("%+v", p)
}
