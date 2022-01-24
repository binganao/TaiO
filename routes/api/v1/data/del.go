package data

import (
	"github.com/binganao/Taio/common"
	"github.com/binganao/Taio/model/response"
	"github.com/binganao/Taio/service/datas"
	"github.com/binganao/Taio/utils/crypto"
	"github.com/gin-gonic/gin"
	"net/http"
)

func DelData(ctx *gin.Context) {
	host, _ := ctx.GetPostForm("host")
	secret, _ := ctx.GetPostForm("secret")

	deHost := crypto.Base64Decrypto(host)
	deSecret := crypto.Base64Decrypto(secret)

	if deHost == "" || deSecret == "" || (deSecret != "" && deSecret != common.SECRET) {
		ctx.JSON(http.StatusBadRequest, response.OptR{
			Code: http.StatusBadRequest,
			Msg:  "操作失败",
		})
		return
	} else {
		datas.Remove(deHost)
		ctx.JSON(http.StatusOK, response.OptR{
			Code: http.StatusOK,
			Msg:  "操作成功",
		})
		return
	}
}
