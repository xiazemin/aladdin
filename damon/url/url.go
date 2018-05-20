package url

import (
	"strings"
)
const sep  ="&"
const sepu="\\u0026"
const pairSep  ="="
const pairLen  =2
func ToJson(uri string)map[string]interface{}{
	var result []string
	if strings.Contains(uri,sepu){
		result=strings.Split(uri,sepu)
	}else if strings.Contains(uri,sep){
		result=strings.Split(uri,sep)
	}

	json:=make(map[string]interface{})
	for _,val:=range(result){
		pair:=strings.Split(val,pairSep)
		if(len(pair)==pairLen){
			json[pair[0]]=pair[1]
		}
	}
	return json
}
