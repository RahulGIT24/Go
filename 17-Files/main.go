package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	fmt.Println("Welcome to files in golang")
	content := "This needs to go in a file - LearnCodeOnline.in"

	file, err := os.Create("./mygofile.txt")

	checkNilError(err)

	length, err := io.WriteString(file, content)

	checkNilError(err)

	fmt.Println("Length is: ", length)
	readFile("./mygofile.txt")
	defer file.Close()
}

func readFile(filename string) {
	dataByte, err := ioutil.ReadFile(filename)

	checkNilError(err)

	fmt.Println("Text data inside the file is \n", string(dataByte))
}

func checkNilError(err error) {
	if err != nil {
		panic(err)
	}
}
