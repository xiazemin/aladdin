package api

import (
	"net/http"
	"fmt"
	"github.com/xiazemin/aladdin/damon/logFile"
	"encoding/json"
	"github.com/xiazemin/aladdin/damon/url"
	jsonEx "github.com/xiazemin/aladdin/damon/json"
	"io/ioutil"
	"text/template"
	"github.com/xiazemin/aladdin/damon/netenv"
	"github.com/xiazemin/aladdin/damon/log"
	"regexp"
	"strings"
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
	case "urlEncode":
		resp=UrlEncode(content)
		logFile.LogNotice(logDir,"\033[32minput\033[0m:"+content+",\033[31murlEncode\033[0m:"+resp)
		fmt.Fprintln(w,resp)
	case "urlDecode":
		resp=UrlDecode(content)
		logFile.LogNotice(logDir,"\033[32minput\033[0m:"+content+",\033[31murlDecode\033[0m:"+resp)
		fmt.Fprintln(w,resp)
	case "jsonPretty":
		resp=JsonPretty(content,logDir)
		logFile.LogNotice(logDir,"\033[32minput\033[0m:"+content+",\033[31mjsonPretty\033[0m:"+resp)
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
	logFile.LogNotice(logDir,u)
	logFile.LogDebug(logDir,strings.Contains(u,"\u0026"))
	logFile.LogDebug(logDir,strings.Contains(u,"\\u0026"))
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

const lienEnd  ="\n"
const sepu="\\u0026"
const sep  ="&"
func (this*UrlJson)ParseUrl(raw string,logDir string) string{
       var urls [] Url
	lines:= strings.Split(raw,lienEnd) //以'\n'为结束符读入一行
	logFile.LogDebug(logDir,lines)
	for i,line:=range lines{
            if line==""{
		    continue
	    }
		var url Url
		url.Id = i
		if strings.Contains(line,sepu){
			line=strings.Replace(line,sepu,sep,-1)
		}
		urlStr:=this.MatchUrl(line,logDir)
		if urlStr==""{
			r:=new(log.Request)
			r=r.ParseLine(logDir,line)
			r=r.Url2Json(r)
			url.Url=r.Url
			url.Arguments=jsonEx.ToForm(logDir,r.Arguments)
		}else {
			res := strings.Split(urlStr, "?")
			r, _ := regexp.Compile("^http[s]://[A-Za-z0-9]+\\.[A-Za-z0-9]+[/A-Za-z0-9\\.]*$")
			if r.Match([]byte(res[0])) {
				url.Url = res[0]
				url.Arguments = strings.Replace(urlStr, res[0] + "?", "", 1)
			} else {
				url.Arguments = urlStr
			}
		}
		urls=append(urls,url)
	}
	r,err:=json.Marshal(urls)
	if err!=nil{
		logFile.LogWarnf(logDir,err)
	}
	return string(r)
}

func (this*UrlJson)MatchUrl(raw string,logDir string) string {
	rexs:=[]string{"^http[s]://[A-Za-z0-9]+\\.[A-Za-z0-9]+[/=\\?%\\-&_~`@\\[\\]\\':+!]*([^<>\"\"])*$",
		"^/[A-Za-z0-9]+\\.[A-Za-z0-9]+[/=\\?%\\-&_~`@\\[\\]\\':+!]*([^<>\"\"])*$",
		"^%\\-&_~`@\\[\\]\\':+!]*([^<>\"\"])*$",
		"^([A-Za-z0-9_\\-]+=[A-Za-z0-9_%\\-.\\+]+)?(&[A-Za-z0-9_\\-\\]+=[A-Za-z0-9_%\\-\\.\\?&=:\\+]+)*$"}
	for _,rex:=range rexs {
		r, _ := regexp.Compile(rex)
		logFile.LogNotice(logDir, raw)
		logFile.LogDebug(logDir, r)
		if (r.MatchString(raw)) {
			logFile.LogNotice(logDir,rex)
			return raw
		}
	}
      return  ""
}