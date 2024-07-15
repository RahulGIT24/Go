package main

import "fmt"

func main() {
	fmt.Println("Arrays")

	// syntax
	var fruitList [4]string

	fruitList[0] = "Apple"
	fruitList[1] = "Mango"
	fruitList[3] = "Blueberry"

	fmt.Println("Fruit list is :", fruitList)
	fmt.Println("Length of Fruit list is :", len(fruitList))

	// this is also valid
	var vegList = [3]string{"Tomatoes", "Potatoes", "Onions"}
	fmt.Println("Vegy list: ", vegList)
	fmt.Println("Length of Vegy list: ", len(vegList))

}
