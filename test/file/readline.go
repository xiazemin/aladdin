package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main()  {
	f, err := os.Open("/Users/didi/aladdin/xiazemin/10.96.76.97/2018-05-14-11/case1/raw.log")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	rd := bufio.NewReader(f)
	for {
		line, err := rd.ReadString('\n') //以'\n'为结束符读入一行

		if err != nil || io.EOF == err {
			break
		}
		fmt.Println(line)
	}
}
