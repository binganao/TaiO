package main

import (
	"github.com/binganao/TaiO/common"
	"github.com/binganao/TaiO/lib"
	middleware "github.com/binganao/TaiO/middlewares"
	"github.com/binganao/TaiO/routes"
	"github.com/binganao/TaiO/service"
	"github.com/binganao/TaiO/service/jobs"
	"github.com/gin-gonic/gin"
	"io/ioutil"
)

func main() {
	lib.Banner()
	lib.InitConfig()
	common.InitValue()
	lib.InitDB()

	if common.NODE_TYPE == "node" {
		jobs.InitJobs()
		go service.AyncProbe()
		for true {

		}
	} else if common.NODE_TYPE == "data" {
		gin.SetMode(gin.ReleaseMode)
		gin.DefaultWriter = ioutil.Discard
		r := gin.Default()

		r.Use(middleware.Cors())
		routes.V1(r)

		r.Run()
	} else {
		jobs.InitJobs()
		go service.AyncProbe()

		gin.SetMode(gin.ReleaseMode)
		gin.DefaultWriter = ioutil.Discard
		r := gin.Default()

		r.Use(middleware.Cors())
		routes.V1(r)

		r.Run()
	}
}
