package service

import (
	"fmt"
	"github.com/binganao/Taio/pkg/logger"
	"github.com/binganao/Taio/service/finger"
	"github.com/binganao/Taio/service/mas"
	"github.com/binganao/Taio/service/nma"
	"github.com/binganao/Taio/service/save"
)

func Probe(ip string) {
	var web [][]string

	ports := mas.MasScan(ip)

	var dP string
	for i, p := range ports {
		if i == 0 {
			dP += p
		} else {
			dP += "," + p
		}
	}

	services := nma.NmapScan(ip, ports)

	fmt.Println(services)

	var dS string
	for i, p := range services {
		if i == 0 {
			dS += p
		} else {
			dS += "," + p
		}
	}

	logger.Info("开始对目标 " + ip + " 进行 WEB 指纹探测")

	for _, s := range services {
		if s != "" {
			wt := finger.FingerScan(ip, s)
			if wt != nil {
				web = append(web, wt)
			}
		}
	}

	var dW string
	for i, wa := range web {
		if i == 0 {
			for j, a := range wa {
				if j == 0 {
					dW += a
				} else {
					dW += "," + a
				}
			}
		} else {
			dW += ";"
			for j, a := range wa {
				if j == 0 {
					dW += a
				} else {
					dW += "," + a
				}
			}
		}
	}

	save.Save(ip, dP, dS, dW)
}
