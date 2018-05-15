package main

import (
	"github.com/xiazemin/aladdin/damon/log"
	"github.com/xiazemin/aladdin/damon/logFile"
	"github.com/xiazemin/aladdin/damon/url"
	"github.com/xiazemin/aladdin/damon/json"
	"github.com/xiazemin/aladdin/damon/curl"
	"github.com/xiazemin/aladdin/damon/config"
	"github.com/xiazemin/aladdin/damon/response"
	"go/src/fmt"
)
/**
log_format main '$remote_addr - $remote_user [$time_local]  '
        '"$request" $status $request_length $request_time $body_bytes_sent '
        '"$http_referer" "$http_user_agent" $server_addr $upstream_addr $host $upstream_cache_status $HEADER "$request_body" $operationid';
 */
const defaultDir  ="/Users/didi/aladdin/xiazemin/10.96.76.97/2018-05-14-11/case1/"
const defaultFile  ="raw.log"
const globalConfig  ="globalConfig.json"
const configParams="configParams.json"
const  lineEnd  ='\n'
func main()  {

	logFile.LogDebug(defaultDir,log.GetAlphaTab("ie=utf-8&f=8&rsv_bp=1&rsv_idx=1&tn=baidu&wd=%20invalid%20character%20%27%7B%27%20after%20top-level%20value&oq=golang%2520%25E8%25B0%2583%25E7%2594%25A8%25E6%25A0%2588&rsv_pq=cd93e94d00051db9&rsv_t=1c5eYwmNwWfA31oioXIREysIkHFRAwl1xCgZlt79euRIPpkelmnkBE9uv4k&rqlang=cn&rsv_enter=0&inputT=483&rsv_sug3=721&rsv_sug2=0&rsv_sug4=483"))
	reqList:=log.Parse(defaultDir,defaultFile,lineEnd);
	//[$time_local] "$request"
	line:=`
	10.0.0.0 - - [11/May/2018:22:55:35 +0800]  "POST /abc/def/v1/hgk/hhhhaave HTTP/1.1" 200 427 0.010 745 "-" "-" 10.0.0.0 127.0.0.1:9000 10.0.0.0 - 13jhgd "k_id=134255&p_id=456787989&version=1.0.0" 454434
	`
	r:=new(log.Request)
	logFile.LogDebug(defaultDir,r.ParseTime(defaultDir,line))
	logFile.LogDebug(defaultDir,url.ToJson("l_id=34544534&y_id=09879&version=1.0.0"))
	for key,value:=range (url.ToJson("lid=56454&h_id=456787&version=1.0.0")){
		if str, ok := value.(string);ok{
			logFile.LogDebug(defaultDir,key+"=>"+str)
		}

	}
	logFile.LogDebug(defaultDir,json.ToForm(defaultDir,url.ToJson("lid=4565&j_id=456756546&version=1.0.0")))

	logFile.LogDebug(defaultDir,"\ntotal req:\n")

	ipPort:=config.GetIpPort(defaultDir,globalConfig)
	logFile.LogDebug(defaultDir,*ipPort)
	confParams:=config.LocadParams(defaultDir,configParams)
	logFile.LogDebug(defaultDir,len(reqList))
	var sucess int
	sucess=0
	for id,req:=range(reqList) {
		logFile.LogDebug(defaultDir,id)
                req=req.ForMatRequest(defaultDir,req,ipPort,*confParams)
		resp:=curl.QueryForm(defaultDir,*req)
		logFile.LogDebug(defaultDir,resp)
		res:=response.Parse(defaultDir,resp)
		if res.Errno==0{
			sucess++
		}else {
			logFile.LogNotice(defaultDir,req.Uri)
		}
		logFile.LogDebug(defaultDir,res)
		config.UpdateParams(defaultDir, configParams, *res)
	}
	fmt.Println("\n total:"+fmt.Sprintf("%d",len(reqList))+",sucess:"+fmt.Sprintf("%d",sucess))

}