package tmp

import (
	"github.com/binganao/TaiO/model/response"
	"github.com/binganao/TaiO/service/datas"
	"github.com/binganao/TaiO/utils/crypto"
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
		datas.RemoveTmp(deHost)
		ctx.JSON(http.StatusOK, response.OptR{
			Code: http.StatusOK,
			Msg:  "操作成功",
		})
		return
	}
}
