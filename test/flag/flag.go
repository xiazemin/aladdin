package main

import (
	"flag"
	"fmt"
)

var music_file *string = flag.String("file", "musicfile", "Use -file <filesource>")

func main() {
	flag.Parse()
	fmt.Println(*music_file)
	getInt()
}

var flagvar int
func getInt() {
	flag.IntVar(&flagvar, "flagname", 1234, "help message for flagname")
	fmt.Println(fmt.Sprintf("%d",flagvar))
	var ip = flag.Int("flagname1", 1234, "help message for flagname")
	fmt.Println(fmt.Sprintf("%d",*ip))
	//flag.Var(&flagVal, "name", "help message for flagname")
}


