package routes

import (
	"github.com/binganao/Taio/routes/api/v1/job"
	"github.com/binganao/Taio/routes/api/v1/search"
	"github.com/gin-gonic/gin"
)

func V1(r *gin.Engine) {
	v1 := r.Group("api/v1")
	{
		v1.GET("search", search.Single)

		jobR := v1.Group("job")
		{
			jobR.GET("add", job.AddJob)
			jobR.GET("del", job.DelJob)
		}
	}
}
