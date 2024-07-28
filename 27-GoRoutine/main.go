package main

import (
	"fmt"
	"net/http"
	"sync"
)

var wg sync.WaitGroup // pointers
var signals= []string{"test"}
var mut sync.Mutex

func main(){
	// go greeter("Hi Rahul")
	// greeter("How's You")
	websiteList := []string{
		"https://google.com",
		"https://facebook.com",
		"https://instagram.com",
		"https://youtube.com",
		"https://leetcode.com",
	}

	for _,web := range websiteList{
		go getStatusCode(web)
		wg.Add(1)
	}

	wg.Wait()
	fmt.Println(signals)
}

// func greeter(s string){
// 	for i := 0; i < 6; i++ {
// 		fmt.Println(s)
// 	}
// }

func getStatusCode(endpoint string){
	defer wg.Done()
	result,err  := http.Get(endpoint)

	if err!=nil{
		fmt.Println("OOPS in endpoint")
	}

	fmt.Printf("%d status code for %s\n",result.StatusCode,endpoint)
	mut.Lock()
	signals = append(signals, endpoint)
	mut.Unlock()
}