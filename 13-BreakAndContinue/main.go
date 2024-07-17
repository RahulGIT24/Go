package main

import "fmt"

func main() {
	fmt.Println("Loops in go lang")

	days := []string{"Sunday", "Tuesday", "Wednesday", "Friday", "Saturday"}

	// fmt.Println(days)

	// for i := 0; i < len(days); i++ {
	// 	fmt.Println(days[i])
	// }

	// for i:=range days{
	// 	fmt.Println(days[i])
	// }

	for index, day := range days {
		fmt.Printf("Index is %v and value is %v\n", index, day)
	}

	rougueValue := 1

	// break
	// for rougueValue < 10 {
	// 	if rougueValue == 5 {
	// 		break
	// 	}
	// 	fmt.Println("Value is: ", rougueValue)
	// 	rougueValue++
	// }

	// continue
	// for rougueValue < 10 {
	// 	if rougueValue == 5 {
	// 		rougueValue++
	// 		continue
	// 	}
	// 	fmt.Println("Value is: ", rougueValue)
	// 	rougueValue++
	// }

	// goto
	for rougueValue < 10 {
		if rougueValue == 5 {
			goto lco
		}
		rougueValue++
	}

lco:
	fmt.Println("Jumping at line 54")
}
