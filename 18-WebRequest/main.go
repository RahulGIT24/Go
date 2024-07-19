package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const url = "https://github.com/RahulGIT24"

func main() {
	fmt.Println("LCO Webrequest")
	response, err := http.Get(url)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Respose type is %T", response)

	defer response.Body.Close()

	databytes, err := ioutil.ReadAll(response.Body)

	if err != nil {
		panic(err)
	}
	content := string(databytes)

	fmt.Println(content)
}
