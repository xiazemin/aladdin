package main

import (
	"log"
	"github.com/howeyc/fsnotify"
	"os"
	"fmt"
	"strconv"
)

func usage()  {
	fmt.Println("")
	fmt.Printf("Usage: %s Montior-directory file-max-bytes-limit\n",os.Args[0])
	fmt.Println("For example:")
	fmt.Printf("%s /opt/logs 1024\n",os.Args[0])
	os.Exit(1)
}

func isDir(dirname string) bool  {
	fhandler, err := os.Stat(dirname);
	if(! (err == nil || os.IsExist(err)) ) {
		return false
	}else {
		return fhandler.IsDir()
	}
}

func isFile(filename string) bool  {
	fhandler, err := os.Stat(filename);
	if(! (err == nil || os.IsExist(err)) ) {
		return false
	}else if (fhandler.IsDir()){
		return false
	}
	return true
}

func emptiedFile(filename string) bool  {
	FN,err := os.Create(filename)
	defer FN.Close()
	if err != nil {
		return false
	}
	fmt.Fprint(FN,"")
	return true
}

func getFileByteSize(filename string) (bool,int64) {
	if (! isFile(filename)) {
		return false,0
	}
	fhandler, _ := os.Stat(filename);
	return true,fhandler.Size()
}

func main() {
	var maxByte int64 = 1024*1024
	if (len(os.Args) < 2) {
		usage()
	}
	if (len(os.Args) >= 3) {
		maxByte_,err := strconv.Atoi(os.Args[2])
		if err != nil {
			log.SetPrefix("[ERROR] ")
			log.Println(os.Args[2],"Is not a legitimate int number")
			usage()
		}
		maxByte = int64(maxByte_)
	}
	dirpath := os.Args[1]
	if (!isDir(dirpath)){
		log.SetPrefix("[ERROR] ")
		log.Println(dirpath,"Is not a legitimate directory")
		usage()
	}
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}
	done := make(chan bool)

	//Process event
	go func() {
		for {
			select {
			case ev := <-watcher.Event:
				if(ev.IsModify()){
					_,size := getFileByteSize(ev.Name)
					log.Println("event:",ev,",byte:",size)
					if (size >= maxByte){
						if ( ! emptiedFile(ev.Name) ) {
							log.SetPrefix("[ERROR] ")
							log.Printf("%s :Can not empty file\n",ev.Name)
						}
					}
				}
			case err := <-watcher.Error:
				log.Println("error:",err)
			}
		}
	}()
	err = watcher.Watch(dirpath)
	if err != nil {
		log.Fatal(err)
	}
	<-done

	watcher.Close()
}