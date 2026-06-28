package main

import (
	"fmt"
	"slices"	
)

func main(){

	slice1:= []int{1, 2, 3, 4, 5}

	fmt.Println("Original slice: ", slice1)

	slice1 = append(slice1, 6)
	fmt.Println("Slice after appending 6: ", slice1)

	slice2 := make([]int, 3)
	slice2 = slice1[2:5]
	fmt.Println("Slice2 after slicing slice1 from index 2 to 5: ", slice2)


	t := []string{"g", "h", "i"}
    fmt.Println("dcl:", t)

    t2 := []string{"g", "h", "i"}
    if slices.Equal(t, t2) {
        fmt.Println("t == t2")
    }
}
