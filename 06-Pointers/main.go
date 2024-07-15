package main

import "fmt"

func main() {
	fmt.Println("Class on Pointers")

	// Creating a pointer
	// var ptr *int

	// fmt.Println("Value of pointer is ", ptr)

	myNumber := 24

	// Creating a pointer with refrencing
	var ptr = &myNumber

	fmt.Println("Value of actual pointer is ", *ptr)
	fmt.Printf("Data type of actual pointer is %T \n", ptr)
	fmt.Printf("Data type of variable is %T \n", *ptr)

	*ptr = *ptr * 2

	fmt.Println("Value of actual pointer is ", myNumber)
}
