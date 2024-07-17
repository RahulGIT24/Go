package main

import "fmt"

func main() {
	fmt.Println("Welcome to functions in golang")
	greeter()

	//* Not allowed
	// func greeter2(){
	// 	fmt.Println("Another Method")
	// }
	// greeter2()

	// result := adder(2, 5)
	result,message := proAdder(2, 5,5,6,7)

	fmt.Println("Result is: ", result)
	fmt.Println("Pro message is: ", message)
}

// function signatures
func adder(valOne int, valTwo int) int {
	return valOne + valTwo
}

func proAdder(values ...int) (int,string) {
	total := 0
	for _,val := range values{
		total += val
	}
	return total,"HELLO"
}

func greeter() {
	fmt.Println("Namaste From GOLANG")
}
