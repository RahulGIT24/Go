package main

import "fmt"

func main() {
	fmt.Println("Maps")

	languages := make(map[string]string)

	languages["JS"] = "Javascript"
	languages["RB"] = "Ruby"
	languages["PY"] = "Python"

	fmt.Println("List of all languages: ", languages)
	fmt.Println(languages["JS"])

	delete(languages, "JS")
	fmt.Println(languages)

	// loops are interesting in go lang

	for key, value := range languages {
		fmt.Printf("For key %v, value is %v\n", key, value)
	}
}
