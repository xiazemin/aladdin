package autorun

import (
	"os/exec"
	"github.com/xiazemin/aladdin/damon/logFile"
	"fmt"
)
func AutoRun(dirLog string,dir string,name string){
	cmd := exec.Command(dir+name, "a-z", "A-Z")
	logFile.LogDebug(dirLog,dir+name)
	out, err := cmd.CombinedOutput()
	if err != nil {
		logFile.LogDebug(dirLog,err)
		logFile.LogWarnf(dirLog,err)
	}
	logFile.LogDebug(dirLog,out)
	fmt.Println(string(out))
}