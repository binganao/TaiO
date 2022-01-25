package job

import (
	"github.com/binganao/TaiO/model/response"
	"github.com/binganao/TaiO/service/jobs"
	"github.com/binganao/TaiO/utils/crypto"
	"github.com/gin-gonic/gin"
	"net/http"
)

func DelJob(ctx *gin.Context) {

	hosts, _ := ctx.GetPostForm("hosts")

	deHosts := crypto.Base64Decrypto(hosts)

	if deHosts != "" {
		jobs.DelJobs(hosts)
		ctx.JSON(http.StatusOK, response.JobResp{Code: http.StatusOK, Msg: "指令已下达，任务即将删除"})
	} else {
		ctx.JSON(http.StatusOK, response.JobResp{Code: http.StatusOK, Msg: "操作失败，请检查请求"})
	}

}
