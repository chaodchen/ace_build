package controllers

import (
	"ace-img2/models"
	"log"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetGameIdInfo(c *gin.Context) {
	var pkgName = "controllers->GetGameIdInfo"
	var sdkVersion = c.DefaultQuery("sdkVersion", sdkVersionArray[0])
	log.Printf("[%s] sdkVersion: %s\n", pkgName, sdkVersion)
	
	code, msg, data := models.GetGameIdInfo(sdkVersion)
	log.Printf("[%s] code: %d msg: %s\n",pkgName, code, msg)
	c.Writer.Header().Set("__code", strconv.Itoa(code))
	c.Writer.Header().Set("__message", msg)
	c.JSON(200, gin.H{
		"data": data,
	})
}