package lib

import (
	"fmt"
	"github.com/binganao/Taio/pkg/logger"
)

func Banner() {
	banner := "\n" +
		"████████╗ █████╗ ██╗ ██████╗ \n" +
		"╚══██╔══╝██╔══██╗██║██╔═══██╗\n" +
		"   ██║   ███████║██║██║   ██║\n" +
		"   ██║   ██╔══██║██║██║   ██║\n" +
		"   ██║   ██║  ██║██║╚██████╔╝\n" +
		"   ╚═╝   ╚═╝  ╚═╝╚═╝ ╚═════╝ \n"
	sub := "" +
		"Author: binganao\t Github: https://github.com/binganao/Taio\n" +
		"TaiO 的定位是一款用于攻击方对靶标资产梳理，快速定位脆弱资产的工具"
	fmt.Println(logger.LightWhite(banner))
	fmt.Println(sub)
}
