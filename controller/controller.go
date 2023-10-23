package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"printdata/model"
	"printdata/tool"
)

func GetAll(c *gin.Context) {
	dataList, err := model.GetAll()
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"err": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"code": 200, "message": "", "data": dataList})
}

func Delete(c *gin.Context) {
	id, ok := c.Params.Get("id")
	if !ok {
		c.JSON(http.StatusOK, gin.H{"err": "无效id"})
	}
	if err := model.Delete(id); err != nil {
		c.JSON(http.StatusOK, gin.H{"err": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"code": 200, "message": "删除成功"})
}

func Create(c *gin.Context) {
	var data model.Data
	// err := c.BindJSON(&data) // 解析JSON数据
	err := c.Bind(&data) // 解析form表单数据  josn数据也可以解析
	if err != nil {
		return
	}
	if err = model.Create(&data); err != nil {
		c.JSON(http.StatusOK, gin.H{"err": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"code": 200, "data": data})
}

func UploadFile(c *gin.Context) {
	// 获取文件
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"err": err.Error()})
		return
	}
	// 生成文件名 和文件路径 并保存文件
	fileName, err := tool.GetFileName(file.Filename)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"err": err.Error()})
	}
	filePath := "uploads/" + fileName
	if err = c.SaveUploadedFile(file, filePath); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"err": err.Error()})
		return
	}

	// 内容提取 保存到数据库
	// fileContent, openErr := tool.ReadDataFromFile(filePath)
	// if openErr != nil || getErr != nil{
	data, getErr := tool.GetDataFromFile(filePath)
	if getErr != nil {
		// 读取文件失败时进行删除
		_ = os.Remove(filePath)
		c.JSON(http.StatusInternalServerError, gin.H{"err": err.Error()})
	}

	// data := model.Data{
	// 	FilePath: filePath,
	// }
	if err = model.Create(data); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"err": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "文件上传成功"})
}
