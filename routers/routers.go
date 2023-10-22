package routers

import (
	"github.com/gin-gonic/gin"
	"printdata/controller"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	// 静态文件
	// 模板文件
	// 路由组
	apiGrop := r.Group("api")
	{
		// 增
		apiGrop.POST("/data", controller.Create)
		apiGrop.POST("/data/uploadfile", controller.UploadFile)
		// 删
		apiGrop.DELETE("/data/:id", controller.Delete)
		// 改
		// 查
		apiGrop.GET("/data", controller.GetAll)
	}
	return r
}
