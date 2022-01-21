package nma

import (
	"context"
	"github.com/binganao/Taio/pkg/logger"
	"strconv"
	"time"

	"github.com/Ullaakut/nmap/v2"
)

func NmapScan(ip string, portSlice []string) []string {
	var ports string
	for i, p := range portSlice {
		if i == 0 {
			ports += p
		} else {
			ports += "," + p
		}
	}
	var results []string
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Minute)
	defer cancel()

	scanner, err := nmap.NewScanner(
		nmap.WithTargets(ip),
		nmap.WithPorts(ports),
		nmap.WithContext(ctx),
		nmap.WithServiceInfo(),
	)
	if err != nil {
		logger.Error("unable to create nma scanner:")
		return nil
	}

	result, warnings, err := scanner.Run()
	if err != nil {
		logger.Error("unable to run nma scan")
		return nil
	}

	if warnings != nil {
		logger.Info("Nmap Warnings")
	}

	for _, host := range result.Hosts {
		if len(host.Ports) == 0 || len(host.Addresses) == 0 {
			continue
		}

		for _, port := range host.Ports {
			results = append(results, strconv.Itoa(int(port.ID))+":"+port.Service.Name)
		}
	}

	return results
}
