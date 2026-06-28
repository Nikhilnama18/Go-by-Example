package main

import "fmt"

type Person struct {
	name string
	age  int
}

func newPerson(name string, age int) *Person {
	return &Person{name: name, age: age}
}

func main() {

	fmt.Println("Person is ", Person{"Nikhil", 20})

	fmt.Println("Person 2 is ", Person{name: "Suhas", age: 21})

	fmt.Println("Person 3 is ", newPerson("Nick", 26))

	fmt.Println("Person 4 is ", &Person{"Varun", 34})

	p1 := newPerson("Arjun", 40)

	p2 := p1

	p2.age = 32

	fmt.Println("P1 & P2 are ", p1, p2)

	p3 := Person{"Ravi", 32}
	p4 := &p3

	p4.age = 33

	fmt.Println("P3 & P4 are ", p3, p4)

	dog := struct {
		name  string
		breed string
	}{
		"Mac",
		"Golden Retriever",
	}

	fmt.Println("New dog in town is ", dog)
}
