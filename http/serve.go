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
        DefaultFile string
}
func (this *Serv)IndexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "hello world")
}

func (this *Serv)AladdinHandler(w http.ResponseWriter, r *http.Request)  {
        //logFile.LogNotice(this.LogDir+"http/",r.Body)
	logFile.LogNotice(this.LogDir+"http/",this.LogDir+"http/")
	r.ParseForm()
        ip:=string(r.Form.Get("ip"))
	port,err:=strconv.Atoi(r.Form.Get("port"))
	logFile.LogNotice(this.LogDir+"http/",fmt.Sprintf(" ip :%s ,port: %d ,from request",ip,port))
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

        logFile.LogNotice(this.LogDir+"http/",fmt.Sprintf(" ip :%s ,port: %d ,from request",ipPort.Ip,ipPort.Port))
	ret:=damon.HandleReq(this.DefaultDir,ipPort,this.ConfigParams,this.LineEnd,this.DefaultFile)
	resp:="hello aladdin\n"+ret
	fmt.Fprintln(w, resp)
	logFile.LogNotice(this.LogDir+"http/",resp)
}

func (this *Serv)Serve(dir string,globalConfig string,configParams string,lineEnd byte,defaultFile string) {
	this.LogDir=dir
	this.DefaultDir=dir
	this.GlobalConfig=globalConfig
	this.ConfigParams=configParams
	this.LineEnd=lineEnd
	this.DefaultFile=defaultFile
	http.HandleFunc("/", this.IndexHandler)
	http.HandleFunc("/aladdin", this.AladdinHandler)
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
