package main

import "fmt"

func main() {
	fmt.Println("If else statement")

	loginCount := 10

	var result string

	// You can't move curly brace down
	if loginCount < 10 {
		result = "Regular User"
	} else if loginCount > 10 {
		result = "Unauthorized User"
	} else {
		result = "Exactly 10 login count"
	}

	fmt.Println(result)

	if 9%2 == 0 {
		fmt.Printf("Even\n")
	} else {
		fmt.Printf("Odd\n")
	}

	if num := 3; num<10{
		fmt.Println("Number is less than 10")
	}else{
		fmt.Println("Number is not less than 10")
	}
}
