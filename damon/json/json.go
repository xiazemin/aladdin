package json

import (
	"strings"
	io "io/ioutil"
	"github.com/xiazemin/aladdin/damon/logFile"
)
const sep  ="&"
const pairSep  ="="
func ToForm(dir string,json map[string]interface{}) string{
	var form string
	for key,value:=range(json){
		var pairs string
		if str, ok := value.(string);ok{
			logFile.LogDebug(dir,key+"=>"+str)
			pair:=make([]string,0)
			pair=append(pair,key)
			pair=append(pair,str)
			pairs=strings.Join(pair,pairSep)
		}
		if(pairs!="") {
			if form ==""{
			  form=pairs
			}else{
				form=form+sep+pairs
			}
		}
	}
    return form
}

func Load(filename string) []byte {
	data, err := io.ReadFile(filename)
	if err != nil{
		return nil
	}
	datajson := []byte(data)
	return datajson
}