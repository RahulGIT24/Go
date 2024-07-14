package main

import "fmt"

//* Inside a method voluros (:=) operator is allowed but not outside
// jwtToken := 300000 error

// valid

// Constant
const LoginToken string = "sdfhkdsjf" // it is a public keyword

var jwtToken int = 300000

func main() {
	var username string = "Rahul"
	fmt.Println(username)
	fmt.Printf("Variable is of type: %T \n", username)

	var isLoogedIn bool = false
	fmt.Println(isLoogedIn)
	fmt.Printf("Variable is of type: %T \n", isLoogedIn)

	var smallVal uint8 = 255
	fmt.Println(smallVal)
	fmt.Printf("Variable is of type: %T \n", smallVal)

	var smallFloat float32 = 255.122
	fmt.Println(smallFloat)
	fmt.Printf("Variable is of type: %T \n", smallFloat)

	// default values and some aliases
	var anoterVariable int
	fmt.Println(anoterVariable)
	fmt.Printf("Variable is of type %T \n", anoterVariable)

	// implicit type
	var website = "Learncodeline.in"
	fmt.Println(website)
	fmt.Printf("Variable type is %T \n", website)

	// no var style
	numberOfUsers := 300000
	fmt.Println(numberOfUsers)
	fmt.Printf("Variable type is %T \n", numberOfUsers)

	fmt.Println(LoginToken)
	fmt.Printf("Variable type is %T \n", LoginToken)
}
