#!/bin/bash
GOOS=linux GOARCH=amd64 go build -o aladdin  src/github.com/xiazemin/aladdin/main/main.go
#$ nohup ./aladdin -dirType=1 -r -tmplType=1 -ip="127.0.0.1" -port=8088 &
GOOS=linux GOARCH=amd64 go build -o watcher  src/github.com/xiazemin/aladdin/main/run.go
#$ nohup ./watcher -r -dirType=1 -ip="127.0.0.1" -port=8088  &

tar czvf watcher.tar configWatch.json  globalConfig.json logConfig.json watcher
tar czvf aladdin.tar  tmpl  globalConfig.json  logConfig.json config/params.json config/data.json data/xiazemin/raw.log aladdin

mkdir download
mv aladdin.tar download/
mv watcher.tar download/
wget http://127.0.0.1:8088/download/aladdin.tar
wget http://127.0.0.1:8088/download/watcher.tar
tar -zxvf aladdin.tar
tar -zxvf watcher.tar

#http://127.0.0.1:8088/download/