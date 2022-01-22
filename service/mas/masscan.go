package mas

import (
	"github.com/binganao/Taio/common"
	"github.com/zan8in/masscan"
	"strconv"
)

func MasScan(ip string) []string {
	var result []string
	scanner, err := masscan.NewScanner(
		masscan.SetParamTargets(ip),
		masscan.SetParamPorts("1-65535"),
		masscan.SetParamWait(0),
		masscan.SetParamRate(common.MASSCAN_RATE),
	)
	if err != nil {
		return nil
	}

	scanResult, _, err := scanner.Run()
	if err != nil {
		return nil
	}

	if scanResult != nil {
		for i, _ := range scanResult.Hosts {
			result = append(result, strconv.Itoa(scanResult.Ports[i].Port))
		}
	}
	return result
}
