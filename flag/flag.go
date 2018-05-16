package flag

import (
	"flag"
	"github.com/xiazemin/aladdin/damon/logFile"
)

var ip = flag.String("ip", "127.0.0.1", "usage: -p=127.0.0.1(your ip)")
var  port = flag.Int("port", 8080, "usage: -port=8080(your port)")
var server=flag.Bool("s",true,"uasge:-s")

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