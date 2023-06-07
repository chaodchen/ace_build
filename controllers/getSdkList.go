package controllers

import (
	"ace-img2/models"
	"log"
	"strconv"
	"github.com/gin-gonic/gin"
)


func GetSdkList(c *gin.Context) {
	var pkgName = "controllers->GetSdkList"
	
	log.Printf("[%s] Start GetSdkList.\n", pkgName)
	sdkVersion := c.DefaultQuery("sdkVersion", sdkVersionArray[0])
	log.Printf("[%s] sdkVersion: %s\n", pkgName, sdkVersion)
	
	code , msg , data := models.GetSdkList(sdkVersion)
	c.Writer.Header().Set("__code", strconv.Itoa(code))
	c.Writer.Header().Set("__message", msg)

	c.JSON(200, gin.H{
		"data": data,
	})
}

