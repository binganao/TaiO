package service

import (
	"fmt"
	"github.com/binganao/Taio/pkg/logger"
	"github.com/binganao/Taio/service/finger"
	"github.com/binganao/Taio/service/mas"
	"github.com/binganao/Taio/service/nma"
)

func Probe(ip string) {
	var web [][]string
	logger.Info("开始对目标 " + ip + " 进行探测")
	ports := mas.MasScan(ip)
	logger.Info("目标 " + ip + " 开启了以下端口: ")
	fmt.Println(ports)
	services := nma.NmapScan(ip, ports)
	logger.Info("目标 " + ip + " 开启了以下服务: ")
	fmt.Println(services)
	for _, s := range services {
		if s != "" {
			wt := finger.FingerScan(ip, s)
			if wt != nil {
				web = append(web, wt)
			}
		}
	}
	logger.Info("目标 " + ip + " 的指纹识别结果: ")
	fmt.Println(web)
}
