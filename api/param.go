package api

import (
	"net/http"
	"github.com/xiazemin/aladdin/damon/file"
	"strings"
	"io/ioutil"
	"text/template"
	"github.com/xiazemin/aladdin/damon/logFile"
	"github.com/xiazemin/aladdin/damon/config"
)

type Param struct {
	DirInfo string
	ParamInfo map[string]interface{}
}
func (this *Data)Edit(w http.ResponseWriter,r *http.Request,defaultDir string,logDir string,viewDir string,configParams string)string{
	dir:=file.GetDir(defaultDir+"config/","config/",logDir)
	str:=file.GetPrintDirs(dir,0)
	str=strings.Replace(strings.Replace(str,"\n","<br/>",-1),
		" ","&nbsp;",-1)
        logFile.LogNotice(logDir,str)
	logFile.LogDebug(logDir,logDir)
	logFile.LogDebug(logDir,str)

	p:=new(Param)
	p.DirInfo=str
	cp:=config.LocadParams(defaultDir,configParams)//configData里包含"config/"
	logFile.LogNotice(logDir,cp)
	p.ParamInfo=cp.Params
	logFile.LogDebug(logDir,cp)
	logFile.LogDebug(logDir,p)

	save:=string(r.PostFormValue("submit"))
	if save=="更新" {
		//str=fmt.Sprintf("%s\n%+v",str,configList)
		var newP=make(map[string]interface{},0)
		for key,_:=range p.ParamInfo{
			newP[key]=string(r.PostFormValue(key))
		}
		logFile.LogNotice(logDir,p.ParamInfo)
		logFile.LogNotice(logDir,newP)
		p.ParamInfo=newP
		if p.ParamInfo==nil{
			logFile.LogWarnf(logDir,p)
		}else {
			config.UpdateParams(defaultDir, configParams, p.ParamInfo)
		}
	}

	logFile.LogDebug(logDir,viewDir+"param/"+"edit.html")
	templ, _ := ioutil.ReadFile(viewDir+"param/"+"edit.html")
	logFile.LogNotice(logDir,templ)
	t := template.New("edit config param file ")
	t.Parse(string(templ))
	t.Execute(w, p)
	return viewDir+"\n"+str
}
