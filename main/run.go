package main

import (
	"github.com/xiazemin/aladdin/autorun"
	"github.com/xiazemin/aladdin/damon/logFile"
	"runtime"
)

const dirLog  ="/Users/didi/aladdin/xiazemin/10.96.76.97/2018-05-14-11/case1/watch/"//"./"
const dirConfig  ="/Users/didi/aladdin/xiazemin/10.96.76.97/2018-05-14-11/case1/"
//"./"
const dirExe  ="/Users/didi/goLang/"
//"./"
const nameExe  ="aladdin"

func main()  {
	var paths []string
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
