package main

import (
	"encoding/hex"
	"fmt"
	"bytes"
)
func main() {
	//[]byte -> String
	src := []byte("Hello")
	encodedStr := hex.EncodeToString(src)
	// 注意"Hello"与"encodedStr"不相等，encodedStr是用字符串来表示16进制
	//String -> []byte
	test, _ := hex.DecodeString(encodedStr)
	fmt.Println(bytes.Compare(test, src)) // 0
}
