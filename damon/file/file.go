package file

import (
	"os"
	"github.com/xiazemin/aladdin/damon/logFile"
	"io"
	"go/src/fmt"
)
/**
 * 判断文件是否存在  存在返回 true 不存在返回false
 */
func checkFileIsExist(filename string) bool {
	var exist = true
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		exist = false
	}
	return exist
}

func Write(dir string,filename string,content string){
	var f *os.File
	var err error
	if checkFileIsExist(dir+filename) { //如果文件存在
		f, err = os.OpenFile(dir+filename, os.O_RDWR, 0666) //打开文件
	} else {
		f, err = os.Create(dir+filename) //创建文件
		logFile.LogNotice(dir,"文件不存在")
	}
	if err!=nil{
		logFile.LogWarnf(dir,err)
		return
	}
	n, err1 := io.WriteString(f, content) //写入文件(字符串)
	if(err1!=nil){
		logFile.LogWarnf(dir,err1)
		return
	}
	logFile.LogNotice(dir,fmt.Sprintf("写入 %d 个字节n", n))
}

