package damon

import (
	"github.com/xiazemin/aladdin/damon/log"
	"github.com/xiazemin/aladdin/damon/logFile"
	"github.com/xiazemin/aladdin/damon/curl"
	"github.com/xiazemin/aladdin/damon/config"
	"github.com/xiazemin/aladdin/damon/response"
	"fmt"
)

func HandleReq(defaultDir string,ipPort *config.IpPort,configParams string,lineEnd byte,defaultFile string) string {
	logFile.LogDebug(defaultDir,*ipPort)
	confParams:=config.LocadParams(defaultDir,configParams)
	reqList:=log.Parse(defaultDir,defaultFile,lineEnd);
	logFile.LogDebug(defaultDir,"\ntotal req:\n")
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
	logFile.LogNotice(defaultDir,"\n total:"+fmt.Sprintf("%d",len(reqList))+",sucess:"+fmt.Sprintf("%d",sucess))
	return "\n total:"+fmt.Sprintf("%d",len(reqList))+",sucess:"+fmt.Sprintf("%d",sucess)
}
