package config

import (
	"encoding/json"
	jsonEx "github.com/xiazemin/aladdin/damon/json"
	"github.com/xiazemin/aladdin/damon/logFile"
	"github.com/xiazemin/aladdin/damon/file"
	"reflect"
)
type Params struct {
	Name string `json:"name"`
	Params map[string] interface{} `json:"params"`
}

func LocadParams(dir string,fileName string) *Params{
	datajson:=jsonEx.Load(dir+fileName)

	p:=new(Params)
	p.Name=fileName
	var v  map[string] interface{}
	err:=json.Unmarshal(datajson,&v)
	if err!=nil{
		logFile.LogWarnf(dir,err)
	}
	p.Params=v
	logFile.LogNotice(dir,*p)
	logFile.LogDebug(dir,string(datajson))
	return p
}

func UpdateParams(dir string,fileName string,param map[string]interface{}){
	datajson:=map[string] interface{}(LocadParams(dir,fileName).Params)
	logFile.LogNotice(dir,datajson)
	for k,_:=range (datajson){
		for kn,vn:=range(param){
			if k==kn && !isEmpty(kn) {
				datajson[k]=vn
				logFile.LogNotice(dir,k)
			}
		}
	}
	logFile.LogNotice(dir,param)
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

