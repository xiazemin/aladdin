package log

import (
	"strings"
	"regexp"
	"time"
	"github.com/xiazemin/aladdin/damon/url"
	"github.com/xiazemin/aladdin/damon/config"
	"github.com/xiazemin/aladdin/damon/logFile"
)
type Request struct {
	Method string
	Uri string
	Param string
	Version string
	Time int64 //s
	Url string
	Arguments map[string]interface{}
}
const length  =3
const paramLength  =1
const timeLength  =1
const uriSep  =" "
const lineSep  ="\""
const timeSepS  ="["
const timeSepE  ="]"

func GetAlphaTab(line string) map[string] int  {
	var alphaTab=make(map[string] int)
	for _,val:=range (line){
		//fmt.Println(val)
		//fmt.Println(uint(val))
		alphaTab[(string(val))]++
		//rune
	}
	return alphaTab
}

func (this * Request) ParseUri(dir string,request string,sep string)*Request{
	result:=strings.Split(request, sep)
	logFile.LogDebug(dir,result)
	if len(result)==length {
		if(result[0]=="GET" || result[0]=="POST"){
			this.Method=result[0]
		}
		r, _ := regexp.Compile("^(/[A-Za-z0-9]+)+(\\?[A-Za-z0-9_\\-]+=[A-Za-z0-9_%\\-.\\+]+)?(&[A-Za-z0-9_\\-]+=[A-Za-z0-9_%/\\-\\.\\?&=:\\+]+)*$")
		logFile.LogDebug(dir,GetAlphaTab(result[1]))
		if r.MatchString(result[1]){
			this.Uri=result[1]
		}
		logFile.LogDebug(dir,"right:"+result[1])
		this.Version=result[2]

	}
	return this
}

func (this * Request)  ParseParam(dir string,request string,sep string) string {
	result:=strings.Split(request, sep)
	logFile.LogDebug(dir,"param result:")
	logFile.LogDebug(dir,result)
	if len(result)==paramLength {
		r, _ := regexp.Compile("^([A-Za-z0-9_\\-]+=[A-Za-z0-9_%\\-.\\+]+)?(&[A-Za-z0-9_\\-]+=[A-Za-z0-9_%\\-\\.\\?&=:\\+]+)*$")
		if r.MatchString(result[0]){
			return result[0]
		}
	}
	return  ""
}
func (this * Request) ParseTime(dir string,request string) int64  {
	logFile.LogDebug(dir,"\ntime start:")
	resultS:=strings.Split(request, timeSepE)
	logFile.LogDebug(dir,resultS)
	if(len(resultS)<2){
		return 0
	}
	logFile.LogDebug(dir,resultS[1])
	resultE:=strings.Split(resultS[0],timeSepS)
	logFile.LogDebug(dir,"\n finish  time start end:")
	logFile.LogDebug(dir,resultE[1])
	logFile.LogDebug(dir,resultE)
	if(len(resultE)<2){
		return 0
	}
        result:=string(resultE[1])
	logFile.LogDebug(dir,"\n\n result:")
	logFile.LogDebug(dir,len(result))

	if result==""{
		return 0
	}
	//11/May/2018:22:56:37 +0800
	logFile.LogDebug(dir,"\n "+result)
	t,err:=time.Parse("02/Jan/2006:15:04:05 +0800",result)
	logFile.LogDebug(dir,"\n time parse :")
	logFile.LogDebug(dir,err)
	if err!=nil{
	  return 0
	}
	logFile.LogDebug(dir,t)
	return t.Unix()
}

func (this * Request) ParseLine(dir string,line string) *Request {
	r:=new(Request)
	for _, v := range strings.Split(line, lineSep) {
		logFile.LogDebug(dir,v)
		request:=r.ParseUri(dir,v,uriSep)
		logFile.LogDebug(dir,request)
		if(request.Uri!=""){
			r.Uri=request.Uri
			if(request.Version!=""){
				r.Version=request.Version
			}
		}
		if(request.Method!=""){
			r.Method=request.Method
		}
		param:=r.ParseParam(dir,v,uriSep)
		logFile.LogDebug(dir,param)
		if(param!=""){
			r.Param=param
		}

	}
	time:=r.ParseTime(dir,line)
	if(time!=0){
		r.Time=time
	}
	return r
}

func (this * Request) ForMatRequest(dir string,req * Request,ipPort *config.IpPort,confParams config.Params)  *Request {
	req.Arguments=url.ToJson(req.Param)
	res:=strings.Split(req.Uri,"?")
	if res[0]!=""{
		req.Uri=res[0]
	}
	if len(res)>1&&res[1]!=""{
		for k,v:=range (url.ToJson(res[1])){
			req.Arguments[k]=v
		}
	}
	req.Url=config.BuildUrl(dir,*ipPort,req.Uri)
	for k,v:=range (confParams.Params){
		req.Arguments[k]=v
	}
	logFile.LogDebug(dir,req)
	return req
}

func (this * Request)Url2Json(req * Request) * Request {
	req.Arguments=url.ToJson(req.Param)
	res:=strings.Split(req.Uri,"?")
	if res[0]!=""{
		req.Uri=res[0]
	}
	if len(res)>1&&res[1]!=""{
		for k,v:=range (url.ToJson(res[1])){
			req.Arguments[k]=v
		}
	}
	return req
}
