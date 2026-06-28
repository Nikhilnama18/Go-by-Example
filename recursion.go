package main

import "fmt"

func fact(n int) int{
	if n<2{
		return 1
	}

	return n * fact(n-1)
}


func main(){

	fmt.Println("Product of first 5 numbers is: ", fact(5))

	var fibonacci func(n int) int

	fibonacci = func(n int) int{
		if n<2{
			return n
		}

		return fibonacci(n-1) + fibonacci(n-2)
	}

	fmt.Println("Fibonacci of 5 is: ", fibonacci(5))
}
