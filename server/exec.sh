#!/bin/bash

#print time
echo $(date +"%Y-%m-%d %H:%M:%S")
echo "项目开始部署~"

cd /Users/xwx/go/src/AutoDeploy/log/

rm -rf myblog

git clone "git@github.com:zty-f/myblog.git"

hexo c && hexo s

hexo d

echo "项目部署完成~"
echo $(date +"%Y-%m-%d %H:%M:%S")