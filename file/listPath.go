package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	myfolder := `/Users/didi/`
	listFile(myfolder)
}

func listFile(myfolder string) {
	files, _ := ioutil.ReadDir(myfolder)
	for _, file := range files {
		if file.IsDir() {
			listFile(myfolder + "/" + file.Name())
		} else {
			fmt.Println(myfolder + "/" + file.Name())
		}
	}
}