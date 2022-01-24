package routes

import (
	"github.com/binganao/Taio/routes/api/v1/data"
	"github.com/binganao/Taio/routes/api/v1/data/tmp"
	"github.com/binganao/Taio/routes/api/v1/job"
	"github.com/binganao/Taio/routes/api/v1/search"
	tmp2 "github.com/binganao/Taio/routes/api/v1/search/tmp"
	"github.com/gin-gonic/gin"
)

func V1(r *gin.Engine) {
	v1 := r.Group("api/v1")
	{
		v1.GET("search", search.Single)
		v1.GET("tmp/search", tmp2.SingleTmp)

		jobR := v1.Group("job")
		{
			jobR.GET("add", job.AddJob)
			jobR.GET("del", job.DelJob)
			jobR.GET("stop", job.StopJob)
		}

		dataR := v1.Group("data")
		{
			dataR.POST("add", data.AddData)
			dataR.POST("tmp/del", tmp.DelDataTmp)
			dataR.POST("tmp/add", tmp.AddDataTmp)
		}
	}
}
