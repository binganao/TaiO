package jobs

import (
	"github.com/binganao/Taio/pkg/logger"
	"strconv"
)

var jobs chan string
var Lock = false

func InitJobs() {
	jobs = make(chan string, 10000)
	logger.Success("初始化任务队列成功!")
}

func GetJobs(f bool) chan string {
	if f {
		logger.Success("读取任务队列成功, 任务队列共计: " + strconv.Itoa(len(jobs)) + " 条!")
	}
	return jobs
}

func AddJobs(raw string, f bool) {
	if raw == "" {
		return
	}
	jobs <- raw
	if f {
		if Lock {
			logger.Success("新任务 " + raw + " 添加成功，当前任务队列共计: " + strconv.Itoa(len(jobs)+1) + " 条!")
		} else {
			logger.Success("新任务 " + raw + " 添加成功，当前任务队列共计: " + strconv.Itoa(len(jobs)) + " 条!")
		}
	}
}

func DelJobs(raw string) {
	go func() {
		for i := range jobs {
			if i == raw {
				logger.Success("任务 " + raw + " 删除成功，当前任务队列共计: " + strconv.Itoa(len(jobs)) + " 条!")
				break
			} else {
				jobs <- i
				continue
			}
		}
	}()
}
