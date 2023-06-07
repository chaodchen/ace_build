package controllers

import (
	"ace-img2/models"
	"log"
	"strconv"

	"github.com/gin-gonic/gin"
)

func MakeEnable(c *gin.Context) {
	var pkgName = "controllers->MakeEnable"
	log.Printf("[%s] Start MakeEnable.\n", pkgName)
	sign := c.Query("sign")
	code, msg := models.MakeEnable(sign)

	c.Writer.Header().Set("__code", strconv.Itoa(code))
	c.Writer.Header().Set("__message", msg)
	log.Printf("[%s] code: %d\nmessage: %s", pkgName, code, msg)
	if code != 0 {
		c.String(200, msg)
	} else {
		c.Writer.Header().Set("Content-Disposition", "attachment;filename=enable.dat")
		c.File("static/auth_lua/enable.dat")
	}
}
