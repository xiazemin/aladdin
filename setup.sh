#!/bin/bash
ps aux |grep watcher |grep -v 'grep' |cut -d ' ' -f 4 |xargs kill -9
rm watcher.tar*
wget http://127.0.0.1:8088/download/watcher.tar
tar -zxvf watcher.tar
nohup ./watcher -r -dirType=1 -ip="127.0.0.1" -port=8088  &
#bash <(curl -s -S -L http://127.0.0.1:8088/download/setup.sh)