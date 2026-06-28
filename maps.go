package main

import (
	"fmt"
	"maps"
)

func main(){

	m:= make(map[string]int)

	m["a"] = 1
	m["b"] = 2

	fmt.Println(m)

	val:= m["a"]
	fmt.Println(val)

	
	delete(m, "a")
	fmt.Println(m)

	clear(m)
	fmt.Println(m)

	k:= map[string]string{"foo":"bar", "baz":"qux"}

	fmt.Println(k)

	n:= map[string]string{"foo":"bar", "baz":"qux"}

	if maps.Equal(k, n){
		fmt.Println("Both maps are equal")
	}
}
