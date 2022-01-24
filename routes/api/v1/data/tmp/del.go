package tmp

import (
	"github.com/binganao/Taio/model/response"
	"github.com/binganao/Taio/utils/crypto"
	"github.com/gin-gonic/gin"
	"net/http"
)

func DelDataTmp(ctx *gin.Context) {
	host, _ := ctx.GetPostForm("host")

	deHost := crypto.Base64Decrypto(host)

	if deHost == "" {
		ctx.JSON(http.StatusBadRequest, response.OptR{
			Code: http.StatusBadRequest,
			Msg:  "操作失败",
		})
		return
	} else {

	}
}
