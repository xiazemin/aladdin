package config

import (
	jsonEx "github.com/xiazemin/aladdin/damon/json"
	"encoding/json"
	"fmt"
	"github.com/xiazemin/aladdin/damon/logFile"
)
type IpPort struct {
	Ip string `json:"ip"`
	Port int `json:"port"`
}

func GetIpPort(dir string,filename string) *IpPort  {
	v:=new(IpPort)
	datajson:=jsonEx.Load(dir+filename)
	logFile.LogDebug(dir,string(datajson))
	err := json.Unmarshal(datajson, v)
	if(err!=nil){
		logFile.LogWarnf(dir,err)
		return &IpPort{}
	}
	return  v
}

func BuildUrl(dir string,ipPort IpPort,uri string) string {
    if(ipPort.Ip==""){
	    logFile.LogNotice(dir,ipPort)
	    return uri
    }
	if(ipPort.Port==0){
		ipPort.Port=80
	}
	logFile.LogDebug(dir,"http://"+ipPort.Ip+":"+fmt.Sprintf("%d", ipPort.Port))
	return "http://"+ipPort.Ip+":"+fmt.Sprintf("%d", ipPort.Port)+uri
}
