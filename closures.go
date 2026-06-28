package main

import "fmt"

func defaultSeq() func() int{
	i:=0
	return func() int{
		i++
		return i
	}
}

func initialSeq(start int) func() int{
	i:=start-1
	return func() int{
		i++
		return i
	}
}


func main(){

	defaultCounter := defaultSeq()
	initialCounter := initialSeq(10)

	fmt.Println(defaultCounter()) // Output: 1
	fmt.Println(defaultCounter()) // Output: 2
	fmt.Println(defaultCounter()) // Output: 3

	fmt.Println(initialCounter()) // Output: 10
	fmt.Println(initialCounter()) // Output: 11

	newDefaultCounter := defaultSeq()
	fmt.Println(newDefaultCounter()) // Output: 1
	fmt.Println(newDefaultCounter()) // Output: 2
}
