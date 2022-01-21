package service

import (
	"github.com/binganao/Taio/service/jobs"
	"github.com/binganao/Taio/utils/parse"
)

func AyncProbe() {
	for true {
		gjobs := jobs.GetJobs()
		for raws := range gjobs {
			hosts := parse.ParseIP(raws)
			for _, host := range hosts {
				Probe(host)
			}
			jobs.AddJobs(raws)
		}
	}
}
