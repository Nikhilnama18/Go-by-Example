package main

import "fmt"

func main() {

	fmt.Println("Range examples")

	nums := []int{2, 4, 6, 8}
	sum := 0
	for _, n := range nums {
		sum += n
	}
	fmt.Println("Sum is: ", sum)

	for i, n := range nums {
		if n == 6 {
			fmt.Println("Index of 6 is: ", i)
		}
	}

	keyValues := map[string]string{"a": "apple", "b": "banana"}
	for k, v := range keyValues {
		fmt.Println(k, v)
	}

	for k := range keyValues {
		fmt.Println("Key is: ", k)
	}

	for i, c := range "golang" {
		fmt.Println(i, c)
	}
}
