package main

import (
	"fmt"
	"io/ioutil"
	"strings"
	"net/http"
)

func main() {
	fmt.Println("Welcome to web verb video ")
	defer PerformGetRequest()
}

func PerformGetRequest() {
	const myUrl = "http://localhost:3000"

	response, err := http.Get(myUrl)
	if err != nil {
		panic(err)
	}

	defer response.Body.Close()
	fmt.Println("Status Code: ",response.StatusCode)
	fmt.Println("Content Length is: ",response.ContentLength)

	// creating a builder
	var responseString strings.Builder

	content,_  := ioutil.ReadAll(response.Body) // content is in byte format
	byteCount,_ := responseString.Write(content)

	fmt.Println(byteCount)
	fmt.Println(responseString.String())

	// fmt.Println(string(content))
}
