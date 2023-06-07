package controllers

import (
	"ace-img2/config"
	"ace-img2/models"
	"fmt"
	"log"
	"net/url"
	"path/filepath"
	"strconv"

	"github.com/gin-gonic/gin"
)

func BuildDemo(c *gin.Context) {
	var pkgName = "controllers.BuildDemo"
	var buildConfig models.BuildConfig
	if err := c.BindJSON(&buildConfig); err != nil {
		log.Printf("[%s] error: %s\n", pkgName, err)
		c.Writer.Header().Set("__code", "4001")
		c.Writer.Header().Set("__message", "解析json失败")
		c.String(200, "解析Json失败")
	} else {
		code, msg := models.BuildDemo(buildConfig)
		c.Writer.Header().Set("__code", strconv.Itoa(code))
		c.Writer.Header().Set("__message", msg)
		if code == 0 {
			var outApkPath = filepath.Join(config.Config.DemoPath, "/launcher/build/outputs/apk/release/launcher-release.apk")
			log.Printf("[%s] outApkPath: %s", pkgName, outApkPath)
			c.Writer.Header().Set("Content-Disposition", fmt.Sprintf("attachment;filename=%s_%s.apk", url.PathEscape(buildConfig.AppName), buildConfig.SdkVersion))

			c.File(outApkPath)
		} else {
			c.String(200, msg)
		}
	}
}
