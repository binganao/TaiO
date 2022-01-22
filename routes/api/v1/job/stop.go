package job

import (
	"github.com/binganao/Taio/model/response"
	"github.com/binganao/Taio/service/jobs"
	"github.com/gin-gonic/gin"
	"net/http"
)

func StopJob(c *gin.Context) {
	jobs.StopJobs()
	c.JSON(http.StatusOK, response.JobResp{Code: http.StatusOK, Msg: "指令已下达"})
}
