package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

func main() {
	// defer PerformPostJSONRequest()
	defer PerformPostFormRequest()
}

func PerformPostJSONRequest() {
	const myUrl = "http://localhost:3000/save"

	// fake json payload
	requestBody := strings.NewReader(`
	{
	"name":"Rahul",
	"age":12
	}
	`)

	response, err := http.Post(myUrl, "application/json", requestBody)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()
	content, _ := ioutil.ReadAll(response.Body)
	var responseString strings.Builder
	bytecount, _ := responseString.Write(content)
	fmt.Println(bytecount)
	fmt.Println(responseString.String())
}

func PerformPostFormRequest() {
	const myURL = "http://localhost:3000/postForm"

	data := url.Values{}
	data.Add("name","Rahul")
	data.Add("age","10")
	data.Add("course","BCA")

	response, err := http.PostForm(myURL,data)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()
	content,_ := ioutil.ReadAll(response.Body)
	fmt.Println(string(content))
}
