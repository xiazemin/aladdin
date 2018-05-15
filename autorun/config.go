package autorun
import (
	jsonEx "github.com/xiazemin/aladdin/damon/json"
	"encoding/json"
	"github.com/xiazemin/aladdin/damon/logFile"
)
type WatchConfig struct {
	Path string `json:"path"`
}

func GetPath(dirLog string,dir string,filename string) string  {
	v:=new(WatchConfig)
	logFile.LogNotice(dirLog,dir+filename)
	datajson:=jsonEx.Load(dir+filename)
	logFile.LogDebug(dirLog,string(datajson))
	err := json.Unmarshal(datajson, v)
	if(err!=nil){
		logFile.LogWarnf(dirLog,err)
		return v.Path
	}
	logFile.LogNotice(dirLog,v)
	return  v.Path
}
