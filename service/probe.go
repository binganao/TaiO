package service

import (
	"github.com/binganao/TaiO/common"
	"github.com/binganao/TaiO/pkg/logger"
	"github.com/binganao/TaiO/service/datas"
	"github.com/binganao/TaiO/service/finger"
	"github.com/binganao/TaiO/service/mas"
	"github.com/binganao/TaiO/service/nma"
	"github.com/binganao/TaiO/utils/crypto"
	"github.com/binganao/TaiO/utils/request"
	"time"
)

func Probe(ip string) {

	start := time.Now()

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

	logger.Info("开始对目标 " + ip + " 进行服务识别")
	services := nma.NmapScan(ip, ports)

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
	if common.NODE_TYPE != "node" {
		datas.Save(ip, dP, dS, dW, start)
	} else {
		/**
		发送数据到主端
		*/
		sip := "host=" + crypto.Base64Enctypto(ip)
		sdp := "ports=" + crypto.Base64Enctypto(dP)
		sds := "services=" + crypto.Base64Enctypto(dS)
		sdw := "fingers=" + crypto.Base64Enctypto(dW)
		request.Post(common.DATA_ADDR+"/api/v1/data/tmp/add", sip+"&"+sdp+"&"+sds+"&"+sdw)
	}
}
