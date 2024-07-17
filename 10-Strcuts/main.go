package main

import "fmt"

func main() {
	fmt.Println("Structs in golang")

	// no inheritance in golang, No super or parent

	hitesh := User{"Rahul", "rg4005450@gmail.com", true, 18}

	// Access values of struts
	fmt.Println(hitesh.Name)

	// Gives you the full details
	fmt.Printf("Hitesh details are: %+v\n", hitesh)

	// Way to print values 
	fmt.Printf("Name is %v and email is %v\n", hitesh.Name, hitesh.Email)
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}
