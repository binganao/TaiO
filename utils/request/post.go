package request

import (
	"crypto/tls"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

func Post(url, body string) string {
	client := &http.Client{
		Timeout: 2 * time.Second,
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true,
			},
		},
	}

	req, err := http.NewRequest("POST", url, strings.NewReader(body))
	if err != nil {
		return ""
	}

	resp, err := client.Do(req)

	if err != nil {
		return ""
	}

	defer resp.Body.Close()

	byteBody, _ := ioutil.ReadAll(resp.Body)

	return string(byteBody)
}
