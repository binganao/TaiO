package job

import (
	"github.com/binganao/TaiO/model/response"
	"github.com/binganao/TaiO/service/jobs"
	"github.com/binganao/TaiO/utils/crypto"
	"github.com/gin-gonic/gin"
	"net/http"
)

func AddJob(ctx *gin.Context) {
	hosts, _ := ctx.GetPostForm("hosts")

	deHosts := crypto.Base64Decrypto(hosts)

	if deHosts != "" {
		jobs.AddJobs(deHosts, true)
		ctx.JSON(http.StatusOK, response.JobResp{Code: http.StatusOK, Msg: "操作成功，任务已经添加"})
	} else {
		ctx.JSON(http.StatusOK, response.JobResp{Code: http.StatusOK, Msg: "操作失败，请检查请求"})
	}
}
