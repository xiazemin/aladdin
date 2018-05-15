package main

import (
	"os/exec"
	"fmt"
)
func main() {
	f, err := exec.LookPath("ls")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(f) //  /bin/ls
}