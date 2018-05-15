package response

import (
	"encoding/json"
	"github.com/xiazemin/aladdin/damon/logFile"
)
type Response struct {
	Errno int `json:"errno,string"`
	ErrMsg string `json:"errmsg"`
	Result map[string] interface{} `json:"result"`
}
func Parse(dir string,result string)*Response{
        v:=new(Response)
	err:=json.Unmarshal([]byte(result),v)
	if(err!=nil){
		logFile.LogWarnf(dir,err)
		logFile.LogWarnf(dir,result)
	}
	return v

}
