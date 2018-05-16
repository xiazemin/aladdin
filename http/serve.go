package http
import (
	"fmt"
	"net/http"
	"github.com/xiazemin/aladdin/damon/logFile"
	"github.com/xiazemin/aladdin/flag"
	"github.com/xiazemin/aladdin/damon"
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
	ret:=damon.HandleReq(this.DefaultDir,this.GlobalConfig,this.ConfigParams,this.LineEnd,this.DefaultFile)
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
	ip,port:=flag.GetIpPort(dir)
	if(ip==""){
		err:=http.ListenAndServe("127.0.0.1:8080", nil)
		if err!=nil{
			logFile.LogWarnf(dir,err)
		}
	}else{
		if port==0{
			port=8080
		}
		err:=http.ListenAndServe(fmt.Sprintf("%s:%d",ip,port), nil)
		if err!=nil{
			logFile.LogWarnf(dir,err)
		}

	}


}
