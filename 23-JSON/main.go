package main

import (
	"encoding/json"
	"fmt"
)

type course struct {
	Name     string `json:"coursename"`
	Price    int
	Platform string   `json:"website"`
	Password string   `json:"-"`
	Tags     []string `json:"tags,omitempty"`
}

func main() {
	fmt.Println("JSON")
	// EncodeJSON()
	DecodeJSON()

}

func EncodeJSON() {

	lcoCourses := []course{
		{"ReactJS", 1200, "LearnCodeOnline", "Abc@1234", []string{"Web dev", "JS"}},
		{"AI", 1100, "LearnCodeOnline", "amit@1234", []string{"Web dev", "JS"}},
		{"MERN", 11200, "LearnCodeOnline", "ht@1234", nil},
	}

	// package this data as json data
	finalJSON, err := json.MarshalIndent(lcoCourses, "", "\t")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", finalJSON)
}

func DecodeJSON() {
	jsonDataFromWeb := []byte(`
	{
                "coursename": "ReactJS",
                "Price": 1200,
                "website": "LearnCodeOnline",
                "tags": [
                        "Web dev",
                        "JS"
                ]
        }
	`)

	var lcoCourse course

	checkValid := json.Valid(jsonDataFromWeb)
	if checkValid {
		fmt.Println("JSON was valid")
		json.Unmarshal(jsonDataFromWeb, &lcoCourse)
		fmt.Printf("%#v\n", lcoCourse)
	} else {
		fmt.Println("The json was not valid")
	}

	// some cases where you want to data to key value

	var myOnlineData map[string]interface{}

	json.Unmarshal(jsonDataFromWeb,&myOnlineData)

	fmt.Printf("%#v\n", myOnlineData)

	for k,v := range myOnlineData{
		fmt.Printf("Key is %v and value is %v\n",k,v)
	}
}
