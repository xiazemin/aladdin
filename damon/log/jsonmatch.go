package log

import (
	"github.com/xiazemin/aladdin/damon/config"
	"go/src/encoding/json"
	"github.com/xiazemin/aladdin/damon/logFile"
	"go/src/fmt"
)

func UpdateJson(logDir string,datajson string,param string)string {
	logFile.LogNotice(logDir,datajson)
	logFile.LogNotice(logDir,param)
	var a interface{}
	var b interface{}
	if err:=json.Unmarshal([]byte(datajson),&a) ; err != nil {
	   	logFile.LogWarnf(logDir,err)
	}
	if err:=json.Unmarshal([]byte(datajson),&b) ; err != nil {
		logFile.LogWarnf(logDir,err)
	}
	c:=Merge(logDir,a,b)
	logFile.LogNotice(logDir,c)
	r,err:=json.Marshal(c)
	if err !=nil{
		logFile.LogWarnf(logDir,err)
	}
	logFile.LogNotice(logDir,r)
	return  string(r)
}


func Merge(logDir string,a interface{},b interface{})interface {} {
	if lv,ok:=a.([]interface{});ok{
		lvb,okb:=b.([]interface{})
		if !okb{
			logFile.LogWarnf(logDir," a is list b is not ")
		}
		return MergeList(logDir,lv,lvb)
	}else if mv,ok:=a.(map[string]interface{});ok{
		mvb,okb:=a.(map[string]interface{})
		if !okb{
			logFile.LogWarnf(logDir," a is map b is not ")
		}
		return MergeMap(logDir,mv,mvb)
	}
	return a
}

func MergeList(logDir string,a[]interface {},b []interface {})[]interface {}  {
	if len(a)!=len(b){
	    logFile.LogWarnf(logDir,fmt.Sprintf(" list length diff %d => %d ",len(a),len(b)))
		return  nil
	}
	for i:=0;i<len(a);i++ {
            a[i]=Merge(logDir,a[i],b[i])
	}
	return a
}

func MergeMap(logDir string,a map[string]interface {},b map[string]interface {})map[string]interface {}  {
	keys:=make([]string,0)
	for k,_:=range b{
		keys=append(keys,k)
	}

	for k,v:=range a {
		if InArray(k,keys){
		    a[k]=Merge(logDir,v,b[k])
		}
	}
	return a
}

func MatchJson(defaultDir string,configName string,datajson string)string{
	confParams:=config.LocadParams(defaultDir,configName)
	var keys []string
	for k,_:=range confParams.Params{
		keys=append(keys,k)
	}

	var raw map[string] interface{}
	logFile.LogNotice(defaultDir,datajson)
	err:=json.Unmarshal([]byte(datajson),&raw)
	if err != nil {
		logFile.LogWarnf(defaultDir,err)
	}
	r:=Match(keys,raw)
	logFile.LogNotice(defaultDir,keys)
	logFile.LogNotice(defaultDir,r)

	j,err:=json.Marshal(r)
	if err !=nil{
		logFile.LogWarnf(defaultDir,err)
	}
	logFile.LogNotice(defaultDir,j)
	return string(j)
}

func Match(conf []string,jsonData map[string]interface{}) map[string]interface {}{
	result:=make(map[string]interface{},0)
	for k,v:=range jsonData{
	re:=SubMatch(conf,k ,v)
		if re!=nil{
			result[k]=re
		}
	}
	return result
}
func SubMatch(conf []string,k string,v interface{}) interface {}{

	if lv,ok:=v.([]interface{});ok{
		list:=make([]interface{},0)
		for _,lvalue:=range lv{
			lval:=SubMatch(conf,k ,lvalue)
			if lval!=nil {
				list = append(list, lval)
			}
		}
		return list
	}else if mv,ok:=v.(map[string]interface{});ok{
		maps:=make(map[string]interface{})
		for k,v:=range mv {
			mapv:=SubMatch(conf,k,v)
			if mapv!=nil {
				maps[k] = mapv
			}
		}
		return maps
	}else if InArray(k,conf){
		return v
	}
	return nil
}

func InArray(key string,array []string) bool {
	for _,v:=range  array{
		if key==v{
			return  true
		}
	}
	return  false
}
