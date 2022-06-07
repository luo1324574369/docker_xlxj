#! /bin/sh
# 编译文件
cd miniprogram && CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build .


