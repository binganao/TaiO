package main

import (
	"github.com/binganao/Taio/common"
	"github.com/binganao/Taio/lib"
	"github.com/binganao/Taio/routes"
	"github.com/binganao/Taio/service"
	"github.com/binganao/Taio/service/jobs"
	"github.com/gin-gonic/gin"
)

func main() {
	lib.InitConfig()
	common.InitValue()
	lib.InitDB()

	jobs.InitJobs()
	go service.AyncProbe()

	r := gin.Default()

	routes.V1(r)

	r.Run()
}
