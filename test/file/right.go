package main

import (
	"log"
	"os"
)

func main() {

	//Write permission
	file, err := os.OpenFile("./test.txt", os.O_WRONLY, 0666)
	if err != nil {
		if os.IsPermission(err) {
			log.Println("Error: Write permission denied.")
		}
	}
	file.Close()

	//Read permission
	file, err = os.OpenFile("./test.txt", os.O_RDONLY, 0666)
	if err != nil {
		if os.IsPermission(err) {
			log.Println("Error: Read permission denied.")
		}
	}
	file.Close()
}