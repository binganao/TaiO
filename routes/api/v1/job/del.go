package job

import (
	"github.com/binganao/Taio/model/response"
	"github.com/binganao/Taio/service/jobs"
	"github.com/binganao/Taio/utils/crypto"
	"github.com/gin-gonic/gin"
	"net/http"
)

func DelJob(c *gin.Context) {
	hosts := crypto.Base64Decrypto(c.Query("hosts"))
	jobs.DelJobs(hosts)
	c.JSON(http.StatusOK, response.JobResp{Code: http.StatusOK, Msg: "指令已下达，任务即将删除"})
}
