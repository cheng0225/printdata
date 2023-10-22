package model

import (
	"printdata/dao"
)

// Data `gorm:"column:age_of_the_beast;default:42" json:"前端的json字段，不写默认列名"`
// binding:"required"修饰的字段，若接收为空值，则报错，是必须字段
// `form:"username" json:"user" uri:"user" xml:"user" binding:"required"`
type Data struct {
	ID                uint
	PrintTime         string // `gorm:"default:42" json:"time"`
	PrintHeight       string
	PrintSpeed        string
	ProfileVelocity   string
	IdleAcceleration  string
	PullbackLength    string
	PullbackSpeed     string
	NozzleTemperature string
	BadTemperature    string
	PrintCondition    string
	ProblemCollection string
	FilePath          string
}

func Create(data *Data) (err error) {
	err = dao.DB.Create(data).Error
	return
}

func GetAll() (dataList []*Data, err error) {
	err = dao.DB.Find(&dataList).Error
	return
}

func Delete(id string) (err error) {
	err = dao.DB.Where("id=?", id).Delete(&Data{}).Error
	return
}
