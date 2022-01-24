package datas

import (
	"github.com/binganao/Taio/lib"
	"github.com/binganao/Taio/pkg/logger"
)

func RemoveTmp(host string) {
	lib.GetDB().Delete("host = ?", host)
	logger.Info("临时数据: " + host + "已删除!")
}
