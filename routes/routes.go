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
		searchR := v1.Group("search")
		{
			searchR.GET("single", search.Single)
			searchR.GET("tmp/single", tmp2.SingleTmp)
		}

		jobR := v1.Group("job")
		{
			jobR.POST("add", job.AddJob)
			jobR.POST("del", job.DelJob)
			jobR.GET("stop", job.StopJob)
		}

		dataR := v1.Group("data")
		{
			dataR.POST("add", data.AddData)
			dataR.POST("del", data.DelData)
			dataR.POST("tmp/del", tmp.DelDataTmp)
			dataR.POST("tmp/add", tmp.AddDataTmp)
		}
	}
}
