package jobs

import (
	"github.com/binganao/Taio/pkg/logger"
	"strconv"
)

var jobs chan string

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
	if f {
		logger.Success("新任务 " + raw + " 添加成功，当前任务队列共计: " + strconv.Itoa(len(jobs)+1) + " 条!")
	}
	jobs <- raw
}

func DelJobs(raw string) {
	go func() {
		if len(jobs) == 0 {
			return
		}
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

func nothing(i string) {
	i = ""
	return
}
