package flag
import (
	"flag"
	"os"
	"fmt"
)
var tmplType=flag.Int("tmplType",2,"usage:-tmplType=1,2,,...(pwd,userdefined,dev,~)")
const tmplDefault  ="/Users/didi/goLang/src/github.com/xiazemin/aladdin/"
func GetTmplDir() string {
	flag.Parse()
	var dir string
	if *tmplType==1{
		dir, _ = os.Getwd()
	}else if(*tmplType==2){
		dir=tmplDefault
	}else  {
		dir, _ = os.Getwd()
	}
	fmt.Println(dir)
	fmt.Println(*tmplType)
	dir=AddSlash(dir)
	return dir+"view/tmpl/"
}