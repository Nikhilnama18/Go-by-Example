package main

import "fmt"

func main(){
	
	i:=10

	if i%2 < 20{
		fmt.Println("I is less than 20")
	}

	if k:= 3; k%2 ==0{
		fmt.Println("K is an even number")
	}else {
		fmt.Println("K is an odd number")
	}

	b:= 100 

	if b>50 && b%100 == 0{
		fmt.Println("B is greater than 50 and divisible by 100")
	}else if b>50 && b%100 != 0{
		fmt.Println("B is greater than 50 but not divisible by 100")
	}else{
		fmt.Println("B is less than or equal to 50")
	}
}
