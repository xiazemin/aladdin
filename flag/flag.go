package flag

import (
	"flag"
	"github.com/xiazemin/aladdin/damon/logFile"
	"github.com/xiazemin/aladdin/damon/file"
	"go/src/fmt"
	"os"
)

var ip = flag.String("ip", "127.0.0.1", "usage: -p=127.0.0.1(your ip)")
var  port = flag.Int("port", 8088, "usage: -port=8080(your port)")
var server=flag.Bool("s",true,"uasge:-s (use http server or use standalone exe)")
var dirType=flag.Int("dirType",3,"usage:-dirType=1,2,3,...(pwd,userdefined,dev,~)")
var dir=flag.String("dir","~","usage:-dir=/home/xxx(userdefined work dir)")
var remoteSer  =flag.Bool("r",false,"usage:-r (use remote server or not )")
func GetIpPort(dir string) (string,int) {
	flag.Parse()
	logFile.LogNotice(dir,*ip)
	logFile.LogNotice(dir,*port)
	return *ip,*port
}

func IsServerType(dir string)bool{
	flag.Parse()
	logFile.LogNotice(dir,*server)
	return *server
}

func IsRemoteServer(dir string)bool  {
	flag.Parse()
	logFile.LogNotice(dir,*remoteSer)
	return *remoteSer
}

const defaultDir  ="/Users/didi/aladdin/xiazemin/10.96.76.97/2018-05-14-11/case1/"

func GetDefaultDir()string  {
	flag.Parse()
	var dirDef string
	if *dirType==1{
		dirDef, _ = os.Getwd()
	}else if(*dirType==2){
                dirDef=*dir
	}else if(*dirType==3){
		dirDef=defaultDir
	}else{
		var err error
		dirDef,err=file.Home()
		if err!=nil{
			fmt.Println(err)
		}
	}
	dirDef=AddSlash(dirDef)
	fmt.Println(fmt.Sprintf("dir Type:%d",*dirType))
	fmt.Println("work dir:"+dirDef)
       return dirDef
}

func AddSlash(dirDef string)  string{
	dirBytes:=[]byte(dirDef)
	if last:=dirBytes[len(dirDef)-1];last!='/'{
		dirDef=dirDef+"/"
	}
	return dirDef
}