package job

import (
	"github.com/binganao/Taio/model/response"
	"github.com/binganao/Taio/service/jobs"
	"github.com/binganao/Taio/utils/crypto"
	"github.com/gin-gonic/gin"
	"net/http"
)

func AddJob(c *gin.Context) {
	hosts := crypto.Base64Decrypto(c.Query("hosts"))
	jobs.AddJobs(hosts, true)
	c.JSON(http.StatusOK, response.JobResp{Code: http.StatusOK, Msg: "操作成功，任务已经添加"})
}
