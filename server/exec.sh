#!/bin/bash

#print time
# shellcheck disable=SC2046
# shellcheck disable=SC2005
echo $(date +"%Y-%m-%d %H:%M:%S")
echo "项目开始部署~"

# shellcheck disable=SC2164
cd /home/data/

rm -rf myblog

git clone "git@github.com:zty-f/myblog.git"

cd myblog

hexo c && hexo g

hexo d

echo "项目部署完成~"
echo $(date +"%Y-%m-%d %H:%M:%S")
