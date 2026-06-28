package main

import "fmt"

func zeroVal( val int){
	val = 0
}

func zeroValPtr( val *int){
	*val = 0
}

func main(){

	num := 5
	fmt.Println("Original value: ", num)
	zeroVal(num)
	fmt.Println("Value after zeroVal: ", num)

	zeroValPtr(&num)
	fmt.Println("Value after zeroValPtr: ", num)

	fmt.Println("Address of num is: ",&num)
}
