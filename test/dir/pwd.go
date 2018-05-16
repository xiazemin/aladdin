package main

import (
	"os"
	"fmt"
)

const(
a = iota
	b
)
func main() {
	var path string
	if os.IsPathSeparator('\\') {
		path = "\\"
	} else {
		path = "/"
	}
	pwd, _ := os.Getwd()
	fmt.Println(pwd)
	fmt.Println("test")
	fmt.Println(path)
	//err := os.Mkdir(pwd+path+"tmp", os.ModePerm)
	//if err != nil {
	//	fmt.Println(err)　　
	//}
	fmt.Println(a)
	fmt.Println(b)
}
