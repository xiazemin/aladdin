package config

import (
	jsonEx "github.com/xiazemin/aladdin/damon/json"
	"encoding/json"
	"github.com/xiazemin/aladdin/damon/logFile"
)

type Description struct {
	User string `json:"user"`
	Date string `json:"date"`
	Model string `json:"model"`
	LogName string `json:"log_name"`
	Selected bool `json:"selected,string"`
	Description string `json:"description"`
//[{"user":"xiazemin","date":"2018051712","model":"model name","log_name":"path/raw.log","selected":1,"description":"model test case1"}]
}

func LoadLogDataDes(dirBase string,fileName string)[]Description  {
	var des []Description
	datajson:=jsonEx.Load(dirBase+fileName)
	logFile.LogNotice(dirBase,datajson)
	err := json.Unmarshal(datajson, &des)
	if(err!=nil){
		logFile.LogWarnf(dirBase,err)
	}
	logFile.LogNotice(dirBase,des)
	return des
}

func GetSelectedLogFiles(dirBase string,fileName string,user string,date string,model string)[]string  {
	desList:=LoadLogDataDes(dirBase,fileName)
	logFile.LogNotice(dirBase,desList)
	var selectedLog []string
	for _,des:=range desList{
		if des.Selected && des.User==user && des.Date==date && model==des.Model{
			selectedLog=append(selectedLog,des.LogName)
		}
	}
	logFile.LogNotice(dirBase,selectedLog)
	return selectedLog
}

