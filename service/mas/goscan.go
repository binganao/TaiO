package mas

import (
	"fmt"
	"github.com/binganao/Taio/lib"
	"github.com/binganao/Taio/model/db"
	"github.com/binganao/Taio/utils/parse"
)

func Run(raw string) {
	hosts := parse.ParseIP(raw)
	for _, ip := range hosts {
		res := MasScan(ip)
		if len(res) > 0 {
			saveHost(ip, res)
		}
	}
}

func saveHost(host string, portslice []string) {
	ports := ""
	for i, p := range portslice {
		if i == 0 {
			ports += p
		} else {
			ports += "," + p
		}
	}
	h := db.MascM{
		Host:  host,
		Ports: ports,
	}
	fmt.Println(h)
	lib.GetDB().Create(&h)
}
