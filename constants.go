package main

import (
	"fmt"
	"math"
)

const s = "Constanst String"

func main(){
	fmt.Println(s)

	// s="New String" Throws error
	const a = 50

	const b = 3e20 / a
	fmt.Println(b)

	fmt.Println(int64(b))

	fmt.Println(math.Sin(a))
}
