package api

import (
	"net/http"
	"fmt"
	"github.com/xiazemin/aladdin/damon/logFile"
	"github.com/xiazemin/aladdin/damon/log"
	"encoding/json"
	"github.com/xiazemin/aladdin/damon/url"
	jsonEx "github.com/xiazemin/aladdin/damon/json"
	"io/ioutil"
	"text/template"
	"github.com/xiazemin/aladdin/damon/netenv"
)
type UrlJson struct {

}

func (this *UrlJson)Handle(uris []string,w http.ResponseWriter,r *http.Request,defaultDir string,logDir string,viewDir string,configData string) string{
	r.ParseForm()
	logFile.LogNotice(logDir,r.Form)
	content:=string(r.Form.Get("content"))
	logFile.LogNotice(logDir,content)
	var resp string
	switch uris[3] {
	case "url2json":
		resp=this.Url2Json(content,logDir)
		logFile.LogNotice(logDir,"\033[32minput\033[0m:"+content+",\033[31murl2json\033[0m:"+resp)
		fmt.Fprintln(w,resp)
	case "json2url":
		resp=this.Json2Url(content,logDir)
		logFile.LogNotice(logDir,"\033[32minput\033[0m:"+content+",\033[31mjson2url\033[0m:"+resp)
		fmt.Fprintln(w,resp)
	case "parseurl":
             resp=this.ParseUrl(content,logDir)
		logFile.LogNotice(logDir,"\033[32minput\033[0m:"+content+",\033[31mparseurl\033[0m:"+resp)
		fmt.Fprintln(w,resp)
	default:
		templ, _ := ioutil.ReadFile(viewDir+"urljson/"+"urljson.html")
		t := template.New("parse log file ")
		t.Parse(string(templ))
		ip:=netenv.GetLocalIp(logDir)
		url:="http://"+ip+":8088"
		t.Execute(w, url)
		resp=r.RequestURI+fmt.Sprintf("   %d  %+v  %s  %s %s",len(uris),uris,uris[2],uris[3],content)
	}
return resp
}

func (this *UrlJson)Url2Json(u string,logDir string)string{
	 m:=url.ToJson(u)
	j,err:=json.Marshal(m)
	if err!=nil{
		logFile.LogWarnf(logDir,err)
	}
	return  string(j)
}

func (this *UrlJson)Json2Url(js string,logDir string)string{
        var i map[string]interface{}
	err:=json.Unmarshal([]byte(js),&i)
	if err!=nil{
		logFile.LogWarnf(logDir,err)
		return ""
	}

	return jsonEx.ToForm(logDir,i)
}

type Url struct {
	Id int `json:"id"`
	Url string `json:"url"`
	Arguments interface{} `json:"arguments"`
} 
func (this*UrlJson)ParseUrl(raw string,logDir string) string{
       var urls [] Url
	reqList:=log.ParseContent(raw,"\n",logDir)
	logFile.LogNotice(logDir,fmt.Sprintf("\n leng of req:%d=>%+v\n",len(reqList),reqList))
	for i,req:=range reqList{
		req=req.Url2Json(req)
		var u Url
		u.Id=i
		u.Url=req.Uri
		u.Arguments=req.Arguments
		urls=append(urls,u)
	}
	r,err:=json.Marshal(urls)
	if err!=nil{
		logFile.LogWarnf(logDir,err)
	}
	return string(r)
}