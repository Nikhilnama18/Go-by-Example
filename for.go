package main

import "fmt"

func main(){

	i:= 1
	for i <= 4{
		fmt.Println("I is ", i)
		i++
	}

	for j :=1; j<=4; j++ {
		fmt.Println("J is ", j)
	}

	for k := range 3{
		fmt.Println("K is ", k)
	}

	for {
		fmt.Println("This is an infinite loop with break statement")
		break
	}

	
	for n:= range 10 {
		if n%2 == 0 {
			continue
		}
		fmt.Println("N is ", n)
	}

	for {
		fmt.Println("This is an infinite loop with return statement")
		return
	}

	// This creates an infinte loop & also this code is not executed because the previous loop has a return statement which will terminate the program
	// for {
	// 	fmt.Println("This is an infinite loop with continue statement")
	// 	continue
	// }
}
