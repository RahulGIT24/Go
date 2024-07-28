package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/RahulGIT24/gomongo/route"
)

func main(){
	fmt.Println("MongoDB")
	r := route.Router()
	fmt.Println("Server is getting started...")
	log.Fatal(http.ListenAndServe(":4000",r))
	fmt.Println("Server is running on port 400")
}