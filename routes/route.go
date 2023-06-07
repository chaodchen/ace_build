package routes

import (
	"ace-img2/controllers"

	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()

	api := r.Group("/api")
	{
		api.GET("/getEnableDat", controllers.MakeEnable)
		api.GET("/getSdkList", controllers.GetSdkList)
		api.POST("/buildDemo", controllers.BuildDemo)
		api.POST("/upLoadSdk", controllers.UpLoadSdk)

		// api.GET("/getGameIdInfo", controllers.GetGameIdInfo)

	}
	return r
}
