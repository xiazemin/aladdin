package file
import (
	"io/ioutil"
	"github.com/xiazemin/aladdin/damon/logFile"
)
func ListDirs(surrentFolder string,dirLog string) []string {

	var allDirs []string
	allDirs=ListSubDirs(allDirs,surrentFolder,dirLog)
	return allDirs
}

func ListSubDirs(allDirs []string,surrentFolder string,dirLog string)[]string  {
	allDirs=append(allDirs,surrentFolder)
	files, _ := ioutil.ReadDir(surrentFolder)
	for _, file := range files {
		if file.IsDir() {
			logFile.LogNotice(dirLog,surrentFolder  + file.Name()+ "/")
			//allDirs=append(allDirs,surrentFolder + file.Name()+ "/")
			allDirs=ListSubDirs(allDirs,surrentFolder + file.Name() + "/",dirLog)
		} else {
			//fmt.Println(surrentFolder + "/" + file.Name())
			logFile.LogNotice(dirLog,surrentFolder + file.Name())
		}
	}
	return  allDirs
}
