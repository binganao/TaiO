package finger

import (
	"bytes"
	"fmt"
	"os/exec"
	"strings"
)

func FingerScan(ip, services string) []string {
	var cmd *exec.Cmd
	var outb, errs bytes.Buffer

	var result []string

	var args []string

	args = append(args, "finger")
	args = append(args, "-u")

	if strings.ContainsAny(services, ":") {
		part := strings.Split(services, ":")
		if part[1] == "https" {
			args = append(args, "https://"+ip+":"+part[0])
			result = append(result, "https://"+ip+":"+part[0])
		} else if part[1] == "http" {
			args = append(args, "http://"+ip+":"+part[0])
			result = append(result, "http://"+ip+":"+part[0])
		} else {
			return nil
		}
	} else {
		return nil
	}

	cmd = exec.Command("bin/ehole", args...)

	cmd.Stdout = &outb
	cmd.Stderr = &errs

	err := cmd.Start()
	if err != nil {
		if errs.Len() > 0 {
			fmt.Printf("err", err)
			return nil
		}
		return nil
	}

	cmd.Wait()

	if len(outb.Bytes()) != 0 {
		out := string(outb.Bytes())
		if !strings.ContainsAny(out, "]") {
			return nil
		}
		part := strings.Split(out, "|")
		if len(part) > 1 {
			part1 := strings.Split(part[1], " ")
			part2 := strings.Split(part[2], " ")
			result = append(result, part1[1])
			result = append(result, part2[1])
		}
		return result
	}

	return nil
}
