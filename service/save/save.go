package save

import (
	"github.com/binganao/Taio/lib"
	"github.com/binganao/Taio/model/db"
	"github.com/binganao/Taio/pkg/logger"
)

func Save(host, ports, services, fingers string) {
	var probT []db.ProbM
	lib.GetDB().Where("host = ?", host).Find(&probT)
	if len(probT) == 0 {
		probt := db.ProbM{
			Host:     host,
			Ports:    ports,
			Services: services,
			Fingers:  fingers,
		}
		if err := lib.GetDB().Create(&probt).Error; err == nil {
			logger.Success("目标 " + host + " 已完成探测!")
		} else {
			logger.Error("目标 " + host + " 已完成探测，但写入数据失败!")
		}
	} else {
		probt := db.ProbM{
			Host:     host,
			Ports:    ports,
			Services: services,
			Fingers:  fingers,
		}
		lib.GetDB().Model(&db.ProbM{}).Where("host = ?", host).Update("ports", probt.Ports)
		lib.GetDB().Model(&db.ProbM{}).Where("host = ?", host).Update("services", probt.Services)
		lib.GetDB().Model(&db.ProbM{}).Where("host = ?", host).Update("fingers", probt.Fingers)
		logger.Success("目标 " + host + " 已完成探测!")
	}
}

func TmpSave(host, ports, services, fingers string) {
	var probT []db.AddTmp
	lib.GetDB().Where("host = ?", host).Find(&probT)
	if len(probT) == 0 {
		probt := db.AddTmp{
			Host:     host,
			Ports:    ports,
			Services: services,
			Fingers:  fingers,
		}
		if err := lib.GetDB().Create(&probt).Error; err == nil {
			logger.Success("目标 " + host + " 已完成探测!")
		} else {
			logger.Error("目标 " + host + " 已完成探测，但写入数据失败!")
		}
	} else {
		probt := db.AddTmp{
			Host:     host,
			Ports:    ports,
			Services: services,
			Fingers:  fingers,
		}
		lib.GetDB().Model(&db.AddTmp{}).Where("host = ?", host).Update("ports", probt.Ports)
		lib.GetDB().Model(&db.AddTmp{}).Where("host = ?", host).Update("services", probt.Services)
		lib.GetDB().Model(&db.AddTmp{}).Where("host = ?", host).Update("fingers", probt.Fingers)
		logger.Success("目标 " + host + " 已完成探测!")
	}
}
