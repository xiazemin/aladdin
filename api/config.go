package api

import (
	"net/http"
	"fmt"
	"github.com/xiazemin/aladdin/damon/file"
	"github.com/xiazemin/aladdin/damon/config"
	"github.com/xiazemin/aladdin/damon/logFile"
	"io/ioutil"
	"strings"
)

type Config struct {
    ConfigDir string
}
func (this*Config)Handle(uris []string,r *http.Request,defaultDir string,logDir string,viewDir string,configData string)(name string,templ string,value interface{})  {
	switch uris[3] {
	case "get":
		name,templ,value=this.Get(r,defaultDir,logDir,viewDir,configData)
	case "update":
		name,templ,value=this.Update(r,defaultDir,logDir,viewDir,configData)
	case "add":
		name,templ,value=this.Add(r,defaultDir,logDir,viewDir,configData)
	default:
		templ=r.RequestURI+fmt.Sprintf("   %d  %+v  %s  %s",len(uris),uris,uris[2],uris[3])

	}
	return  name,templ,value
}

type ConfigData struct {
	FileInfo string
	FileContent []config.Description
}

func (this *Config)Get(r *http.Request,defaultDir string,logDir string,viewDir string,configData string)( string, string, interface{}){
	this.ConfigDir=defaultDir+"config/"
	dir:=file.GetDir(this.ConfigDir,"config/",logDir)
	str:=file.GetPrintDirs(dir,0)

	conf:=new(ConfigData)
	conf.FileInfo=strings.Replace(strings.Replace(str,"\n","<br/>",-1),
		" ","&nbsp;",-1)

	logFile.LogNotice(logDir,this.ConfigDir+configData)
	configList:=config.LoadLogDataDes(defaultDir,configData) //configData里包含"config/"
	//str=fmt.Sprintf("%s\n%+v",str,configList)
	conf.FileContent=configList

	bytes, _ := ioutil.ReadFile(viewDir+"config/"+"get.html")
	fmt.Println(viewDir+"config/"+"get.html")
	logFile.LogNotice(logDir,string(viewDir+"config/"+"get.html"))
       logFile.LogNotice(logDir,string(bytes))
	return "get config",string(bytes),conf
}

func (this *Config)Update(r *http.Request,defaultDir string,logDir string,viewDir string,configData string)(name string,templ string,value interface{}){
	return defaultDir,viewDir,configData
}

func (this *Config)Add(r *http.Request,defaultDir string,logDir string,viewDir string,configData string)(name string,templ string,value interface{}){
	return defaultDir,viewDir,configData
}
