package nma

import (
	"github.com/binganao/Taio/lib"
	"github.com/binganao/Taio/model/db"
	"github.com/binganao/Taio/utils/parse"
)

func Run(raw string, ports string) {
	hosts := parse.ParseIP(raw)
	for _, ip := range hosts {
		res := nmapScan(ip, ports)
		if len(res) > 0 {
			saveHost(ip, res)
		}
	}
}

func saveHost(host string, serviceSlice []string) {
	services := ""
	for i, s := range serviceSlice {
		if i == 0 {
			services += s
		} else {
			services += "," + s
		}
	}
	h := db.NamM{
		Host:    host,
		Service: services,
	}

	lib.GetDB().Create(&h)
}
