package main

import (
	"github.com/gin-gonic/gin"
	"printdata/dao"
	"printdata/middlewares"
	"printdata/model"
	"printdata/routers"
)

func main() {
	err := dao.InitMysql() // 初始化数据库
	if err != nil {
		panic(err)
	}
	_ = dao.DB.AutoMigrate(&model.Data{}) // 模型绑定

	r := gin.Default()
	// 中间件 跨域
	r.Use(middlewares.Cors())

	// 注册路由
	routers.SetupRouter(r)
	routers.Download(r)

	err = r.Run("127.0.0.1:8000")
	if err != nil {
		return
	}
}
