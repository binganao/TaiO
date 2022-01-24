package datas

import (
	"github.com/binganao/Taio/lib"
	"github.com/binganao/Taio/model/db"
	"github.com/binganao/Taio/pkg/logger"
)

func RemoveTmp(host string) {
	lib.GetDB().Where("host = ?", host).Delete(&db.ProbTmpM{})
	logger.Info("临时数据: " + host + "已删除!")
}

func Remove(host string) {
	lib.GetDB().Where("host = ?", host).Delete(&db.ProbM{})
	logger.Warning("数据: " + host + "已删除!")
}
