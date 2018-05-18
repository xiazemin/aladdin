package netenv
import (
	"net"
	"github.com/xiazemin/aladdin/damon/logFile"
	"github.com/xiazemin/aladdin/flag"
	"go/src/fmt"
)
func GetLocalIp(dir string)string  {
	fmt.Print(flag.IsRemoteServer(dir))
	if !flag.IsRemoteServer(dir){
		return "127.0.0.1"
	}
	var ip string
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		logFile.LogWarnf(dir,err)
		return ip
	}
	logFile.LogNotice(dir,addrs)
	for _, address := range addrs {
		// 检查ip地址判断是否回环地址
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				logFile.LogNotice(dir,ipnet.IP.String())
				ip=ipnet.IP.String()
				break;
			}
		}
	}
	return ip
}
