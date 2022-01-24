package data

import (
	"github.com/binganao/Taio/common"
	"github.com/binganao/Taio/model/response"
	"github.com/binganao/Taio/service/datas"
	"github.com/binganao/Taio/utils/crypto"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func AddData(c *gin.Context) {
	host, _ := c.GetPostForm("host")
	ports, _ := c.GetPostForm("ports")
	services, _ := c.GetPostForm("services")
	fingers, _ := c.GetPostForm("fingers")
	secret, _ := c.GetPostForm("secret")

	deSecret := crypto.Base64Decrypto(secret)
	if deSecret == "" || (deSecret != "" && deSecret != common.SECRET) {
		c.JSON(http.StatusBadRequest, response.OptR{
			Code: http.StatusBadRequest,
			Msg:  "操作失败",
		})
		return
	} else {
		deHost := crypto.Base64Decrypto(host)
		dePorts := crypto.Base64Decrypto(ports)
		deServices := crypto.Base64Decrypto(services)
		deFingers := crypto.Base64Decrypto(fingers)
		datas.Save(deHost, dePorts, deServices, deFingers, time.Now())
		c.JSON(http.StatusOK, response.OptR{
			Code: http.StatusOK,
			Msg:  "操作成功",
		})
		return
	}
}
