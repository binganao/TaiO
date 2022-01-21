package main

import (
	"github.com/binganao/Taio/common"
	"github.com/binganao/Taio/lib"
	"github.com/binganao/Taio/service"
)

func main() {
	lib.InitConfig()
	common.InitValue()
	lib.InitDB()

	//r := gin.Default()
	//
	//v1 := r.Group("api/v1")
	//{
	//	v1.GET("test", test.Test)
	//	v1.GET("job", job.AddJob)
	//}
	//
	//r.Run()
	service.Probe("")
}
