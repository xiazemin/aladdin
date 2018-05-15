package main

import (
	"fmt"
	"time"
)

func main() {
	str := "11/May/2018:22:55:35 +0800"

	t, _ := time.Parse("02/Jan/2006:15:04:05 +0800",str)
	fmt.Println(t)
	fmt.Println(t.Unix())
}