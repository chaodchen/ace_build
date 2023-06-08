package main

import (
	"ace-img2/config"
	"ace-img2/routes"
	"fmt"
	"log"
    "os"
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
    // create zips acesdks anosdks
    need_dirs := []string {
        "static/zips",
        "static/acesdks",
        "static/anosdks",
    }
    
    for _, dir := range need_dirs {
        if _, err := os.Stat(dir); os.IsNotExist(err) {
            err := os.Mkdir(dir, 0755)
            if err != nil {
                log.Println("create need_dir failed.")
                return
            }
        }

    }

	ggin = routes.InitRouter()
	log.Printf("[%s] main init.\n", pkgName)

}

func main() {
	var pkgName = "main.main"
	log.Printf("[%s] Hello AceServer!\n", pkgName)
	ggin.Run(fmt.Sprintf(":%d", config.Config.Port))
}
