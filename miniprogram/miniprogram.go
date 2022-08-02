package main

import (
	"log"
	"miniprogram/server"
	"miniprogram/util/config"
	"os"
)

func main() {
	// 初始化配置
	rootPath, _ := os.Getwd()
	if err := config.InitConfig(rootPath); err != nil {
		log.Fatalf("initConfig err %v", err)
	}

	logPath := rootPath + "/output.log"
	if s := server.NewServer(logPath); s.Run() != nil {
		return
	}
}
