package controllers

import (
	"ace-img2/models"
	"log"
	"path/filepath"
	"strconv"

	"github.com/gin-gonic/gin"
)


func UpLoadSdk(c *gin.Context) {
	var code = 0
	var msg = "ok"
	var pkgName = "controllers.UpLoadSdk"
	rFile, err := c.FormFile("file")

	if err != nil {
		code = 5001
		msg = err.Error()
	} else if rFile.Size > 100000000 { // 100m
		code = 5002
		msg = err.Error()
	} else {
		var dst = filepath.Join("static", "zips", rFile.Filename)
		log.Printf("[%s] dst: %s", pkgName, dst)
		err = c.SaveUploadedFile(rFile, dst)
		if err != nil {
			code = 5003
			msg = err.Error()
		} else {
			code, msg = models.UpLoadSdk(dst)
		}
	}

	c.Writer.Header().Set("__code", strconv.Itoa(code))
	c.Writer.Header().Set("__message", msg)
	c.String(200, msg)
}
