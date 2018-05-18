package http

import (
	"net/http"
	"fmt"
	"go/src/strings"
	"github.com/xiazemin/aladdin/damon/logFile"
	"github.com/xiazemin/aladdin/damon/file"
	"github.com/xiazemin/aladdin/api"
	"github.com/xiazemin/aladdin/flag"
	"text/template"
)

func (this *Serv)FileRoute(w http.ResponseWriter, r *http.Request)  {
	var resp string
	uris:=strings.Split(r.RequestURI,"/")
	tmplDir:=flag.GetTmplDir()
	if(len(uris)<3){
		logFile.LogNotice(this.LogDir,r.RequestURI+fmt.Sprintf("   %d  %+v",len(uris),uris))
		fmt.Fprintln(w, r.RequestURI+fmt.Sprintf("   %d  %+v",len(uris),uris))
		return
	}
	switch uris[2] {
	case "add":
		resp=this.Add(r)
	case "update":
		resp=this.Update(r)
	case "get":
		resp=this.Get(r)
	case "data":
		if len(uris)<4{
			resp=r.RequestURI+fmt.Sprintf("   %d  %+v  %s ",len(uris),uris,uris[2])
		}else{
			data:=new(api.Data)
			resp=data.Handle(uris,w,r,this.DefaultDir,this.LogDir,tmplDir,this.ConfigData)
			logFile.LogNotice(this.LogDir,resp)
			return
		}
	case "config":
		if len(uris)<4{
			resp=r.RequestURI+fmt.Sprintf("   %d  %+v  %s ",len(uris),uris,uris[2])
		}else{
			conf:=new(api.Config)
			name,templ,value:=conf.Handle(uris,r,this.DefaultDir,this.LogDir,tmplDir,this.ConfigData)
			logFile.LogNotice(this.LogDir,fmt.Sprintf("\n conf.Handel result:%s\n%s\n%+v\n",name,templ,value))
			t := template.New(name)
			t.Parse(templ)
			t.Execute(w, value)
			return
		}
	default:
		resp=r.RequestURI+fmt.Sprintf("   %d  %+v  %s ",len(uris),uris,uris[2])

	}
	logFile.LogNotice(this.LogDir,r.RequestURI+fmt.Sprintf("   %d  %+v  %s   %s",len(uris),uris,uris[2],resp))
	fmt.Fprintln(w,resp)
	return
}

func (this *Serv)Add( r *http.Request) string {
	return  string(r.RequestURI)
}

func (this *Serv)Update( r *http.Request)string {
	return  string(r.RequestURI)
}

func (this *Serv)Get( r *http.Request) string  {
	dir:=file.GetDir(this.DefaultDir,".",this.LogDir)
	str:=file.GetPrintDirs(dir,0)
	return  string(r.RequestURI+"\n"+str)
}