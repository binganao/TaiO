package main

import (
	"github.com/binganao/Taio/routes/api/v1/job"
	"github.com/binganao/Taio/routes/api/v1/test"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	v1 := r.Group("api/v1")
	{
		v1.GET("test", test.Test)
		v1.GET("job", job.AddJob)
	}

	r.Run()
}
