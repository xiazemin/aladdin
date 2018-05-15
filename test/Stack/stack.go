package main

import (
	"fmt"
	"runtime/debug"
)

func test1() {
	test2()
}

func test2() {
	test3()
}

func test3() {
	fmt.Printf("%s", debug.Stack())
	debug.PrintStack()
}

func main() {
	test1()
}