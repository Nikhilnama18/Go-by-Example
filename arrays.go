package main

import "fmt"

func main() {
	var a [5]int

	a[0] = 2

	fmt.Println(a)

	b := [3]int{1, 2, 3}
	fmt.Println(b)

	c := [...]int{4, 5, 6, 7}
	fmt.Println(c)

	var d [3]string

	for i := 0; i < 3; i++ {
		fmt.Scan(&d[i])
	}
	fmt.Println(d)

	var e [3][2]int

	for i := 0; i < 3; i++ {
		for j := 0; j < 2; j++ {
			e[i][j] = i + j
			fmt.Print(e[i][j], " ")
		}
		fmt.Println()
	}
}
