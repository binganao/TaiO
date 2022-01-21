package utils

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func ParseIP(rawInput string) []string {
	var ips []string
	var parts [4][]string

	// 判断是否为 IP 段的形式
	if strings.Contains(rawInput, ";") {
		multiIp := strings.Split(rawInput, ";")
		for _, raw := range multiIp {
			if strings.Contains(raw, "/") {

			} else {
				ipSlice := strings.Split(raw, ".")

				for i, part := range ipSlice {

					// 解析是否含有 , 若有则逐一加入临时列表中
					if strings.Contains(part, ",") {
						ev := strings.Split(part, ",")
						for _, item := range ev {

							// 解析是否含有 - ，若有则将全部都加入一个临时的列表中
							if strings.Contains(item, "-") {
								ft := strings.Split(item, "-")
								if len(ft) == 2 {
									f, _ := strconv.Atoi(ft[0])
									t, _ := strconv.Atoi(ft[1])
									if f < t {
										for from := f; from <= t; from++ {
											// 如果不重复则加入临时列表
											if !sliceContains(strconv.Itoa(from), parts[i]) {
												parts[i] = append(parts[i], strconv.Itoa(from))
											}
										}
									}
								}
							} else {
								// 判断是否为 单个 IP
								if len(item) <= 3 {
									// 如果不重复则加入临时列表
									if !sliceContains(item, parts[i]) {
										parts[i] = append(parts[i], item)
									}
								}
							}
						}
					} else if strings.Contains(part, "-") {
						ft := strings.Split(part, "-")
						if len(ft) == 2 {
							f, _ := strconv.Atoi(ft[0])
							t, _ := strconv.Atoi(ft[1])
							if f < t {
								for from := f; from <= t; from++ {
									// 如果不重复则加入临时列表
									if !sliceContains(strconv.Itoa(from), parts[i]) {
										parts[i] = append(parts[i], strconv.Itoa(from))
									}
								}
							}
						}
					} else {
						parts[i] = append(parts[i], part)
					}
				}
			}
		}
	} else {
		if strings.Contains(rawInput, "/") {

		} else {
			ipSlice := strings.Split(rawInput, ".")

			for i, part := range ipSlice {

				// 解析是否含有 , 若有则逐一加入临时列表中
				if strings.Contains(part, ",") {
					ev := strings.Split(part, ",")
					for _, item := range ev {

						// 解析是否含有 - ，若有则将全部都加入一个临时的列表中
						if strings.Contains(item, "-") {
							ft := strings.Split(item, "-")
							if len(ft) == 2 {
								f, _ := strconv.Atoi(ft[0])
								t, _ := strconv.Atoi(ft[1])
								if f < t {
									for from := f; from <= t; from++ {
										// 如果不重复则加入临时列表
										if !sliceContains(strconv.Itoa(from), parts[i]) {
											parts[i] = append(parts[i], strconv.Itoa(from))
										}
									}
								}
							}
						} else {
							// 判断是否为 单个 IP
							if len(item) <= 3 {
								// 如果不重复则加入临时列表
								if !sliceContains(item, parts[i]) {
									parts[i] = append(parts[i], item)
								}
							}
						}
					}
				} else if strings.Contains(part, "-") {
					ft := strings.Split(part, "-")
					if len(ft) == 2 {
						f, _ := strconv.Atoi(ft[0])
						t, _ := strconv.Atoi(ft[1])
						if f < t {
							for from := f; from <= t; from++ {
								// 如果不重复则加入临时列表
								if !sliceContains(strconv.Itoa(from), parts[i]) {
									parts[i] = append(parts[i], strconv.Itoa(from))
								}
							}
						}
					}
				} else {
					parts[i] = append(parts[i], part)
				}
			}
		}
	}

	for i := range parts {
		sort.Strings(parts[i])
	}

	for _, i := range parts[0] {
		for _, j := range parts[1] {
			for _, k := range parts[2] {
				for _, l := range parts[3] {
					ips = append(ips, fmt.Sprintf("%s.%s.%s.%s", i, j, k, l))
				}
			}
		}
	}

	ips = removeRepByMap(ips)

	return ips
}

func sliceContains(sub string, list []string) bool {
	flag := true
	for _, already := range list {
		al, _ := strconv.Atoi(already)
		it, _ := strconv.Atoi(sub)
		if al == it {
			flag = false
		}
	}
	if flag {
		return false
	}
	return true
}

//slice去重
func removeRepByMap(slc []string) []string {
	result := []string{}
	tempMap := map[string]byte{}
	for _, e := range slc {
		l := len(tempMap)
		tempMap[e] = 0

		if len(tempMap) != l {
			result = append(result, e)
		}
	}
	return result
}
