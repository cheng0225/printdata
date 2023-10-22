package main

import (
	"printdata/dao"
	"printdata/model"
	"printdata/routers"
)

func main() {
	err := dao.InitMysql() // 初始化数据库
	if err != nil {
		panic(err)
	}
	_ = dao.DB.AutoMigrate(&model.Data{})
	// 注册路由
	r := routers.SetupRouter()
	err = r.Run("127.0.0.1:8000")
	if err != nil {
		return
	}
}
