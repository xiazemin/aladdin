package main

import (
	"log"
	"os"
	"fmt"
)

func main() {
	originalPath := "./test.txt"
	newPath := "test_new.txt"
	err := os.Rename(originalPath, newPath)
	if err != nil {
		log.Fatal(err)
	}

	originalPath1 := "test"
	newPath1 := "test_new"
	err1 := os.Rename(originalPath1, newPath1)
	if err != nil {
		log.Fatal(err1)
	}

	originalPath2 := "test.txt"
	result := Exists(originalPath2)
	fmt.Println(result)
}

func Exists(name string) bool {
	if _, err := os.Stat(name); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}