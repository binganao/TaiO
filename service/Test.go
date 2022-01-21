package service

import (
	"fmt"
	"github.com/binganao/Taio/service/finger"
)

func Test() {
	var web [][]string
	ip := "121.42.148.29"
	services := []string{"999:http", "443:https", "2121:ftp", "3306:mysql", "6006:tcpwrapped", "6009:tcpwrapped", "6010:tcpwrapped", "6014:tcpwrapped", "6021:tcpwrapped", "6027:tcpwrapped", "6028:tcpwrapped", "6034:tcpwrapped", "6037:tcpwrapped", "6038:tcpwrapped", "6039:tcpwrapped", "6041:tcpwrapped", "8090:http", "49153:msrpc", "49154:msrpc", "49160:msrpc", "49168:msrpc"}
	for _, s := range services {
		if s != "" {
			wt := finger.FingerScan(ip, s)
			if wt != nil {
				web = append(web, wt)
			}
		}
	}
	fmt.Println(web)
}
