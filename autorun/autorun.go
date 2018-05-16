package autorun

import (
	"os/exec"
	"github.com/xiazemin/aladdin/damon/logFile"
	"github.com/xiazemin/aladdin/flag"
	"github.com/xiazemin/aladdin/damon/curl"
	"fmt"
)

type RemoteConf struct {
	Ip string
	Port int
}
func AutoRun(dirLog string,dir string,name string){
	if !remoteRun(dirLog) {
		localRun(dirLog, dir, name)
	}

}

func localRun(dirLog string,dir string,name string)  {
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

func remoteRun(dirLog string) bool {
     conf:=GetRemoteConf(dirLog)
	if(conf.Ip==""){
		return false
	}

	url:="http://"+conf.Ip+fmt.Sprintf(":%d",conf.Port)+"/aladdin"
	resp:=curl.QueryFormSimple(dirLog,url)
	logFile.LogNotice(dirLog,resp)
	logFile.LogDebug(dirLog,resp)
	if(resp!="") {
		return true
	}else{
		logFile.LogWarnf(dirLog,resp)
		return false
	}
}

func GetRemoteConf(dirLog string) *RemoteConf {
	ip,port:=flag.GetIpPort(dirLog)
	conf:=new(RemoteConf)
	conf.Ip=ip
	conf.Port=port
	return conf
}