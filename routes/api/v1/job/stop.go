package job

import (
	"github.com/binganao/TaiO/model/response"
	"github.com/binganao/TaiO/service/jobs"
	"github.com/gin-gonic/gin"
	"net/http"
)

func StopJob(c *gin.Context) {
	jobs.StopJobs()
	c.JSON(http.StatusOK, response.JobResp{Code: http.StatusOK, Msg: "指令已下达"})
}
