package service

import (
	"github.com/binganao/Taio/pkg/logger"
	"github.com/binganao/Taio/service/jobs"
	"github.com/binganao/Taio/utils/parse"
	"strconv"
	"time"
)

func AyncProbe() {
	for true {
		if len(jobs.GetJobs(false)) == 0 {
			time.Sleep(5 * time.Second)
			logger.Warning("任务列表为空!")
			continue
		}
		for raws := range jobs.GetJobs(true) {
			if raws == "" {
				continue
			}
			jobs.Lock = true
			if jobs.DWorking == raws {
				logger.Success("任务 " + raws + " 的删除操作已完成!")
				jobs.Working = ""
				jobs.DWorking = ""
				jobs.DelJobs(raws)
				continue
			}
			jobs.Working = raws
			logger.Info("检测到任务: " + raws + ", 开始扫描...")
			hosts := parse.ParseIP(raws)
			if len(hosts) == 0 {
				jobs.Lock = false
				break
			}
			for i, host := range hosts {
				logger.Info("启动探测 " + host + ", 当前任务进度: " + strconv.Itoa(i+1) + "/" + strconv.Itoa(len(hosts)))
				Probe(host)
				time.Sleep(5 * time.Second)
			}
			jobs.Working = ""
			jobs.Lock = false
			jobs.AddJobs(raws, false)
		}
	}
}
