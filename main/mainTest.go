package main
import (
	"github.com/xiazemin/aladdin/damon/log"
	"github.com/xiazemin/aladdin/damon/logFile"
	"github.com/xiazemin/aladdin/damon/url"
	"github.com/xiazemin/aladdin/damon/json"
	"github.com/xiazemin/aladdin/damon/file"
	"github.com/xiazemin/aladdin/damon/config"
	"go/src/fmt"
	"github.com/xiazemin/aladdin/api"
	"strings"
)
/**
log_format main '$remote_addr - $remote_user [$time_local]  '
        '"$request" $status $request_length $request_time $body_bytes_sent '
        '"$http_referer" "$http_user_agent" $server_addr $upstream_addr $host $upstream_cache_status $HEADER "$request_body" $operationid';
 */
const defaultDir  ="/Users/didi/aladdin/xiazemin/10.96.76.97/2018-05-14-11/case1/"
const configData  = "config/data.json"
const globalConfig  ="globalConfig.json"
func main()  {
	fmt.Print(strings.Contains("price=91\u0026passe=113\u0026nt=1\u0026trav=34","\u0026"))
	uj:=new(api.UrlJson)
	fmt.Println(uj.Url2Json("price=91\u0026passe=113\u0026nt=1\u0026trav=34",defaultDir))
	return
	uri:="/www.baidu.com?rsv_enter=1&rqlang=cn&rsv_bp=0&rsv_sug3=5&rsv_idx=1&rsv_sug7=100&rsv_t=e3d0gUXTMejJdTYOwTOhMwTKhTmIVyTYRVM5gYbU%2FiiLObrnxUcZBLs74SQ&tn=baidu&rsv_sug4=4847&wd=tes&f=8&https://www.baidu.com/s?ie=utf-8&inputT=1058&rsv_pq=ff88e6380000669f&rsv_sug1=3&rsv_sug2=0"

	fmt.Println(uj.MatchUrl(uri,defaultDir))
	uri2:="abc=desf&"
	fmt.Println(uj.MatchUrl(uri2,defaultDir))

	r0:=new(log.Request)
	r1:=r0.ParseParam(defaultDir,uri," ")
	fmt.Print(r1)

	logFile.LogDebug(defaultDir,r1)
	at:=log.GetAlphaTab(uri)
	for k,_:=range at{
		if !strings.Contains("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789",k){
			logFile.LogDebug(defaultDir,k)
		}
		if strings.Contains("()$*+.[]?\\[]|^-",k){
			logFile.LogDebug(defaultDir,"@@"+k)
		}
	}

	logFile.LogDebug(defaultDir,at)


	userConf:=config.GetUserConf(defaultDir,globalConfig)
	fmt.Print(userConf)
	logFile.LogNotice(defaultDir,defaultDir)
	list:=config.GetSelectedLogFiles(defaultDir,configData,userConf.User,userConf.Date,userConf.Model)

	fmt.Print(list)
	logFile.LogDebug(defaultDir,log.GetAlphaTab("ie=utf-8&f=8&rsv_bp=1&rsv_idx=1&tn=baidu&wd=%20invalid%20character%20%27%7B%27%20after%20top-level%20value&oq=golang%2520%25E8%25B0%2583%25E7%2594%25A8%25E6%25A0%2588&rsv_pq=cd93e94d00051db9&rsv_t=1c5eYwmNwWfA31oioXIREysIkHFRAwl1xCgZlt79euRIPpkelmnkBE9uv4k&rqlang=cn&rsv_enter=0&inputT=483&rsv_sug3=721&rsv_sug2=0&rsv_sug4=483"))

	//[$time_local] "$request"
	line:=`
	10.0.0.0 - - [11/May/2018:22:55:35 +0800]  "POST /abc/def/v1/hgk/hhhhaave HTTP/1.1" 200 427 0.010 745 "-" "-" 10.0.0.0 127.0.0.1:9000 10.0.0.0 - 13jhgd "k_id=134255&p_id=456787989&version=1.0.0" 454434
	`
	r:=new(log.Request)
	logFile.LogDebug(defaultDir,r.ParseTime(defaultDir,line))
	logFile.LogDebug(defaultDir,url.ToJson("l_id=34544534&y_id=09879&version=1.0.0"))
	for key,value:=range (url.ToJson("lid=56454&h_id=456787&version=1.0.0")){
		if str, ok := value.(string);ok{
			logFile.LogDebug(defaultDir,key+"=>"+str)
		}

	}
	logFile.LogDebug(defaultDir,json.ToForm(defaultDir,url.ToJson("lid=4565&j_id=456756546&version=1.0.0")))

	allDirs:=file.ListDirs(defaultDir,defaultDir)
	logFile.LogDebug(defaultDir,allDirs)
	fmt.Print(allDirs)

	dir:=file.GetDir(defaultDir,".",defaultDir)
	fmt.Print(dir)
	fmt.Println("\n----------")
	str:=file.PrintDirs(dir,0)
	fmt.Println("\n----------")
	fmt.Println(str)

}
