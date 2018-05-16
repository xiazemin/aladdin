package logFile
import (
	"log"
	"os"
	io "io/ioutil"
	"fmt"
	"encoding/json"
	"runtime/debug"
)

type LogConfig struct {
	LogLevel int `json:"log_level"`
}
const logWarnf  = "warn.log" //1
const logNotice  ="notice.log" //2
const logDebug  ="debug.log"  //4
const logConfig  ="logConfig.json"

func Load(filename string) []byte {
	data, err := io.ReadFile(filename)
	if err != nil{
		return nil
	}
	datajson := []byte(data)
	return datajson
}

func GetLogLevel (filename string) (*LogConfig,error)  {
	v:=new(LogConfig)
	datajson:=Load(filename)
	//fmt.Print(string(datajson))
	err := json.Unmarshal(datajson, v)
	return  v,err
}

func SetLogLevel(filename string){
	_, err1 := os.Stat(filename)
	var (
		logConf *os.File
		err error
	)
	if err1!=nil{
		logConf,err = os.Create(filename)

	}else {
		logConf,err=os.OpenFile(filename,os.O_RDWR|os.O_CREATE|os.O_APPEND,0644)
	}
	defer logConf.Close()
	if err != nil {
		log.Fatalln("open file error !")
	}
	n3, err := logConf.WriteString("{\"log_level\":7}")
	fmt.Println(fmt.Sprintf("\n%d words write to logConfig.json",n3))

}

func log2file(dir string,level string,prefix string,content interface{},trace string){
	errDir := os.MkdirAll(dir+"log/", 0777)
	if errDir != nil {
		fmt.Printf("%s", errDir)
	}
	// 定义一个文件
	fileName := dir+"log/"+level
	_, err1 := os.Stat(fileName)
	var (
		logFile *os.File
	        err error
	)
	if err1!=nil{
		logFile,err = os.Create(fileName)

	}else {
		logFile,err=os.OpenFile(fileName,os.O_RDWR|os.O_CREATE|os.O_APPEND,0644)
	}
	defer logFile.Close()
	if err != nil {
		log.Fatalln("open file error !")
	}
	commlog := log.New(logFile,prefix,log.LstdFlags)
	commlog.Println(fmt.Sprintf("%+v",content)+trace)
}

func LogDebug(dir string,content interface{}){
	logLevel,err:=GetLogLevel(dir+logConfig)
	if(err!=nil){
		fmt.Print(err)
		logLevel.LogLevel=7
		SetLogLevel(dir+logConfig)
	}
	//fmt.Println(fmt.Sprintf("%d",logLevel.LogLevel))
	if logLevel.LogLevel&4==0{
		return
	}
	//// 创建一个日志对象
	//debugLog := log.New(logFile,"[Debug]",log.LstdFlags)
	//debugLog.Println("A debug message here")
	////配置一个日志格式的前缀
	//debugLog.SetPrefix("[Info]")
	//debugLog.Println("A Info Message here ")
	////配置log的Flag参数
	//debugLog.SetFlags(debugLog.Flags() | log.LstdFlags)
	//debugLog.Println("A different prefix")
	log2file(dir,logDebug,"[Debug]",content,fmt.Sprintf("%s", debug.Stack()))
}

func LogNotice(dir string,content interface{})  {
	errDir := os.MkdirAll(dir+"log/", 0777)
	if errDir != nil {
		fmt.Printf("%s", errDir)
		fmt.Println(fmt.Sprintf("%s", debug.Stack()))
	}
	logLevel,err:=GetLogLevel(dir+logConfig)
	if(err!=nil){
		fmt.Print(err)
		logLevel.LogLevel=7
		SetLogLevel(dir+logConfig)
	}
	if logLevel.LogLevel&2==0{
		return
	}
	log2file(dir,logNotice,"[Notice]",content,"")
}

func LogWarnf(dir string,content interface{})  {
	logLevel,err:=GetLogLevel(dir+logConfig)
	if(err!=nil){
		fmt.Print(err)
		logLevel.LogLevel=7
		SetLogLevel(dir+logConfig)
	}
	if logLevel.LogLevel&1==0{
		return
	}
	log2file(dir,logWarnf,"[Warning]",content,fmt.Sprintf("%s", debug.Stack()))
}