package main

import (
	"github.com/xiazemin/aladdin/http"
	"github.com/xiazemin/aladdin/flag"
	"github.com/xiazemin/aladdin/damon"
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

	if(flag.IsServerType(defaultDir)){
		serv:=new(http.Serv)
		serv.Serve(defaultDir,globalConfig,configParams,lineEnd,defaultFile)
		return
	}
	 ret:=damon.HandleReq(defaultDir,globalConfig,configParams,lineEnd,defaultFile)
	fmt.Println(ret)
}
