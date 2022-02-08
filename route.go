package main

import (
	"github.com/gin-gonic/gin"
	"io"
	"os"
)

func GinRun() {
	var logfile *os.File
	exits := CheckFileIsExits("./gin.log")
	if !exits {
		logfile, _ = os.Create("./gin.log")
	} else {
		//如果存在文件则 追加log
		logfile, _ = os.OpenFile("./gin.log", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
	}
	gin.SetMode(gin.DebugMode)
	gin.DefaultWriter = io.MultiWriter(logfile)
	router := gin.Default()
	router.POST("/login", Login)
	router.Run(":8001")
}
