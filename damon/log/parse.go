package log
import (
	"bufio"
	"io"
	"os"
	"github.com/xiazemin/aladdin/damon/logFile"
	"go/src/strings"
)

func Parse(dir string,fileName string,lineEnd byte) []*Request {
	var reqList[]*Request
	f, err := os.Open(dir+fileName)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	rd := bufio.NewReader(f)
	for {
		line, err := rd.ReadString(lineEnd) //以'\n'为结束符读入一行

		if err != nil || io.EOF == err {
			break
		}
		r:=new(Request)
		r=r.ParseLine(dir,line)
		logFile.LogDebug(dir,line)
		logFile.LogDebug(dir,r)
		if(r!=nil && r.Uri!=""){
			logFile.LogDebug(dir,"\n get uri and param:")
			logFile.LogDebug(dir,r.Uri)
			logFile.LogDebug(dir,r)
			reqList=append(reqList,r)
		}
	}
	return reqList
}

func ParseContent(content string,lienEnd string,logDir string)[]*Request{
	var reqList[]*Request
	lines:= strings.Split(content,lienEnd) //以'\n'为结束符读入一行
	logFile.LogDebug(logDir,lines)
	for _,line:=range lines{
		if line==""{
			continue
		}
		r:=new(Request)
		r=r.ParseLine(logDir,line)
		logFile.LogDebug(logDir,r)
		if(r!=nil && r.Uri!=""){
			logFile.LogDebug(logDir,"\n get uri and param:")
			logFile.LogDebug(logDir,r.Uri)
			logFile.LogDebug(logDir,r)
			reqList=append(reqList,r)
		}
	}
	return  reqList
}

