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
const nameConf="configWatch.json"

func main()  {
	dirConfig:=flag.GetDefaultDir()
	var paths []string
	errDir := os.MkdirAll(dirConfig+"watch/", 0777)
	if errDir != nil {
		fmt.Printf("%s", errDir)


	}
	dirLog:=dirConfig+"watch/"
	logFile.LogNotice(dirLog,dirConfig)//cp logConfig.json watch/
	for _,p:=range (autorun.GetPath(dirLog,dirConfig,nameConf)) {
		paths = append(paths, p)
	}

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
