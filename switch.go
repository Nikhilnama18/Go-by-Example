package main

import (
	"fmt"
	"time"
)

func main() {

	i := 2

	switch i {
	case 1:
		fmt.Println("I is 1")
	case 2:
		fmt.Println("I is 2")
	default:
		fmt.Println("I is not 1 or 2")
	}

	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("It's the weekend")
	default:
		fmt.Println("It's a weekday")
	}

	t := time.Now() // Swithc without no expression is an alternative way to express if/else logic.
	switch {
	case t.Hour() < 12:
		fmt.Println("It's before noon")
	default:
		fmt.Println("It's after noon")
	}
}
