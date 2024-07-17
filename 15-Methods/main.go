package main

import "fmt"

func main() {
	fmt.Println("Methods")
	hitest := User{"Hitest", "hitest@go.com", true, 19}
	hitest.GetStatus()
	email := hitest.NewMail()
	fmt.Println(email)
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

func (u User) GetStatus() {
	if u.Status == true {
		fmt.Println("Online")
	} else {
		fmt.Println("Offline")
	}
}

func (u User) NewMail() string {
	u.Email = "test@go.dev"
	return u.Email
}
