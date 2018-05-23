package autorun
import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/howeyc/fsnotify"
	"github.com/xiazemin/aladdin/damon/logFile"
	"github.com/xiazemin/aladdin/damon/file"
)

const defaultTimeSpan  =1 //s

var (
	eventTime    = make(map[string]int64)
	scheduleTime time.Time
)

func NewWatcher(dir string,paths []string,dirExe string,name string) {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		logFile.LogWarnf(dir,fmt.Sprintf("[ERRO] Fail to create new Watcher[ %s ]\n", err))
		os.Exit(2)
	}

	go func() {
		for {
			select {
			case e := <-watcher.Event:
				isbuild := true

			logFile.LogNotice(dir,e)
			// Skip TMP files for Sublime Text.
				if checkTMPFile(e.Name) {
					continue
				}
				if !checkIfWatchExt(e.Name) {
					continue
				}

				mt := getFileModTime(dir,e.Name)
				if t := eventTime[e.Name]; mt == t {
					logFile.LogWarnf(dir,fmt.Sprintf("[SKIP] # %s #\n", e.String()))
					//isbuild = false
				}

				eventTime[e.Name] = mt
			        logFile.LogDebug(dir,fmt.Sprintf("eventTime:%+v,time %+v,isBuild:%+v,event:%+v",eventTime[e.Name],mt,isbuild,e))
				if isbuild {
					logFile.LogWarnf(dir,fmt.Sprintf("[EVEN] %s\n", e))
					go func() {
						// Wait 1s before autobuild util there is no file change.
						scheduleTime = time.Now().Add(defaultTimeSpan * time.Second)
						for {
							logFile.LogDebug(dir,scheduleTime)
							time.Sleep(scheduleTime.Sub(time.Now()))
							if time.Now().After(scheduleTime) {
							logFile.LogDebug(dir,time.Now())
									break
							}
							return
						}

						AutoRun(dir,dirExe,name)
					}()
					//AutoRun(dir,dirExe,name)
				}
			case err := <-watcher.Error:
				logFile.LogWarnf(dir,fmt.Sprintf("[WARN] %s\n", err.Error())) // No need to exit here
			}
		}
	}()

	logFile.LogNotice(dir,"[INFO] Initializing watcher...\n")
	for _, path := range paths {
		logFile.LogNotice(dir,fmt.Sprintf("[TRAC] Directory( %s )\n", path))
		logFile.LogDebug(dir,fmt.Sprintf("[TRAC] Directory( %s )\n", path))
		allDirs:=file.ListDirs(path,dir)
                logFile.LogNotice(dir,allDirs)
		logFile.LogDebug(dir,allDirs)
		for _,subpath:=range (allDirs) {
			err = watcher.Watch(subpath)
			if err != nil {
				logFile.LogWarnf(dir, fmt.Sprintf("[ERRO] Fail to watch directory[ %s ]\n", err))
				os.Exit(2)
			}
		}
	}

}

// getFileModTime retuens unix timestamp of `os.File.ModTime` by given path.
func getFileModTime(dir string,path string) int64 {
	path = strings.Replace(path, "\\", "/", -1)
	f, err := os.Open(path)
	if err != nil {
		logFile.LogWarnf(dir,fmt.Sprintf("[ERRO] Fail to open file[ %s ]\n", err))
		return time.Now().Unix()
	}
	defer f.Close()

	fi, err := f.Stat()
	if err != nil {
		logFile.LogWarnf(dir,fmt.Sprintf("[ERRO] Fail to get file information[ %s ]\n", err))
		return time.Now().Unix()
	}

	return fi.ModTime().Unix()
}

// checkTMPFile returns true if the event was for TMP files.
func checkTMPFile(name string) bool {
	if strings.HasSuffix(strings.ToLower(name), ".tmp") {
		return true
	}
	return false
}

var watchExts = []string{".go",".php"}

// checkIfWatchExt returns true if the name HasSuffix <watch_ext>.
func checkIfWatchExt(name string) bool {
	for _, s := range watchExts {
		if strings.HasSuffix(name, s) {
			return true
		}
	}
	return false
}