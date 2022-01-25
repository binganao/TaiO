package mas

import (
	"context"
	"github.com/binganao/TaiO/common"
	"github.com/binganao/TaiO/pkg/logger"
	"github.com/zan8in/masscan"
	"time"
)

//func MasScan(ip string) []string {
//	var result []string
//	scanner, err := masscan.NewScanner(
//		masscan.SetParamTargets(ip),
//		masscan.SetParamPorts("1-65535"),
//		masscan.SetParamWait(0),
//		masscan.SetParamRate(common.MASSCAN_RATE),
//	)
//	if err != nil {
//		return nil
//	}
//
//	scanResult, _, err := scanner.Run()
//	if err != nil {
//		return nil
//	}
//
//	if scanResult != nil {
//		for i, _ := range scanResult.Hosts {
//			result = append(result, strconv.Itoa(scanResult.Ports[i].Port))
//		}
//	}
//	return result
//}

func MasScan(ip string) []string {
	logger.Info("开始对目标 " + ip + " 进行端口探测")
	var result []string
	context, cancel := context.WithTimeout(context.Background(), 1*time.Minute)
	defer cancel()

	var (
		scannerResult []masscan.ScannerResult
		errorBytes    []byte
	)

	scanner, err := masscan.NewScanner(
		masscan.SetParamTargets(ip),
		masscan.SetParamPorts("1-65535"),
		masscan.SetParamWait(0),
		masscan.SetParamRate(common.MASSCAN_RATE),
		masscan.WithContext(context),
	)

	if err != nil {
		logger.Fatal("Unable to create masscan scanner")
		return result
	}

	if err := scanner.RunAsync(); err != nil {
		logger.Error("Run Async Masscan Scan Error!")
		return result
	}

	stdout := scanner.GetStdout()

	stderr := scanner.GetStderr()

	go func() {
		for stdout.Scan() {
			srs := masscan.ParseResult(stdout.Bytes())
			logger.Info(srs.IP + " " + srs.Port)
			result = append(result, srs.Port)
			scannerResult = append(scannerResult, srs)
		}
	}()

	go func() {
		for stderr.Scan() {
			errorBytes = append(errorBytes, stderr.Bytes()...)
		}
	}()

	if err := scanner.Wait(); err != nil {
		return result
	}

	return result
}
