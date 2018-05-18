package autorun

import (
	"os/exec"
	"github.com/xiazemin/aladdin/damon/logFile"
	"github.com/xiazemin/aladdin/flag"
	"github.com/xiazemin/aladdin/damon/curl"
	"github.com/xiazemin/aladdin/damon/config"
	"github.com/xiazemin/aladdin/damon/netenv"
	"fmt"
)
const globalConfig  ="globalConfig.json"

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
	defaultDir:=flag.GetDefaultDir()
	ipPort:=config.GetIpPort(defaultDir,globalConfig)
	if flag.IsRemoteServer(defaultDir){
		ip := netenv.GetLocalIp(defaultDir)
		logFile.LogNotice(dirLog, fmt.Sprintf(" local ip %s", ip))
		if ip!="" {
		ipPort.Ip = ip
		}
	}
	userConf:=config.GetUserConf(defaultDir,globalConfig)
	url:="http://"+conf.Ip+fmt.Sprintf(":%d",conf.Port)+"/aladdin?"+fmt.Sprintf("ip=%s&port=%d&user=%s&model=%s&date=%s",ipPort.Ip,ipPort.Port,userConf.User,userConf.Model,userConf.Date)
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