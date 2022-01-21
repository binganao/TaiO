package test

import (
	"github.com/binganao/Taio/utils/crypto"
	"github.com/binganao/Taio/utils/parse"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Test(c *gin.Context) {
	ips := crypto.Base64Decrypto(c.Query("ips"))
	c.JSON(http.StatusOK, gin.H{"code": 200, "msg": parse.ParseIP(ips)})
}
