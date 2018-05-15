package config

import (
	"encoding/json"
	jsonEx "github.com/xiazemin/aladdin/damon/json"
	"github.com/xiazemin/aladdin/damon/logFile"
	"github.com/xiazemin/aladdin/damon/response"
	"github.com/xiazemin/aladdin/damon/file"
	"reflect"
)
type Params struct {
	Name string `json:"name"`
	Params map[string] interface{} `json:"params"`
}

func LocadParams(dir string,fileName string) *Params{
	datajson:=jsonEx.Load(dir+fileName)
	data:=[]byte("{\"name\":\""+fileName+"\",\"params\":"+string(datajson)+"}")
	logFile.LogDebug(dir,"{\"name\":\""+fileName+"\",\"params\":"+string(datajson)+"}")
	p:=new(Params)
	err:=json.Unmarshal(data,p)
	if err!=nil{
		logFile.LogWarnf(dir,err)
	}
	return p
}

func UpdateParams(dir string,fileName string,result response.Response){
	datajson:=map[string] interface{}(LocadParams(dir,fileName).Params)
	logFile.LogNotice(dir,datajson)
	for k,_:=range (datajson){
		for kn,vn:=range(result.Result){
			if k==kn && !isEmpty(kn) {
				datajson[k]=vn
				logFile.LogNotice(dir,k)
			}
		}
	}
	logFile.LogNotice(dir,result)
	logFile.LogNotice(dir,datajson)
	res,err:=json.Marshal(datajson)
	if err!=nil{
		logFile.LogWarnf(dir,err)
	}else {
		file.Write(dir, fileName, string(res))
	}
}

func isEmpty(a interface{}) bool {
	v := reflect.ValueOf(a)
	if v.Kind() == reflect.Ptr {
		v=v.Elem()
	}
	return v.Interface() == reflect.Zero(v.Type()).Interface()
}

