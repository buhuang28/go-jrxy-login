package main

import (
	"log"
	"os"
	"time"
)

var (
	logger *log.Logger
)

func init() {
	logFileNmae := `./log/` + time.Now().Format("20060102") + "loginapi.log"
	logFileAllPath := logFileNmae
	_, err := os.Stat(logFileAllPath)
	exits := CheckFileIsExits(`log`)
	if !exits {
		_ = os.Mkdir("./log", os.ModePerm)
	}
	var f *os.File
	if err != nil {
		f, _ = os.Create(logFileAllPath)
	} else {
		//如果存在文件则 追加log
		f, _ = os.OpenFile(logFileAllPath, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
	}
	logger = log.New(f, "", log.LstdFlags)
}

func main() {
	GinRun()
}
