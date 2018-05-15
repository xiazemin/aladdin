package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	myfolder := `/Users/didi/`

	files, _ := ioutil.ReadDir(myfolder)
	for _, file := range files {
		if file.IsDir() {
			continue
		} else {
			fmt.Println(file.Name())
		}
	}
}