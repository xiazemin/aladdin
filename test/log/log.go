package main

import (
	"log"
	"os"
)

func main(){
	arr := []int {2,3}
	log.Print("Print array ",arr,"\n")
	log.Println("Println array",arr)
	log.Printf("Printf array with item [%d,%d]\n",arr[0],arr[1])
	logf();
}
func logf(){
	// 定义一个文件
	fileName := "ll.log"
	logFile,err  := os.Create(fileName)
	defer logFile.Close()
	if err != nil {
		log.Fatalln("open file error !")
	}
	// 创建一个日志对象
	debugLog := log.New(logFile,"[Debug]",log.LstdFlags)
	debugLog.Println("A debug message here")
	//配置一个日志格式的前缀
	debugLog.SetPrefix("[Info]")
	debugLog.Println("A Info Message here ")
	//配置log的Flag参数
	debugLog.SetFlags(debugLog.Flags() | log.LstdFlags)
	debugLog.Println("A different prefix")
}