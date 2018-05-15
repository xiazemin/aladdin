package log
import (
	"bufio"
	"io"
	"os"
	"github.com/xiazemin/aladdin/damon/logFile"
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

