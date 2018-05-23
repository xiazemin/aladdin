package api

import (
	"net/http"
	"fmt"
	"github.com/xiazemin/aladdin/damon/logFile"
	"github.com/xiazemin/aladdin/damon/log"
	"strings"
	"github.com/xiazemin/aladdin/damon/url"
	"go/src/encoding/json"
	"github.com/xiazemin/aladdin/damon/file"
	jsonEx "github.com/xiazemin/aladdin/damon/json"
)

const lineEnd  = '\n'

type UrlInfo struct{
   Uri string
	Arguments map[string]interface{}
}
func (this*Data)Log2json(w http.ResponseWriter,r *http.Request,defaultDir string,logDir string,viewDir string,configParams string) string{
	r.ParseForm()
	fileName:=string(r.Form.Get("file_name"))
	logFile.LogNotice(logDir,fileName)
	reqL:=log.Parse(defaultDir,fileName,lineEnd)
	logFile.LogNotice(logDir,reqL)
	var ul []UrlInfo
	for _,req:=range reqL{
		fmt.Print(req)
		res:=strings.Split(req.Uri,"?")
		 var u UrlInfo
                u.Arguments=make(map[string]interface{})
		if res[0]!=""{
			u.Uri=res[0]
		}
		if len(res)>1&&res[1]!=""{
			for k,v:=range (url.ToJson(res[1])){
				u.Arguments[k]=v
				fmt.Println(k)
				fmt.Println(v)
			}
		}
		for k,v:=range req.Arguments{
			u.Arguments[k]=v
			fmt.Println(k)
			fmt.Println(v)
		}
		if len(u.Arguments)!=0 {
			ul = append(ul, u)
		}
	}
	logFile.LogNotice(logDir,ul)
	js,err:=json.Marshal(ul)
	if err !=nil{
		logFile.LogWarnf(logDir,err)
	}
	logFile.LogNotice(logDir,string(js))
	jsFileName:=strings.Replace(fileName,".log",".json",1)
	file.Write(defaultDir, jsFileName, string(js))
	fmt.Fprintln(w, string(jsFileName))
return fileName
}
func (this*Data)Log2conf(w http.ResponseWriter,r *http.Request,defaultDir string,logDir string,viewDir string,configParams string)string{
	r.ParseForm()
	fileName:=string(r.Form.Get("file_name"))
	logFile.LogNotice(logDir,fileName)

	raw:=string(jsonEx.Load(defaultDir+fileName))
	logFile.LogNotice(logDir,raw)
	res:=log.MatchJson(defaultDir,configParams,raw)
        logFile.LogNotice(logDir,res)

	jsConfName:=strings.Replace(fileName,".json","Conf.json",1)
	file.Write(defaultDir, jsConfName, string(res))
	fmt.Fprintln(w, jsConfName)
	return fileName
}
func (this*Data)ConfLoad(w http.ResponseWriter,r *http.Request,defaultDir string,logDir string,viewDir string,configParams string)string{
	r.ParseForm()
	fileName:=string(r.Form.Get("file_name"))
	logFile.LogNotice(logDir,fileName)
	raw:=string(jsonEx.Load(defaultDir+fileName))
	logFile.LogNotice(logDir,raw)
	fmt.Fprintln(w, raw)
	return fileName
}
func (this*Data)ConfEdit(w http.ResponseWriter,r *http.Request,defaultDir string,logDir string,viewDir string,configParams string)string{
	r.ParseForm()
	jsonData:=string(r.Form.Get("json_data"))
	logFile.LogNotice(logDir,jsonData)
	//fmt.Println(log.UpdateJson(defaultDir,raw,res))
	fmt.Fprintln(w, jsonData)
	return jsonData
}
