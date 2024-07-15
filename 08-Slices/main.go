package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Slices in GO")

	var fruitList = []string{"Apple", "Tomato", "Peach"}

	fmt.Printf("Type of fruitlist is %T \n", fruitList)

	// not valid
	// fruitList[0]

	fruitList = append(fruitList, "Mango", "Banana")
	fmt.Println(fruitList)

	fruitList = append(fruitList[:4])
	fmt.Println(fruitList)

	highScores := make([]int, 4)

	highScores[0] = 234
	highScores[1] = 2
	highScores[2] = 24
	highScores[3] = 23
	// highScores[4] = 211 // out of bound

	// You can append using this way, append reallocate the memory
	highScores = append(highScores, 555, 122, 33, 45, 0)

	fmt.Println(highScores)

	sort.Ints(highScores)

	fmt.Println(highScores)
	fmt.Println(sort.IntsAreSorted(highScores))

	// how to remove a value from slice based on index

	var courses = []string{"reactjs", "javascript", "swift", "python", "ruby"}
	var index int = 2
	courses = append(courses[:index],courses[index+1:]...)
	fmt.Println(courses)
}
