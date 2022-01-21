package main

import (
	"github.com/binganao/Taio/common"
	"github.com/binganao/Taio/lib"
	"github.com/binganao/Taio/routes/api/v1/job"
	"github.com/binganao/Taio/routes/api/v1/search"
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

	v1 := r.Group("api/v1")
	{
		v1.GET("search", search.Single)

		jobR := v1.Group("job")
		{
			jobR.GET("add", job.AddJob)
			jobR.GET("del", job.DelJob)
		}
	}

	r.Run()
	//service.Probe("121.42.148.29")
	//service.Test()
}
