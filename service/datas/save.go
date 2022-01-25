package datas

import (
	"github.com/binganao/TaiO/lib"
	"github.com/binganao/TaiO/model/db"
	"github.com/binganao/TaiO/pkg/logger"
	"strconv"
	"time"
)

func Save(host, ports, services, fingers string, start time.Time) {
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
			elapsed := time.Since(start).Minutes()
			pE := strconv.FormatFloat(elapsed, 'E', -1, 64)
			logger.Success("目标 " + host + " 已完成探测，用时: " + pE)
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
		elapsed := time.Since(start).Minutes()
		pE := strconv.FormatFloat(elapsed, 'G', 3, 64)
		logger.Success("目标 " + host + " 已完成探测，用时: " + pE + " m，5 秒后进入下一任务")
	}
}

func TmpSave(host, ports, services, fingers string) {
	var probT []db.ProbTmpM
	lib.GetDB().Where("host = ?", host).Find(&probT)
	if len(probT) == 0 {
		probt := db.ProbTmpM{
			Host:     host,
			Ports:    ports,
			Services: services,
			Fingers:  fingers,
		}
		if err := lib.GetDB().Create(&probt).Error; err == nil {
			logger.Success("目标 " + host + " 已完成探测，已经添加至临时数据中")
		} else {
			logger.Error("目标 " + host + " 已完成探测，添加至临时数据失败")
		}
	} else {
		probt := db.ProbTmpM{
			Host:     host,
			Ports:    ports,
			Services: services,
			Fingers:  fingers,
		}
		lib.GetDB().Model(&db.ProbTmpM{}).Where("host = ?", host).Update("ports", probt.Ports)
		lib.GetDB().Model(&db.ProbTmpM{}).Where("host = ?", host).Update("services", probt.Services)
		lib.GetDB().Model(&db.ProbTmpM{}).Where("host = ?", host).Update("fingers", probt.Fingers)
		logger.Success("目标 " + host + " 已完成探测，已经存在临时数据，更新完成！")
	}
}
