package service

import (
	"fmt"
	"github.com/binganao/TaiO/pkg/logger"
	"github.com/binganao/TaiO/service/jobs"
	"strings"
)

func NodeLoop() {
	var command string
	commands := []string{"add ", "del ", "stop"}
	fmt.Print("> ")
	fmt.Scanf("%s", &command)
	for i, j := range commands {
		if strings.Contains(command, j) {
			if i == 0 {
				host := strings.Split(command, " ")[1]
				jobs.AddJobs(host, true)
				break
			} else if i == 1 {
				host := strings.Split(command, " ")[1]
				jobs.DelJobs(host)
				break
			} else if i == 2 {
				jobs.StopJobs()
				break
			} else {
				logger.Error("您的命令输入有误，请重新输入！")
				break
			}
		}
	}
}
