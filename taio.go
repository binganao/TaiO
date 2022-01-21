package main

import (
	"fmt"
	"github.com/binganao/Taio/utils"
)

func main() {
	fmt.Println(len(utils.ParseIP("127.0.0.1-255;127.0.1-233.2")))
}
