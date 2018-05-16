package main

import (
	"github.com/xiazemin/aladdin/autorun"
	"github.com/xiazemin/aladdin/damon/logFile"
	"github.com/xiazemin/aladdin/flag"
	"runtime"
	"os"
	"fmt"
)
//"./"
const dirExe  ="/Users/didi/goLang/"
//"./"
const nameExe  ="aladdin"

func main()  {
	dirConfig:=flag.GetDefaultDir()
	var paths []string
	errDir := os.MkdirAll(dirConfig+"watch/", 0777)
	if errDir != nil {
		fmt.Printf("%s", errDir)


	}
	dirLog:=dirConfig+"watch/"
	logFile.LogNotice(dirLog,dirConfig)//cp logConfig.json watch/
	paths=append(paths,autorun.GetPath(dirLog,dirConfig,"configWatch.json"))
	logFile.LogNotice(dirLog,paths)
	exit := make(chan bool)
	autorun.NewWatcher(dirLog,paths, dirExe,nameExe)
	autorun.AutoRun(dirLog,dirExe,nameExe)
	for {
		select {
		case <-exit:
			runtime.Goexit()
		}
	}
}
