package http
import (
	"fmt"
	"net/http"
	"github.com/xiazemin/aladdin/damon/logFile"
	"github.com/xiazemin/aladdin/flag"
	"github.com/xiazemin/aladdin/damon"
	"github.com/xiazemin/aladdin/damon/config"
	"strconv"
)

type Serv struct {
	LogDir string
	DefaultDir string
	GlobalConfig string
        ConfigParams string
        LineEnd byte
        LogFiles []string
	ConfigData string
}
func (this *Serv)IndexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "hello world")
}

func (this *Serv)getIpPort(r *http.Request) *config.IpPort {
	ip:=string(r.Form.Get("ip"))
	port,err:=strconv.Atoi(r.Form.Get("port"))
	logFile.LogNotice(this.LogDir,fmt.Sprintf(" ip :%s ,port: %d ,from request",ip,port))
	if err!=nil{
		logFile.LogWarnf(this.LogDir,err)
	}
	var ipPort *config.IpPort
	ipPort=new(config.IpPort)
	if ip==""{
		ipPort=config.GetIpPort(this.DefaultDir,this.GlobalConfig)
		logFile.LogWarnf(this.LogDir," test service ip is empty,use ip in globalConfig.json instead ")
	}
	ipPort.Ip=ip
	ipPort.Port=port
	if port==0{
		ipPort.Port=8080
		logFile.LogWarnf(this.LogDir," port is 0 use 8080 instead ")
	}

	logFile.LogNotice(this.LogDir,fmt.Sprintf(" ip :%s ,port: %d ,from request",ipPort.Ip,ipPort.Port))
	return ipPort
}

func (this *Serv)getLogFiles(r *http.Request) []string {
	userConf:=config.GetUserConf(this.DefaultDir,this.GlobalConfig)
	logFile.LogNotice(this.LogDir,userConf)
	user:=string(r.Form.Get("user"))
	date:=string(r.Form.Get("date"))
	model:=string(r.Form.Get("model"))
	if user!=""{
		userConf.User=user
	}
	if date!=""{
		userConf.Date=date
	}
	if model!=""{
		userConf.Model=model
	}
	logFile.LogNotice(this.LogDir,userConf)
	logFiles:=config.GetSelectedLogFiles(this.DefaultDir,this.ConfigData,userConf.User,userConf.Date,userConf.Model)
	logFile.LogNotice(this.LogDir,logFiles)
	return logFiles
}

func (this *Serv)AladdinHandler(w http.ResponseWriter, r *http.Request)  {
        //logFile.LogNotice(this.LogDir+"http/",r.Body)
	logFile.LogNotice(this.LogDir,this.LogDir)
	r.ParseForm()
        ipPort:=this.getIpPort(r)
	logFiles:=this.getLogFiles(r)
	if len(logFiles)>0{
		this.LogFiles=logFiles
	}
	ret:=damon.HandleReq(this.DefaultDir,ipPort,this.ConfigParams,this.LineEnd,this.LogFiles)
	resp:="hello aladdin\n"+ret
	fmt.Fprintln(w, resp)
	logFile.LogNotice(this.LogDir,resp)
}

func (this *Serv)Serve(dir string,globalConfig string,configParams string,lineEnd byte,defaultFile  string,configData string) {
	this.LogDir=dir+"http/"
	this.DefaultDir=dir
	this.GlobalConfig=globalConfig
	this.ConfigParams=configParams
	this.LineEnd=lineEnd
	this.LogFiles=append(this.LogFiles,defaultFile)
	this.ConfigData=configData
	http.HandleFunc("/", this.IndexHandler)
	http.HandleFunc("/aladdin", this.AladdinHandler)
	http.HandleFunc("/file/",this.FileRoute)
	// 注意斜杠！
	http.Handle("/download/", http.StripPrefix("/download/", http.FileServer(http.Dir(dir+"download/"))))
	ip,port:=flag.GetIpPort(dir)
	if(ip==""){
		err:=http.ListenAndServe("127.0.0.1:8088", nil)
		if err!=nil{
			logFile.LogWarnf(dir,err)
		}
	}else{
		if port==0{
			port=8088
		}
		err:=http.ListenAndServe(fmt.Sprintf("%s:%d",ip,port), nil)
		if err!=nil{
			logFile.LogWarnf(dir,err)
		}

	}


}
