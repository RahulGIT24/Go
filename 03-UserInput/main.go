package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("This is a simple welcome to input user programme")
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter the rating for our Pizza: ")

	// comma ok || err syntax

	input, _ := reader.ReadString('\n')
	fmt.Println("Thanks for rating, ", input)
	fmt.Printf("Type of rating is %T, ", input)

}
