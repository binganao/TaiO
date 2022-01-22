package lib

import (
	"fmt"
	"github.com/binganao/Taio/pkg/logger"
)

func Banner() {
	banner := "" +
		"████████╗ █████╗ ██╗ ██████╗ \n" +
		"╚══██╔══╝██╔══██╗██║██╔═══██╗\n" +
		"   ██║   ███████║██║██║   ██║\n" +
		"   ██║   ██╔══██║██║██║   ██║\n" +
		"   ██║   ██║  ██║██║╚██████╔╝\n" +
		"   ╚═╝   ╚═╝  ╚═╝╚═╝ ╚═════╝ \n"
	sub := "Author: binganao\t Github: https://github.com/binganao/Taio"
	fmt.Println(logger.LightWhite(banner))
	fmt.Println(sub)
}
