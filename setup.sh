#!/bin/bash
pid=`ps aux |grep watcher |grep -v 'grep' |cut -d ' ' -f 4`
if [[ $pid = 0 ]]; then
echo 'no process ';
else
ps aux |grep watcher |grep -v 'grep' |cut -d ' ' -f 4 |xargs kill -9
fi

if [[ -f "watcher.tar" ]]; then
rm watcher.tar*
fi
wget http://127.0.0.1:8088/download/watcher.tar
tar -zxvf watcher.tar
nohup ./watcher -r -dirType=1 -ip="127.0.0.1" -port=8088  &
#bash <(curl -s -S -L http://127.0.0.1:8088/download/setup.sh)