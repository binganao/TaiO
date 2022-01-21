package job

import (
	"github.com/binganao/Taio/model/response"
	"github.com/binganao/Taio/utils/crypto"
	"github.com/binganao/Taio/utils/parse"
	"github.com/gin-gonic/gin"
	"net/http"
)

func AddJob(c *gin.Context) {
	ips := crypto.Base64Decrypto(c.Query("ips"))
	result := parse.ParseIP(ips)
	c.JSON(http.StatusOK, response.JobResp{Code: http.StatusOK, Result: result, Msg: ""})
}
