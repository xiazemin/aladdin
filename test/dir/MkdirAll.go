/*
  递归创建目录
  os.MkdirAll(path string, perm FileMode) error

  path  目录名及子目录
  perm  目录权限位
  error 如果成功返回nil，如果目录已经存在默认什么都不做
*/
package main

import (
	"fmt"
	"os"
)

func main() {
	err := os.MkdirAll("/Users/didi/aladdin/xiazemin/10.96.76.97/2018-05-14-11/case1/watch/", 0777)
	if err != nil {
		fmt.Printf("%s", err)
	} else {
		fmt.Print("Create Directory OK!")
	}
	MakNewDir()
}
//该片段来自于http://outofmemory.cn

func MakNewDir()  {
	err := os.MkdirAll("/Users/didi/aladdin/xiazemin/10.96.76.97/2018-05-14-11/case1/watch/1/a/b/", 0777)
	if err != nil {
		fmt.Printf("%s", err)
	} else {
		fmt.Print("Create Directory OK!")
	}
}