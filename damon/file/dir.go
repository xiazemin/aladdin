package file

import (
	"github.com/xiazemin/aladdin/damon/logFile"
	"io/ioutil"
	"go/src/fmt"
	"strings"
)

type Dir struct {
	Path string
	Name string
	SubDirs []*Dir
	FileNames []string
}

func GetDir(path string,name string,dirLog string)*Dir  {
	dir:=new(Dir)
	dir.Name=name
	dir.Path=path
       return SubDirs(dir,dirLog)
}

func SubDirs(dir *Dir,dirLog string)*Dir  {
	files, _ := ioutil.ReadDir(dir.Path)
	for _, file := range files {
		if file.IsDir() {
			logFile.LogNotice(dirLog,dir.Name  + file.Name()+ "/")
			subDir:=GetDir(dir.Path+file.Name()+ "/",file.Name(),dirLog)
			//allDirs=append(allDirs,surrentFolder + file.Name()+ "/")
			dir.SubDirs=append(dir.SubDirs,subDir)
		} else {
			//fmt.Println(surrentFolder + "/" + file.Name())
			logFile.LogNotice(dirLog,dir.Name + file.Name())
			dir.FileNames=append(dir.FileNames,file.Name())

		}
	}
	return  dir
}

func PrintDirs(dir*Dir,k int)string{
	space := strings.Repeat(" ", 3 * k)
	fmt.Println(space+"|"+"--"+dir.Name)
	str:=space+"|"+"--"+dir.Name+"\n"
	str=str+PrintSubDir(dir,k+1)

	return str
}

func PrintSubDir(dir*Dir,k int)string{
	space := strings.Repeat(" ", 3 * k)
	str:=""
	for _,file:=range dir.FileNames{
		fmt.Println(space+"|"+"--"+file)
		str=str+space+"|"+"--"+file+"\n"
	}
	for _,subDir:=range dir.SubDirs{
		str=str+PrintDirs(subDir,k)

	}
	return str
}

func GetPrintDirs(dir*Dir,k int)string{
	space := strings.Repeat(" ", 3 * k)
	str:=space+"|"+"--"+dir.Name+"\n"
	str=str+GetPrintSubDir(dir,k+1)

	return str
}

func GetPrintSubDir(dir*Dir,k int)string{
	space := strings.Repeat(" ", 3 * k)
	str:=""
	for _,file:=range dir.FileNames{
		str=str+space+"|"+"--"+file+"\n"
	}
	for _,subDir:=range dir.SubDirs{
		str=str+GetPrintDirs(subDir,k)

	}
	return str
}