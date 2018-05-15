package main

import (
	"fmt"
	"strings"
)

func main() {
	str := []string{"Hello", "World", "Good"}
	fmt.Println(strings.Join(str, " "))
	lastIndex();
	repeat();
	replace();
	split();
	splitAfter();
	test:="123456789"
	fmt.Print(strings.Split(test,"0"))
	fmt.Print(strings.Split(test,"7"))
}
//程序输出 Hello World Good

func lastIndex()  {
	str := "Hello World"
	fmt.Println(strings.LastIndex(str, "l"))
}

func repeat()  {
	str := "Hello "
	fmt.Println(strings.Repeat(str, 5))
}

func replace()  {
	str := "hi hi hi are you ok"
	fmt.Println(strings.Replace(str, "hi", "ok", 3))
}

func split()  {
	str := "one,two,three"
	for _, v := range strings.Split(str, ",") {
		fmt.Println(v)
	}
}

func splitAfter()  {
	str := "one,two,three"
	for _, v := range strings.SplitAfter(str, ",") {
		fmt.Println(v)
	}
}