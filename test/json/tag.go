package main

import (
	"fmt"
	"encoding/json"
	"log"
)
type Account struct {
	Email    string  `json:"email"`
	Password string  `json:"password"`
	Money    float64 `json:"money,string"`
}

var jsonString string = `{
    "email": "phpgo@163.com",
    "password" : "123456",
    "money" : "100.5"
}`
//  "money" : "100.5" 不能没有 双引号，否则会报错

func main() {

	account := Account{}
	err := json.Unmarshal([]byte(jsonString), &account)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%+v\n", account)
}
