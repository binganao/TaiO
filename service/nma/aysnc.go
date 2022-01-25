package nma

import (
	"context"
	"github.com/Ullaakut/nmap/v2"
	"github.com/binganao/TaiO/pkg/logger"
	"strconv"
	"time"
)

func aynscNmap(aHost, port string) string {

	if port == "" {
		return ""
	}

	var aResult string

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Minute)
	defer cancel()

	scanner, err := nmap.NewScanner(
		nmap.WithTargets(aHost),
		nmap.WithPorts(port),
		nmap.WithContext(ctx),
		nmap.WithServiceInfo(),
		nmap.WithSkipHostDiscovery(),
	)
	if err != nil {
		logger.Error("Unable to create nma scanner:")
		return ""
	}

	result, warnings, err := scanner.Run()
	if err != nil {
		logger.Error("Unable to run nma scan")
		return ""
	}

	if warnings != nil {
		logger.Info("Nmap Warnings")
	}

	host := result.Hosts[0]
	aPort := host.Ports[0]
	logger.Info(aHost + " " + strconv.Itoa(int(aPort.ID)) + " " + aPort.Service.Name)
	aResult = strconv.Itoa(int(aPort.ID)) + ":" + aPort.Service.Name

	return aResult
}
