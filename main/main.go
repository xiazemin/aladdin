package main

import (
	"github.com/xiazemin/aladdin/http"
	"github.com/xiazemin/aladdin/flag"
	"github.com/xiazemin/aladdin/damon"
	"github.com/xiazemin/aladdin/damon/config"
	"go/src/fmt"
	"github.com/xiazemin/aladdin/damon/logFile"
)
/**
log_format main '$remote_addr - $remote_user [$time_local]  '
        '"$request" $status $request_length $request_time $body_bytes_sent '
        '"$http_referer" "$http_user_agent" $server_addr $upstream_addr $host $upstream_cache_status $HEADER "$request_body" $operationid';
 */
const defaultFile  ="raw.log"
const globalConfig  ="globalConfig.json"
const configParams="configParams.json"
const configData  = "configData.json"
const  lineEnd  ='\n'
func main()  {
	defaultDir:=flag.GetDefaultDir()
	if(flag.IsServerType(defaultDir)){
		serv:=new(http.Serv)
		serv.Serve(defaultDir,globalConfig,configParams,lineEnd,defaultFile,configData)
		return
	}
	ipPort:=config.GetIpPort(defaultDir,globalConfig)
	userConf:=config.GetUserConf(defaultDir,globalConfig)
	logFile.LogNotice(defaultDir,userConf)
	logFiles:=config.GetSelectedLogFiles(defaultDir,configData,userConf.User,userConf.Date,userConf.Model)
	if len(logFiles)==0{
		logFiles=append(logFiles,defaultFile)
	}
	logFile.LogNotice(defaultDir,logFiles)
	 ret:=damon.HandleReq(defaultDir,ipPort,configParams,lineEnd,logFiles)
	fmt.Println(ret)
}
