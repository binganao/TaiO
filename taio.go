package main

import (
	"github.com/binganao/Taio/common"
	"github.com/binganao/Taio/lib"
	"github.com/binganao/Taio/routes"
	"github.com/binganao/Taio/service"
	"github.com/binganao/Taio/service/jobs"
	"github.com/gin-gonic/gin"
	"io/ioutil"
)

func main() {
	lib.Banner()
	lib.InitConfig()
	common.InitValue()
	lib.InitDB()

	jobs.InitJobs()
	go service.AyncProbe()

	gin.SetMode(gin.ReleaseMode)
	gin.DefaultWriter = ioutil.Discard
	r := gin.Default()

	routes.V1(r)

	r.Run()
}
