package main

import (
	"ace-img2/config"
	"ace-img2/routes"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
)

var ggin *gin.Engine

func init() {
	var pkgName = "main.init"
	// 读取配置并初始化
	// 设置Log保存路径
	// logFile, err := os.OpenFile(prop.LogPath, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
	// if err != nil {
	// 	log.Fatalln(err.Error())
	// }
	// log.SetOutput(logFile)
	// logFile.Close()

	ggin = routes.InitRouter()
	log.Printf("[%s] main init.\n", pkgName)

}

func main() {
	var pkgName = "main.main"
	log.Printf("[%s] Hello AceServer!\n", pkgName)
	ggin.Run(fmt.Sprintf(":%d", config.Config.Port))
}
