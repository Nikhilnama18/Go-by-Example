package main

import "fmt"



func sum(a, b int) int {
		return a + b
}

func concat(str1, str2 string) string {
		return str1 + str2
}

func swap (a, b int) (int, int) {
		return b, a
}

func infiniteSum(nums ...int) int {
	total := 0

	for index, num := range nums {
		total += num
		fmt.Printf("Index: %d, Number: %d, Current Total: %d\n", index, num, total)
	}
	return total
}

func main(){

	result := sum(5, 10)
	fmt.Println("The sum is:", result)

	resultStr := concat("Hello, ", "World!")
	fmt.Println("The concatenated string is:", resultStr)

	a:=10
	b:=20
	a, b = swap(a, b)
	fmt.Println("After swapping: a =", a, ", b =", b)


	resultInfinite := infiniteSum(1, 2, 3, 4, 5, 6)
	fmt.Println("The infinite sum is:", resultInfinite)

}
