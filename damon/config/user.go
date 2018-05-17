package config

import (
	"github.com/xiazemin/aladdin/damon/logFile"
	jsonEx "github.com/xiazemin/aladdin/damon/json"
	"encoding/json"
)

type User struct {
	User string `json:"user"`
	Date string `json:"date"`
	Model string `json:"model"`
}
func GetUserConf(dir string,fileName string) *User {
	v:=new(User)
	datajson:=jsonEx.Load(dir+fileName)
	logFile.LogDebug(dir,dir)
	logFile.LogDebug(dir,string(datajson))
	err := json.Unmarshal(datajson, v)
	if(err!=nil){
		logFile.LogWarnf(dir,err)
	}
	logFile.LogNotice(dir,v)
	return  v
}
