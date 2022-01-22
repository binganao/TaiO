package nma

import (
	"sync"
)

func NmapScan(ip string, portSlice []string) []string {

	var result []string

	if len(portSlice) > 10 {
		for i := 0; i < len(portSlice); i++ {
			if i%9 == 0 {
				continue
			}
			var t []string
			t = append(t, portSlice[i])
			s := NmapScan(ip, t)
			for _, j := range s {
				result = append(result, j)
			}
		}
	} else {
		var wg sync.WaitGroup
		for _, p := range portSlice {
			wg.Add(1)
			go func(port string) {
				defer wg.Done()
				r := aynscNmap(ip, port)
				if r != "" {
					result = append(result, r)
				}
			}(p)
		}
		wg.Wait()
	}

	return result
}
