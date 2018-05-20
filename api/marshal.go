package api
import(
	"encoding/json"
	"github.com/xiazemin/aladdin/damon/logFile"
)
func JsonPretty(js string,dir string)string  {
	var v interface{}
	logFile.LogNotice(dir,js)
	err:=json.Unmarshal([]byte(js),&v)
	if err!=nil{
		logFile.LogWarnf(dir,err)
	}
	logFile.LogNotice(dir,v)
	data, err := json.MarshalIndent(v, "", "    ")
	if err != nil {
		logFile.LogWarnf(dir,err)
	}
	logFile.LogNotice(dir,data)
	//fmt.Printf("%s\n", data)
	return string(data)
}