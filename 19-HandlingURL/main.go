package main

import (
	"fmt"
	"net/url"
)

const myurl = "https://github.com/RahulGIT24?tab=repositories&chatbot=vercel"

func main() {
	fmt.Println("Handling URLs in Go Lang")
	fmt.Println(myurl)

	// parsing url
	result, _ := url.Parse(myurl)

	fmt.Println(result.Scheme)
	fmt.Println(result.Host)
	fmt.Println(result.Path)
	fmt.Println(result.RawQuery)
	fmt.Println(result.Port())

	qparams := result.Query()

	fmt.Printf("Type of qparams is %T \n", qparams)
	fmt.Println(qparams["tab"])

	for _, val := range qparams {
		fmt.Println("Param is :", val)
	}

	// Remember pass reference
	partsOfURL := &url.URL{Scheme: "https", Host: "github.com", Path: "/RahulGIT24", RawQuery: "tab=repositories&chatbot=vercel"}

	anotherURL := partsOfURL.String()
	fmt.Println(anotherURL)
}
