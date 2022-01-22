package nma

import (
	"github.com/binganao/Taio/pkg/logger"
	"runtime"
	"sync"
)

//func NmapScan(ip string, portSlice []string) []string {
//	var ports string
//	for i, p := range portSlice {
//		if i == 0 {
//			ports += p
//		} else {
//			ports += "," + p
//		}
//	}
//	var results []string
//	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Minute)
//	defer cancel()
//
//	scanner, err := nmap.NewScanner(
//		nmap.WithTargets(ip),
//		nmap.WithPorts(ports),
//		nmap.WithContext(ctx),
//		nmap.WithServiceInfo(),
//	)
//	if err != nil {
//		logger.Error("Unable to create nma scanner:")
//		return nil
//	}
//
//	result, warnings, err := scanner.Run()
//	if err != nil {
//		logger.Error("Unable to run nma scan")
//		return nil
//	}
//
//	if warnings != nil {
//		logger.Info("Nmap Warnings")
//	}
//
//	for _, host := range result.Hosts {
//		if len(host.Ports) == 0 || len(host.Addresses) == 0 {
//			continue
//		}
//
//		for _, port := range host.Ports {
//			results = append(results, strconv.Itoa(int(port.ID))+":"+port.Service.Name)
//		}
//	}
//
//	return results
//}

func NmapScan(ip string, portSlice []string) []string {
	logger.Info("开始对目标 " + ip + " 进行服务识别")

	var wg sync.WaitGroup
	//var m sync.Mutex

	maxProcs := runtime.NumCPU()
	runtime.GOMAXPROCS(maxProcs)

	var result []string

	for _, p := range portSlice {
		wg.Add(1)
		go func(port string) {
			r := aynscNmap(ip, port)
			if r != "" {
				result = append(result, r)
			}
			wg.Done()
		}(p)
	}
	wg.Wait()

	return result
}
